package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	proto "weatherservice/proto"

	handler "weatherservice/handler"
	ow "weatherservice/openweather"

	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/patrickmn/go-cache"
)

func main() {

	service := micro.NewService(
		micro.Name("go.micro.api.weatherservice"),
		micro.Flags(
			&cli.StringFlag{
				Name:  "owm_api_key",
				Usage: "Set OpenWeatherMap API key",
			},
		),
	)
	APIKey := os.Getenv("OWM_API_KEY")

	service.Init(
		micro.Action(func(c *cli.Context) error {
			fmt.Printf("The OWM_API_KEY flag is: %s\n", c.String("owm_api_key"))
			if len(APIKey) == 0 {
				APIKey = c.String("owm_api_key")
			}
			return nil
		}),
	)

	if len(APIKey) == 0 {
		log.Fatalf("No API keys present. Set OWM_API_KEY env variable.")
	}

	// http client
	httpClient := &http.Client{
		Timeout: time.Second * 2,
	}

	// ow client and cache initialize
	client := ow.NewOpenWeatherMapClient(APIKey, httpClient)

	c := cache.New(5*time.Minute, 10*time.Minute)

	proto.RegisterWeatherServiceHandler(service.Server(), &handler.Weather{
		Client: ow.NewCachedClient(client, c),
	})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
