package handler

import (
	"context"
	"errors"
	"testing"

	proto "weatherservice/proto"
	weather "weatherservice/proto"

	"github.com/stretchr/testify/require"
)

type MockClient struct {
	GetCall struct {
		Recives struct {
			City string
		}
		Returns struct {
			Response []byte
			Error    error
		}
	}
}

func (m *MockClient) CurrentWeatherFromCity(city string) ([]byte, error) {
	m.GetCall.Recives.City = city
	return m.GetCall.Returns.Response, m.GetCall.Returns.Error
}

func TestGetHandler(t *testing.T) {
	t.Run("hanlder returns empty response for emty request", func(t *testing.T) {
		mockClient := &MockClient{}
		mockClient.GetCall.Returns.Response = nil
		mockClient.GetCall.Returns.Error = nil
		api := &Weather{
			Client: mockClient,
		}
		rsp := &proto.Response{}
		err := api.Get(context.TODO(), &proto.Request{
			Cities: []string{},
		}, rsp)

		require.NoError(t, err)
		require.Equal(t, []*weather.CurrentWeatherResponse{}, rsp.Response)

	})
	t.Run("hanlder returns error", func(t *testing.T) {
		mockClient := &MockClient{}

		mockClient.GetCall.Returns.Response = nil
		mockClient.GetCall.Returns.Error = errors.New("Some error")
		api := &Weather{
			Client: mockClient,
		}
		rsp := &proto.Response{}
		err := api.Get(context.TODO(), &proto.Request{
			Cities: []string{"Paris"},
		}, rsp)

		require.Error(t, err)
	})
	t.Run("handler returns response", func(t *testing.T) {
		mockClient := &MockClient{}
		data := []byte(`{"coord":{"lon":2.35,"lat":48.85},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"main":{"temp":53.92,"feelsLike":44.35,"tempMin":52,"tempMax":55.4,"pressure":1013,"humidity":87},"wind":{"speed":17.22,"deg":240},"sys":{"type":1,"id":6550,"country":"FR","sunrise":1583820831,"sunset":1583862477},"clouds":{"all":90},"name":"Paris"}`)
		mockClient.GetCall.Returns.Response = data
		mockClient.GetCall.Returns.Error = nil
		api := &Weather{
			Client: mockClient,
		}
		rsp := &proto.Response{}
		err := api.Get(context.TODO(), &proto.Request{
			Cities: []string{"Paris"},
		}, rsp)
		require.NoError(t, err)
		require.Equal(t, "Paris", rsp.Response[0].Name)
	})
}
