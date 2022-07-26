package cache

import (
	"context"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	"myshop/client/global"
	"strconv"
)

type CacheRedis struct {
	cacheManager *cache.Cache
}

func NewRedis() *CacheRedis {
	redisStore := store.NewRedis(redis.NewClient(&redis.Options{
		Addr: string(global.Config.Cache.Redis.Host) + ":" + strconv.Itoa(global.Config.Cache.Redis.Port),
	}), nil)
	cacheManager := cache.New(redisStore)
	return &CacheRedis{cacheManager: cacheManager}
}

func (f *CacheRedis) Get(ctx context.Context,key interface{}) (interface{},error) {
	return f.cacheManager.Get(ctx,key)
}
func (f *CacheRedis)  Set(ctx context.Context,key,val interface{},s *store.Options) error {
	return f.cacheManager.Set(ctx,key,val,s)
}

func (f *CacheRedis)  Delete(ctx context.Context,key interface{}) error {
	return f.cacheManager.Delete(ctx,key)
}