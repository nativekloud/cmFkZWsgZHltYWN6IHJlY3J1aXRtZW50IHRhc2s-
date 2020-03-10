package handler

import (
	"context"
	"encoding/json"
	ow "weatherservice/openweather"
	proto "weatherservice/proto"

	"golang.org/x/sync/errgroup"
)

// Weather service
type Weather struct {
	Client ow.Client
}

// Get fetches weather in requested cities
func (w *Weather) Get(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	results, err := fetchAllErrGroup(ctx, w.Client, req.Cities)
	if err != nil {
		return err
	}
	// unmarshall results to proto
	response, errProto := toProto(results)
	if errProto != nil {
		return err
	}
	rsp.Response = response
	return nil
}

func fetchAllErrGroup(ctx context.Context, c ow.Client, cities []string) ([][]byte, error) {
	g, ctx := errgroup.WithContext(ctx)
	response := make([][]byte, len(cities))
	// range over cities and fetch results
	for i, city := range cities {
		i, city := i, city //https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			currentWeather, err := c.CurrentWeatherFromCity(city)
			if err != nil {
				return err
			}
			response[i] = currentWeather
			return nil
		})
	}
	// wait for all request to finish or get error
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return response, nil
}

func toProto(response [][]byte) ([]*proto.CurrentWeatherResponse, error) {
	results := make([]*proto.CurrentWeatherResponse, len(response))
	for i, r := range response {
		var cwr proto.CurrentWeatherResponse
		// unmarshal the byte stream into a proto
		err := json.Unmarshal(r, &cwr)
		if err != nil {
			return results, err
		}
		results[i] = &cwr
	}
	return results, nil
}
