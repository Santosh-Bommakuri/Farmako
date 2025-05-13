package cache

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
    r := redis.NewClient(&redis.Options{
        Addr:     os.Getenv("add"), 
        Password: os.Getenv(""), 
        DB:       0,
    })
    if err := r.Ping(Ctx).Err(); err != nil {
        log.Fatalf("failed to connect to Redis: %v", err)
    }
    Redis = r
}


func SaveCoupon(key string, data []byte) {
    Redis.Set(Ctx, key, data, time.Hour)
}


func GetCoupon(key string) (string, error) {
    return Redis.Get(Ctx, key).Result()
}
