package infrastruct

import (
	"net/url"

	"github.com/go-redis/redis"
)

// NewRedis .
func NewRedis(url *url.URL) (*redis.Client, error) {
	opts, err := redis.ParseURL(url.String())
	if err != nil {
		return nil, err
	}

	return redis.NewClient(opts), nil
}
