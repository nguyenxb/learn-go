package Cprocess

import (
	"encoding/json"
	"fmt"
	"gocode/item/chatroom/client/Cutils"
	"gocode/item/chatroom/common/message"
	"net"
)

// 用于处理与用户相关的业务
type UserProcess struct {
	conn net.Conn
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// 1 连接到服务器
	this.conn, err = net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer this.conn.Close()
	// 2 将用户id 和 用户 密码序列化后赋值给mes
	var mes message.Message
	var loginMes message.LoginMes
	// 设置消息类型为登录类型
	mes.Type = message.LoginMesType

	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// 将loginMes 序列化
	data, _ := json.Marshal(loginMes)

	mes.Data = string(data)

	Cutils := &Cutils.Utils{}
	Cutils.WriteDataToServer(this.conn, &mes)
	var resLoginMes message.LoginResMes
	resMes, _ := Cutils.ReadDataFromServer(this.conn)
	err = json.Unmarshal([]byte(resMes.Data), &resLoginMes)
	// 输出返回的结果
	fmt.Println("code=", resLoginMes.Code)
	fmt.Println("error = ", resLoginMes.Error)
	if resLoginMes.Code == 200 {
		fmt.Println("登录成功")
		return
	} else if resLoginMes.Code == 400 {
		fmt.Println("用户未注册")
	}
	fmt.Println("")
	return
}

func (this *UserProcess) Register(userId int, userPwd, userName string) (err error) {
	// 链接到服务器
	this.conn, err = net.Dial("tcp", "localhost:9999")
	if err != nil {
		return
	}
	defer this.conn.Close()
	// 给registerMes 实例赋值

	registerMes := message.RegisterMes{
		User: message.User{
			UserId:   userId,
			UserPwd:  userPwd,
			UserName: userName,
		}}
	// 将register序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	mes := &message.Message{
		Type: message.RegisterMesType,
		Data: string(data),
	}
	// 创建一个工具实例
	util := Cutils.Utils{}
	// 向服务端发送数据
	err = util.WriteDataToServer(this.conn, mes)
	if err != nil {
		fmt.Println("发送数据失败")
	}

	// 接收服务端返回的结果
	resMes, err := util.ReadDataFromServer(this.conn)
	if err != nil {
		fmt.Println("从服务端接收数据失败")
		return
	}
	// 将发回的数据进行反序列化成resRegisterMes
	var resRegisterMes message.RegisterResMes
	err = json.Unmarshal([]byte(resMes.Data), &resRegisterMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	//获取并判断结果
	if resRegisterMes.Code == 200 {
		fmt.Println("注册成功", resRegisterMes.Code)
		return
	}
	fmt.Println("注册失败")
	return
}
