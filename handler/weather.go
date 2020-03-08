package handler

import (
	"context"
	"encoding/json"
	"os"
	ow "weatherservice/openweather"
	proto "weatherservice/proto"

	"golang.org/x/sync/errgroup"
)

type Weather struct{}

func (w *Weather) Get(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	//rsp.Response = req.Cities
	client := ow.OpenWeatherMap{API_KEY: os.Getenv("OWM_API_KEY")}
	var g errgroup.Group

	var response []*proto.CurrentWeatherResponse
	for _, city := range req.Cities {
		g.Go(func() error {
			currentWeather, err := client.CurrentWeatherFromCity(city)
			if err != nil {
				return err
			}
			var cwr proto.CurrentWeatherResponse
			// unmarshal the byte stream into a Go data type
			jsonErr := json.Unmarshal(currentWeather, &cwr)
			if jsonErr != nil {
				return jsonErr
			}
			response = append(response, &cwr)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	rsp.Response = response
	return nil
}
