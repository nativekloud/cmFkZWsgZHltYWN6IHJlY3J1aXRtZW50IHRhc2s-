package openweather

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func NewOpenWeatherMapClient(key string) Client {
	return OpenWeatherMap{
		API_KEY: key,
	}
}

type OpenWeatherMap struct {
	API_KEY string
}

const (
	API_URL string = "api.openweathermap.org"
)

func makeAPIRequest(url string) ([]byte, error) {
	// Build an http client so we can have control over timeout
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	res, getErr := client.Get(url)
	if getErr != nil {
		return nil, getErr
	}

	// defer the closing of the res body
	defer res.Body.Close()

	// read the http response body into a byte stream
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}

func (owm OpenWeatherMap) CurrentWeatherFromCity(city string) ([]byte, error) {
	if owm.API_KEY == "" {
		// No API keys present, return error
		return nil, errors.New("No API keys present")
	}
	url := fmt.Sprintf("http://%s/data/2.5/weather?q=%s&units=imperial&APPID=%s", API_URL, city, owm.API_KEY)

	body, err := makeAPIRequest(url)
	if err != nil {
		return nil, err
	}

	return body, nil
}
