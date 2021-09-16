package Smodel

import (
	"encoding/json"
	"fmt"
	"gocode/item/chatroom/common/message"

	"github.com/garyburd/redigo/redis"
)

type UserDao struct {
	Pool *redis.Pool
}

// 初始化线程池
var MyuserDao UserDao

func (this *UserDao) GetUserById(conn redis.Conn, userId int) (user *message.User, err error) {
	// 从数据库中获取数据
	// 查询到就返回空，没有查询到就返回错误
	data, _ := conn.Do("Hget", "users", userId)
	if data == nil { //表示在 users 哈希中，没有找到对应id
		err = message.ErrorUserNotexists
		return
	}
	if data != nil {
		// 如果存在该用户，
		//先将该用户进行类型转化转成切片
		err = message.ErrorUserExists
		buf, ok := data.([]byte)
		if !ok {
			err = message.ErrorTypeAssertionFailure
		}
		// 将该用户进行反序列化,并为user进行赋值
		json.Unmarshal(buf, &user)
	}
	return
}

func (this *UserDao) LoginCheck(userId int) (user *message.User, err error) {
	// 从线程池中获取数据库的链接
	conn := this.Pool.Get()
	user, err = this.GetUserById(conn, userId)
	if err != nil {
		return
	}
	return

}
func (this *UserDao) RegisterCheck(userId int, strOfRegisterMes string) (err error) {
	conn := this.Pool.Get()
	_, err = this.GetUserById(conn, userId)
	fmt.Println("GetUserById111=", err)
	if err == message.ErrorUserNotexists {
		// 将数据存入数据库中
		n, _ := conn.Do("HSet", "users", userId, strOfRegisterMes)
		fmt.Println("n=", n)
		fmt.Println("conn Do Hset err=", err)

		return nil
	}

	return
}
