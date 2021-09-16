package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 创建全局的连接池指针
var pool *redis.Pool

// 当启动程序时就初始化连接池
func init() {
	// 创建redis 连接池
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大的空闲链接数
		MaxActive:   0,   //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: 100, // 最大空闲时间 // 如果没有100秒内没有使用链接则自动放回连接池
		Dial: func() (redis.Conn, error) {
			// 初始化链接的代码，链接哪一个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}

}

func main() {
	// 从pool中取出一个链接
	conn := pool.Get()
	defer conn.Close()

	// 向redis中存入数据
	_, err := conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("connet err=", err)
		return
	}

	// 从redis中取出数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	// str := string(r.())
	fmt.Println("r=", r)
	//如果要从pool中取出链接，那么一定要保证连接池是没有关闭的
	pool.Close()
	conn2 := pool.Get()
	// error :connet err= redigo: get on closed pool
	_, err = conn2.Do("Set", "name", "tom---")
	if err != nil {
		fmt.Println("connet err=", err)
		return
	}

	// 从redis中取出数据
	r1, err := redis.String(conn2.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	fmt.Println("r1=", r1)
	fmt.Println(conn2)

}
