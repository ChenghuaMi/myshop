package config

import "time"

type Cache struct {
	CacheType string `mapstructure:"cacheType"`
	FreeCache  `mapstructure:"freecache"`
	Redis `mapstructure:"redis"`
}

type FreeCache struct {
	Size int   `mapstructure:"size"`
	Expiration time.Duration  `mapstructure:"expiration"`
}