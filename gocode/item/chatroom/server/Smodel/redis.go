package Smodel

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

func InitPool(maxIdel, maxActive int, time time.Duration) {
	pool := redis.Pool{
		MaxIdle:     maxIdel,
		MaxActive:   maxActive,
		IdleTimeout: time,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	// 将数据库操作连接池初始化
	MyuserDao.Pool = &pool
}
