.PHONY: proto

proto:
	protoc --proto_path=$$GOPATH/src:. --micro_out=. --go_out=. proto/weather.proto
run:
	go run main.go
api:
	micro api
web:
	micro web
test:
	curl http://localhost:8080/weatherservice/WeatherService/get?cities%5B%5D=London