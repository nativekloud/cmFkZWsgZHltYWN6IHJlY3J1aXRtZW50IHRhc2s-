package openweather

import "github.com/patrickmn/go-cache"

func NewCachedClient(client Client, cache *cache.Cache) Client {
	return cachedClient{
		client: client,
		cache:  cache,
	}
}

type cachedClient struct {
	client Client
	cache  *cache.Cache
}

func (c cachedClient) CurrentWeatherFromCity(city string) ([]byte, error) {
	cached, found := c.cache.Get(city)
	if found {
		return cached.([]byte), nil

	}
	live, err := c.client.CurrentWeatherFromCity(city)
	c.cache.Set(city, live, cache.DefaultExpiration)
	return live, err
}
