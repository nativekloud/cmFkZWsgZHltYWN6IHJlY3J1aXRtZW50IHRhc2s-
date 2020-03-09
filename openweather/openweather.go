package openweather

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type HTTPClient interface {
	Get(string) (*http.Response, error)
}

func NewOpenWeatherMapClient(key string, httpClient HTTPClient) Client {
	return API{
		APIKey:     key,
		httpClient: httpClient,
	}
}

type API struct {
	APIKey     string
	httpClient HTTPClient
}

const (
	APIUrl string = "api.openweathermap.org"
)

// request API URL and return response []byte
func (api API) request(url string) ([]byte, error) {
	// just need GET
	response, err := api.httpClient.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, errHTTPFailed)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New("request failed")
	}
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return nil, errors.Wrap(readErr, errReadResponseFailed)
	}
	return body, nil
}

func (api API) CurrentWeatherFromCity(city string) ([]byte, error) {
	if api.APIKey == "" {
		// No API keys present, return error
		return nil, errors.New(errNoAPIKeys)
	}
	url := fmt.Sprintf("http://%s/data/2.5/weather?q=%s&units=imperial&APPID=%s", APIUrl, city, api.APIKey)

	body, err := api.request(url)
	if err != nil {
		return nil, err
	}

	return body, nil
}
