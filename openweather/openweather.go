package openweather

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func NewOpenWeatherMapClient(key string, httpClient *http.Client) Client {
	return API{
		API_KEY:    key,
		httpClient: httpClient,
	}
}

type API struct {
	API_KEY    string
	httpClient *http.Client
}

const (
	API_URL string = "api.openweathermap.org"
)

// request API URL and return response []byte
func (api API) request(url string) ([]byte, error) {
	// just need GET
	response, err := api.httpClient.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP request failed.")
	}

	defer response.Body.Close()
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}
	return body, nil
}

func (api API) CurrentWeatherFromCity(city string) ([]byte, error) {
	if api.API_KEY == "" {
		// No API keys present, return error
		return nil, errors.New(errNoAPIKeys)
	}
	url := fmt.Sprintf("http://%s/data/2.5/weather?q=%s&units=imperial&APPID=%s", API_URL, city, api.API_KEY)

	body, err := api.request(url)
	if err != nil {
		return nil, err
	}

	return body, nil
}
