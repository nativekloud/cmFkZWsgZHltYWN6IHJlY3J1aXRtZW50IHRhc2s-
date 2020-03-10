.PHONY: proto, test, build, docker, request,build-linux

proto:
	protoc --proto_path=$$GOPATH/src:. --micro_out=. --go_out=. proto/weather.proto
run:
	go run main.go

build: proto
	go build -o weather-srv *.go

build-linux: proto
	GOOS=linux go build -o weather-srv *.go
api:
	micro api
web:
	micro web

docker:
	docker build . -t weather-srv:latest
test:
	go test -v ./... -cover

request:
	curl "http://localhost:8080/weatherservice/WeatherService/get?cities[]=Paris&cities[]=London"
