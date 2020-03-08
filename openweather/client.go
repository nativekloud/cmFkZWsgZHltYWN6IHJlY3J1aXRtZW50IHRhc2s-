package openweather

type Client interface {
	CurrentWeatherFromCity(string) ([]byte, error)
}
