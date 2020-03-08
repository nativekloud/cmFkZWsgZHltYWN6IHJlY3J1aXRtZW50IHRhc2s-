package handler

import (
	"context"
	"encoding/json"
	ow "weatherservice/openweather"
	proto "weatherservice/proto"

	"golang.org/x/sync/errgroup"
)

type Weather struct {
	Client ow.Client
}

// Get fetches weather in requested cities
func (w *Weather) Get(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	var g errgroup.Group
	var response []*proto.CurrentWeatherResponse

	// range over cities and fetch results
	for _, city := range req.Cities {
		g.Go(func() error {
			currentWeather, err := w.Client.CurrentWeatherFromCity(city)
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
	// wait for all request to finish or get error
	if err := g.Wait(); err != nil {
		return err
	}
	// happy path
	rsp.Response = response
	return nil
}
