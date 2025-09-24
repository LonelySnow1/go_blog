package initialize

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"os"
	"server/global"
)

func ConnectRedis() redis.Client {
	redisCfg := global.Config.Redis

	client := redis.NewClient(&redis.Options{ // 新建redis链接
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	_, err := client.Ping().Result() // test联通
	if err != nil {
		global.Log.Error("Failed to connect to redis", zap.Error(err))
		os.Exit(1)
	}

	return *client
}
