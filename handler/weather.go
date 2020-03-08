package handler

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	ow "weatherservice/openweather"
	proto "weatherservice/proto"
)

type Weather struct{}

func (w *Weather) Get(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	//rsp.Response = req.Cities
	client := ow.OpenWeatherMap{API_KEY: os.Getenv("OWM_API_KEY")}
	var wg sync.WaitGroup
	wg.Add(len(req.Cities))
	var response []*proto.CurrentWeatherResponse
	for i := 0; i < len(req.Cities); i++ {
		go func(city string) {
			defer wg.Done()
			currentWeather, _ := client.CurrentWeatherFromCity(city)
			// if err != nil {
			// 	return err
			// }
			var cwr proto.CurrentWeatherResponse
			// unmarshal the byte stream into a Go data type
			_ = json.Unmarshal(currentWeather, &cwr)
			// if jsonErr != nil {
			// 	return jsonErr
			// }
			response = append(response, &cwr)
		}(req.Cities[i])
	}
	wg.Wait()
	rsp.Response = response
	return nil
}
