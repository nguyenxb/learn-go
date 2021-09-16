package Sprocess

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/item/chatroom/common/message"
	"gocode/item/chatroom/server/Smodel"
	"gocode/item/chatroom/server/Sutils"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) LoginCheck(conn net.Conn, strOfLoginMes string) (err error) {
	// 将登录消息反序列化
	fmt.Println("data=", strOfLoginMes)
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(strOfLoginMes), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal，", err)
		return
	}

	// 从数据库中根据用户id读取数据
	user, err := Smodel.MyuserDao.LoginCheck(loginMes.UserId)
	if err != message.ErrorUserExists {
		return
	}
	if user == nil {
		return message.ErrorUserNotexists
	}
	var loginResMes message.LoginResMes
	// 这个时候user已经获取到了
	if loginMes.UserId != user.UserId || user.UserPwd != loginMes.UserPwd {
		loginResMes.Code = 400
	} else {
		loginResMes.Code = 200
	}
	// 将验证结果返回给客户端

	var mes message.Message
	mes.Type = message.LoginResMesType

	// 将loginResMes 进行序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		err = message.ErrorSerializationFailure
		return
	}
	mes.Data = string(data)
	// 将mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		err = message.ErrorSerializationFailure
		return
	}
	// 获取 数据包的长度的字节
	var pkgLen uint32
	var buf [4]byte
	dataPkgLen := len(data)
	pkgLen = uint32(dataPkgLen)
	// 将pkhLen 转成 二进制存到数组切片buf中
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	// 将数据包的长度字节发送给客户端
	conn.Write(buf[:4])
	n, err := conn.Write(data)
	if n != dataPkgLen {
		fmt.Println("conn.Write err", err)
		return
	}
	fmt.Println("服务端发回指令成功", pkgLen)
	return
}

// 将对数据进行注册检查
func (this *UserProcess) RegisterCheck(conn net.Conn, strOfRegisterMes string) (err error) {
	// 先将数据进行反序列化
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(strOfRegisterMes), &registerMes)
	if err != nil {
		err = message.ErrorDeserializationFailure
		return
	}
	// 信息存入数据库中
	err = Smodel.MyuserDao.RegisterCheck(registerMes.UserId, strOfRegisterMes)

	var mes message.Message
	mes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	if err != nil {
		registerResMes.Code = 300
		registerResMes.Error = err.Error()
	} else {
		registerResMes.Code = 200
	}

	// 将registerResMes 进行序列化
	data, _ := json.Marshal(registerResMes)
	mes.Data = string(data)
	util := Sutils.Utils{}
	util.WriteDataToClient(conn, &mes)

	// 返回
	return
}
