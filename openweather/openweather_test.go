package openweather

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type MockHTTPClient struct {
	CallGet func(url string) (*http.Response, error)
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.CallGet(url)
}

func TestOpenWeatherMapClient(t *testing.T) {
	t.Run("Server error", func(t *testing.T) {
		client := NewOpenWeatherMapClient("key", &MockHTTPClient{
			CallGet: func(string) (*http.Response, error) {
				return nil, errors.New("Error on server")
			},
		})
		_, err := client.CurrentWeatherFromCity("London")
		require.Error(t, err)
	})
	t.Run("No API Key", func(t *testing.T) {
		client := NewOpenWeatherMapClient("", &MockHTTPClient{
			CallGet: func(string) (*http.Response, error) {
				return nil, errors.New("Error on server")
			},
		})
		_, err := client.CurrentWeatherFromCity("London")
		require.Error(t, err)
	})
	t.Run("Sucessful empty response", func(t *testing.T) {
		json := `{}`
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		client := NewOpenWeatherMapClient("key", &MockHTTPClient{
			CallGet: func(string) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       r,
				}, nil
			},
		})
		res, err := client.CurrentWeatherFromCity("London")
		require.NoError(t, err)
		require.Equal(t, []byte(json), res)
	})
	t.Run("401 Unauthorized", func(t *testing.T) {
		json := `{}`
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		client := NewOpenWeatherMapClient("key", &MockHTTPClient{
			CallGet: func(string) (*http.Response, error) {
				return &http.Response{
					StatusCode: 401,
					Body:       r,
				}, nil
			},
		})
		_, err := client.CurrentWeatherFromCity("London")
		require.Error(t, err)
	})
}
