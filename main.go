package main

import (
	"fmt"
	"log"
	"os"
	"time"

	proto "weatherservice/proto"

	handler "weatherservice/handler"
	ow "weatherservice/openweather"

	micro "github.com/micro/go-micro/v2"
	"github.com/patrickmn/go-cache"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.weatherservice"),
	)
	service.Init()

	APIKey := os.Getenv("OWM_API_KEY")

	if len(APIKey) == 0 {
		log.Fatalf("No API keys present. Set OWM_API_KEY env variable.")
	}
	// ow client and cache initialize
	client := ow.NewOpenWeatherMapClient(APIKey)
	c := cache.New(5*time.Minute, 10*time.Minute)

	proto.RegisterWeatherServiceHandler(service.Server(), &handler.Weather{
		Client: ow.NewCachedClient(client, c),
	})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
