package handler

import (
	"context"
	"encoding/json"
	"os"
	ow "weatherservice/openweather"
	proto "weatherservice/proto"
)

type Weather struct{}

func (w *Weather) Get(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	//rsp.Response = req.Cities
	client := ow.OpenWeatherMap{API_KEY: os.Getenv("OWM_API_KEY")}
	currentWeather, err := client.CurrentWeatherFromCity("London")
	if err != nil {
		return err
	}
	var cwr proto.CurrentWeatherResponse

	// unmarshal the byte stream into a Go data type
	jsonErr := json.Unmarshal(currentWeather, &cwr)
	if jsonErr != nil {
		return jsonErr
	}
	rsp.Response = []*proto.CurrentWeatherResponse{&cwr}
	return nil
}
