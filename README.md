# Description

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

http://localhost:8080/weatherservice/weather/get?cities[]=test,linux
