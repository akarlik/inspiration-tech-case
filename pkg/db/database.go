package db

import (
	"inspiration-tech-case/configuration"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

func NewDatabase(config configuration.Config) *cache.Cache {
	de := config.Get("DATASOURCE.DEFAULT_EXPIRATION")
	defaultExpiration, _ := strconv.ParseUint(de, 10, 32)

	ci := config.Get("DATASOURCE.CLEANUP_INTERVAL")
	cleanupInterval, _ := strconv.ParseUint(ci, 10, 32)

	accountCache := cache.New(time.Duration(defaultExpiration)*time.Minute, time.Duration(cleanupInterval)*time.Minute)

	return accountCache
}
