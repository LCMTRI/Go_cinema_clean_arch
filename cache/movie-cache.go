package cache

import (
	"Go_cinema_reconstructed/model"

	"github.com/go-redis/redis/v7"
)

type MovieCache interface {
	getClient() *redis.Client
	Set(key string, value *model.MovieRes)
	Get(key string) *model.MovieRes
	// SetEx(key string, value *model.MovieRes, time int)
	// GetEx(key string)
}
