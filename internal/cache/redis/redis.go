package redis

import (
	"context"
	"gin-frame-base/internal/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:           global.Config.Redis.Host + ":" + global.Config.Redis.Port,
		Username:       global.Config.Redis.Username,
		Password:       global.Config.Redis.Password,
		DB:             global.Config.Redis.Db,
		IdentitySuffix: global.Config.Redis.Prefix,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
