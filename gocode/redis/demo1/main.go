package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	// 操作哈希数据类型
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("链接数据库失败，%v\n", err)
		return
	}
	defer conn.Close()

	// 向redis 存入数据 一个一个存和取
	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Printf("set err=", err)
		return
	}
	_, err = conn.Do("HSet", "user01", "name", "张三")
	if err != nil {
		fmt.Printf("set err=", err)
		return
	}
	fmt.Println("hash 存入数据成功")

	//取出数据
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("HGet err=", err)
		return
	}

	r2, err := redis.String(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("HGet err=", err)
		return
	}
	fmt.Printf("r1 =%v, r1 type =%T\n", r1, r1)
	fmt.Printf("r2 =%v, r2 type =%T\n", r2, r2)

}
