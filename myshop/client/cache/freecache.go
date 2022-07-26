package cache

import (
	"github.com/coocood/freecache"
	"github.com/eko/gocache/v2/store"
	"context"
	"myshop/client/global"
	"time"
)

type FreeCache struct {
	cacheManager *store.FreecacheStore
}
func NewFreeCache() *FreeCache {

	freecacheStore := store.NewFreecache(freecache.NewCache(global.Config.Cache.Size),&store.Options{
		Expiration: global.Config.Cache.Expiration * time.Second,
	})

	return &FreeCache{cacheManager: freecacheStore}
}

func (f *FreeCache) Get(ctx context.Context,key interface{}) (interface{},error) {
	return f.cacheManager.Get(ctx,key)
}
func (f *FreeCache)  Set(ctx context.Context,key,val interface{},s *store.Options) error {
	return f.cacheManager.Set(ctx,key,val,s)
}
func (f *FreeCache)  Delete(ctx context.Context,key interface{}) error {
	return f.cacheManager.Delete(ctx,key)
}
