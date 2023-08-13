package redis

// redis 的配置及初始化

import (
	"02-GO_Web_CLI/settings"
	"fmt"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf(
				"%s:%d", cfg.Host, cfg.Port,
			),
			Password: cfg.Password, // no password set
			DB:       cfg.DB,       // use default DB
			PoolSize: cfg.PoolSize,
		},
	)

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
