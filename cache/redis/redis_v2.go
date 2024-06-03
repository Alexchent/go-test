package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// https://redis.io/docs/clients/go/

type CacheService struct {
	*redis.Client
}

func NewCacheService(options redis.Options) *CacheService {
	client := redis.NewClient(&options)
	return &CacheService{client}
}

func (s *CacheService) Set(key string, value interface{}, expiration time.Duration) {
	err := Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
	return
}

func (s *CacheService) Get(key string) (val string) {
	val, _ = Client.Get(ctx, key).Result()
	return
}

func (s *CacheService) GetToStruct(key string) any {
	var value interface{}
	err := Client.Get(context.Background(), key).Scan(value)
	if err != nil {
		panic(err)
	}
	return value
}

func (s *CacheService) StoreMap(key string, value map[string]string) {
	for k, v := range value {
		err := Client.HSet(ctx, key, k, v).Err()
		if err != nil {
			panic(err)
		}
	}
}

func (s *CacheService) GetMap(key string) (val map[string]string) {
	val, err := Client.HGetAll(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}

// SAdd 集合操作
func (s *CacheService) SAdd(key string, value interface{}) int64 {
	res, err := Client.SAdd(ctx, key, value).Result()
	if err != nil {
		panic(err)
	}
	return res
}

func (s *CacheService) SMembers(key string) (val []string) {
	val, err := Client.SMembers(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}

func (s *CacheService) SRem(key string, members ...interface{}) {
	err := Client.SRem(ctx, key, members).Err()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}

func (s *CacheService) SPop(key string) (val string) {
	val, err := Client.SPop(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}

func (s *CacheService) SPopN(key string, count int64) (val []string) {
	val, err := Client.SPopN(ctx, key, count).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}

func (s *CacheService) SScan(key string, cursor uint64, match string, count int64) (val []string) {
	val, _, err := Client.SScan(ctx, key, cursor, match, count).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}
