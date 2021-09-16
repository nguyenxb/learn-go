package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	// 链接redis
	conn, err := redis.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()

	var name string
	var age int
	var key string
	fmt.Println("请输入key")
	fmt.Scanln(&key)
	fmt.Println("请输入name")
	fmt.Scanln(&name)
	fmt.Println("请输入age")
	fmt.Scanln(&age)
	fmt.Println(key)
	fmt.Println(name)
	fmt.Println(age)
	// 存多个数据，取多个数据
	// 向redis写入数据 string [key-val]
	// _, err = conn.Do("HMSet", "user02", "name", "john", "age", 19)
	_, err = conn.Do("HMSet", key, "name", name, "age", age)
	if err != nil {
		fmt.Println("HMSet err", err)
		return
	}
	fmt.Println("存入成功")

	//  从redis 取出数据
	r, err := redis.Strings(conn.Do("HMGet", key, "name", "age"))
	// r, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("HMGet err", err)
		return
	}
	fmt.Printf("r =%v,r type=%T\n", r, r)
	for i, v := range r {
		fmt.Printf("r[%v]=%v\n", i, v)
	}

}
