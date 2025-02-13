FROM golang:1.23.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o weather-api ./cmd/myweather-api/main.go

CMD ["./weather-api"]
