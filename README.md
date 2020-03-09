
# Weather service with caching 

Service is based on [go-micro](https://micro.mu/docs/framework.html) framework simplyfying creating micorservices in go.
Cache uses go-cache but implementation is decupled nad can be swapped as openweather client is implemented via interface

# Run

Configure OpenWeatherMap API key in docker-compose file

Build docker image

```
make docker
```

Run docker compose 

```
docker-compose up &
```

You can as well scale number of backend weather services

```
docker-compose up --scale weather=5
```

You can browser running services at [http://localhost:8082](http://localhost:8082) and API is served at [http://localhost:8080](http://localhost:8080)

Weatherservice runs here [http://localhost:8080/weatherservice/WeatherService/get?cities[]=Paris&cities[]=London](http://localhost:8080/weatherservice/WeatherService/get?cities[]=Paris&cities[]=London)




## Usage

A Makefile and docker-compose files are included for convenience

Build the binary for current OS

```
make build
```
Build the binary for use in Docker container

```
make build-linux
```

Run the service
```
./weather-srv
```

Build a docker image
```
make docker
```

Run curl example request 

```
make live-test
```

Run tests 

```
make test
```


##  Task Description

Let’s say we building a small application enabling users to retrieve information
about the weather in the places of their choosing. Your task is creating a microservice
responsible for fetching current weather conditions in cities specified in the requests.
Specification:
- As a source of the weather information you should use a free API described [here](https://openweathermap.org/current).
- Service should expose one HTTP endpoint that takes a list of city names as a
query parameter and returns information about current weather in each city.
- Since free tier account of the OpenWeather API has limited number of API calls,
the service has to have some kind of caching layer that would prevent
subsequent calls for the same city in short time interval.
- The application has to expose some mechanism of configuration. An option to
specify the HTTP port of the server and an API key is a minimum.

Nice to have:
- Provide a dockerfile that can be used to build and run the application without the
need of having the Go toolchain installed.

## notes

https://github.com/briandowns/openweathermap

http://localhost:8080/weatherservice/WeatherService/get?cities[]=Paris&cities[]=London

go run main.go --owm_api_key="test"

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./weather-srv
```

Build a docker image
```
make docker
```