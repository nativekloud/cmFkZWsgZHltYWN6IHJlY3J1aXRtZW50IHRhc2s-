.PHONY: proto, test

proto:
	protoc --proto_path=$$GOPATH/src:. --micro_out=. --go_out=. proto/weather.proto
run:
	go run main.go
api:
	micro api
web:
	micro web
test:
	curl "http://localhost:8080/weatherservice/WeatherService/get?cities[]=Paris&cities[]=London"