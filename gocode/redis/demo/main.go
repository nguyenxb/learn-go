package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//  通过 go向redis 写入数据和读取数据
	// 1. 连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("链接失败", err)
		return
	}
	defer conn.Close()

	// fmt.Println("conn succ...", conn)

	//  通过go 向redis 写入数据 string [key-val]
	r1, err := conn.Do("Set", "name", "jjjjj")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	fmt.Printf("r1 =%v, r1 type=%T\n", r1, r1)
	// 3 通过 go向redis 读取数据 string[key-value]
	r2, err := conn.Do("Get", "name")
	r3, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	fmt.Printf("r2 = %v,r2 type = %T\n", r2, r2)
	fmt.Printf("r2 =%s\n", r2)
	fmt.Printf("r3 = %v,r3 type = %T\n", r3, r3)
	fmt.Printf("r3 =%s\n", r3)
	fmt.Println("操作ok")
}
