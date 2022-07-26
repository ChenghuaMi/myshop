package cache

import (
	"context"
	"github.com/eko/gocache/v2/store"
	"time"
)

type Options struct {
	Cost int64

	Expiration time.Duration

	Tags []string
}


type CacheInterface interface {
	Get(ctx context.Context,key interface{}) (interface{},error)
	Set(ctx context.Context,key,val interface{},s *store.Options) error
	Delete(ctx context.Context,key interface{}) error

}

type Cache struct {
	ctx context.Context
	cache CacheInterface
}
func NewCache(cache CacheInterface) *Cache{
	ctx := context.Background()
	return &Cache{
		ctx:   ctx,
		cache: cache,
	}
}
func (ca *Cache) Get(key interface{}) (interface{},error) {
	return ca.cache.Get(ca.ctx,key)
}
func (ca *Cache) Set(key,val interface{},s *Options) error {
	ss := store.Options(*s)
	return ca.cache.Set(ca.ctx,key,val, &ss)

}
func (ca *Cache) Delete(key interface{}) error {
	return ca.cache.Delete(ca.ctx,key)
}
