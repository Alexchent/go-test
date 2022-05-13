package Cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var Client *redis.Client
var ctx = context.Background()

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	//if err != nil {
	//    panic(err)
	//}
	//
	//Client := redis.NewClient(opt)
}

func Test() {
	err := Client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := Client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := Client.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func Get(key string) (val string) {
	val, err := Client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}

func Set(key string, value interface{}, expiration time.Duration) {
	err := Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
	return
}

// SAdd 集合操作
func SAdd(key string, value interface{}) int64 {
	res, err := Client.SAdd(ctx, key, value).Result()
	if err != nil {
		panic(err)
	}
	return res
}

func SMembers(key string) (val []string) {
	val, err := Client.SMembers(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}

func SRem(key string, members ...interface{}) {
	err := Client.SRem(ctx, key, members).Err()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}

func SPop(key string) (val string) {
	val, err := Client.SPop(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}

func SPopN(key string, count int64) (val []string) {
	val, err := Client.SPopN(ctx, key, count).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}

func SScan(key string, cursor uint64, match string, count int64) (val []string) {
	val, _, err := Client.SScan(ctx, key, cursor, match, count).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	}
	return
}
