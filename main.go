package main

import (
	"fmt"

	proto "weatherservice/proto"

	handler "weatherservice/handler"

	micro "github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.weatherservice"),
	)
	service.Init()

	proto.RegisterWeatherServiceHandler(service.Server(), new(handler.Weather))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
