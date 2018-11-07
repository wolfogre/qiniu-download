package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var base *cache.Cache

func init() {
	base = cache.New(time.Hour, time.Minute)
}
