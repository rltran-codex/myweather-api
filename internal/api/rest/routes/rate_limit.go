package routes

// rate limiter that links to IP address.
// resource: https://blog.logrocket.com/rate-limiting-go-application/

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type connection struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu          sync.Mutex
	connections = make(map[string]*connection)
)

func InitRate() {
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, client := range connections {
				if time.Since(client.lastSeen) > (30 * time.Minute) {
					log.Printf("Removed rate limiter for %s, last seen %s", ip, client.lastSeen)
					delete(connections, ip)
				}
			}
			mu.Unlock()
		}
	}()
}

func RateLimit(next func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		mu.Lock()
		if _, found := connections[ip]; !found {
			log.Printf("Adding rate limiter for %s", ip)
			connections[ip] = &connection{limiter: rate.NewLimiter(rate.Every(10*time.Minute), 5)}
		}
		connections[ip].lastSeen = time.Now()

		if !connections[ip].limiter.Allow() {
			mu.Unlock()
			log.Printf("too many requests from %s", ip)
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("API at capacity, try again later."))
			return
		}
		mu.Unlock()
		next(w, r)
	})
}
