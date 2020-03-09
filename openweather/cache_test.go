package openweather

import (
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/require"
)

type CacheTestClient struct {
	results []byte
}

func (c CacheTestClient) CurrentWeatherFromCity(string) ([]byte, error) {
	return c.results, nil
}

func TestCachedClient(t *testing.T) {
	var cache = cache.New(1*time.Minute, 1*time.Minute)
	var expected = []byte("testing")
	var client = CacheTestClient{results: expected}

	var api = NewCachedClient(client, cache)

	t.Run("get fresh results", func(t *testing.T) {
		res, err := api.CurrentWeatherFromCity("test")
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("get results from cache", func(t *testing.T) {
		var oldExpected = expected
		client.results = []byte(string(oldExpected) + "new")
		res, err := api.CurrentWeatherFromCity("test")
		require.NoError(t, err)
		require.Equal(t, oldExpected, res)
	})

	t.Run("flush cache and get fresh.", func(t *testing.T) {
		cache.Flush()
		res, err := api.CurrentWeatherFromCity("test")
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})
}
