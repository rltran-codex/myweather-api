# Weather API

- [Intro](#overview)
- [Installation](#installation)
- [Usage](#usage)

# Overview

Project from [roadmap.sh](https://roadmap.sh/projects/weather-api-wrapper-service) built using Go.

This API is just a simple API wrapper that fetches weather data from an external source.

As a prerequisite, sign up for a [redis](https://redis.io/) account to use
as a cache for each request. Additionally, create an account with
[visual crossing](https://www.visualcrossing.com/sign-up) to obtain an API key
for the weather data. Both of these services are **free**.
The rate limit is set to a maximum of 5 requests per 10 minutes per IP address. If no requests are made from the same IP address within 30 minutes, the rate limit counter will be reset.

# Installation

1. Clone this repo
2. Create environment variables or use a `.env` file (must be placed in configs folder). Edit the `configs.json`.

```
redis_username=
redis_password=
visual_crossing_api_key=
```

```json
{
  "Address": "0.0.0.0:8080",
  "ReadTimeout": 200, // seconds
  "WriteTimeout": 600, // seconds
  "CacheExpire": 30 // minutes
}
```

4. run project from the project's root directory: `go run ./cmd/myweather-api/main.go`

## Build

### Docker

```
docker build -t myweather-api .
docker run -p <port1>:<port2> myweather-api
```

**Note**: The port2 is the port from config.json, which is needed to expose the port from the docker container.

### Make

Builds executable and copies all configs into the `bin` directory

```
make all
```

**_OPTIONAL_**:
It is possible to build this application for a Raspberry Pi 5 by doing the command:
`make build-raspdist`. This will create the executable suitable for the device.

# Usage

Once the application is running, perform a GET request on the endpoint.

```
curl http://localhost:8080/api/v1/weather?city=London
```

- **Response Format**: This will return weather information for the city of London in JSON format.

# Lessons

Listed below are topics and skills I learned from this:

- Learned how to structure API
  - HTTP handling and routing [1](https://www.manning.com/books/go-web-programming)
  - Utilizing environment variables to protect sensitive information such as keys and account info.
- Use Redis for caching
- Rate limiting [2](https://blog.logrocket.com/rate-limiting-go-application/)
- Build and deploy
  - Docker [3](https://www.geeksforgeeks.org/what-is-dockerfile/) [4](https://docs.docker.com/guides/golang/build-images/)
  - Makefile [5](https://tutorialedge.net/golang/makefiles-for-go-developers/)
