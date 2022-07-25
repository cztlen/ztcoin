package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var GoCache *cache.Cache

//存储链上数据

func init() {
	GoCache = cache.New(2*time.Minute, 5*time.Minute)
}
