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

// func (this *UserProcess) WriteDataToServer(mes message.Message) {
// 	// 将mes 序列化
// 	data, err := json.Marshal(mes)
// 	if err != nil {
// 		fmt.Println("json.Marshal err=", err)
// 	}
// 	// 获取序列化后的数据长度
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	// 向服务端发送数据长度
// 	_, err = this.conn.Write(buf[:4])
// 	if err != nil {
// 		return
// 	}

// 	// 向服务端发送数据
// 	n, err := this.conn.Write(data)
// 	if err != nil {
// 		fmt.Println("conn.Write err", err)
// 		return
// 	}
// 	fmt.Println("接收到的数据包长度，", n)
// 	fmt.Println("发送成功")

// }
// func (this *UserProcess) ReadDataFormServer() (resMes message.Message, err error) {
// 	var resBuf [4096]byte

// 	// 获取服务端发回的数据包长度字节
// 	_, err = this.conn.Read(resBuf[:4])
// 	if err != nil {
// 		fmt.Println("conn.Read err", err)
// 		return
// 	}
// 	// 获取数据包长度
// 	resPkgLen := binary.BigEndian.Uint32(resBuf[:4])

// 	n, err := this.conn.Read(resBuf[:resPkgLen])
// 	if int(resPkgLen) != n || err != nil {
// 		fmt.Println("conn.Read err11=", err)
// 		return
// 	}

// 	// 将数据包进行发序列化resMes
// 	err = json.Unmarshal(resBuf[:resPkgLen], &resMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal err111", err)
// 		return
// 	}
// 	return
// }

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// 1 连接到服务器
	this.conn, err = net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer this.conn.Close()

	// defer this.conn.Close()
	// fmt.Println("login conn=", this.conn)
	// fmt.Println("cutils conn=", Cutils.Conn)
	// 2 将用户id 和 用户 密码序列化后赋值给mes
	var mes message.Message
	var loginMes message.LoginMes
	// 设置消息类型为登录类型
	mes.Type = message.LoginMesType

	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	// // 将mes 序列化
	// data, err = json.Marshal(mes)
	// if err != nil {
	// 	fmt.Println("json.Marshal err=", err)
	// }

	// // 获取序列化后的数据长度
	// var pkgLen uint32
	// pkgLen = uint32(len(data))
	// var buf [4]byte
	// binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// // 向服务端发送数据长度
	// _, err = conn.Write(buf[:4])
	// if err != nil {
	// 	return
	// }

	// // 向服务端发送数据
	// n, err := conn.Write(data)
	// if err != nil {
	// 	fmt.Println("conn.Write err", err)
	// 	return
	// }
	// fmt.Println("接收到的数据包长度，", n)
	// fmt.Println("发送成功")
	// this.WriteDataToServer(mes)
	Cutils := &Cutils.Utils{}
	Cutils.WriteDataToServer(this.conn, &mes)
	//  接收服务端发回的结果
	// var resBuf [4096]byte

	// // 获取服务端发回的数据包长度字节
	// _, err = conn.Read(resBuf[:4])
	// if err != nil {
	// 	fmt.Println("conn.Read err", err)
	// 	return
	// }
	// // 获取数据包长度
	// resPkgLen := binary.BigEndian.Uint32(resBuf[:4])

	// n, err := conn.Read(resBuf[:resPkgLen])
	// if int(resPkgLen) != n || err != nil {
	// 	fmt.Println("conn.Read err11=", err)
	// 	return
	// }

	// // 将数据包进行发序列化resMes
	// var resMes message.Message
	// var resLoginMes message.LoginResMes
	// err = json.Unmarshal(resBuf[:resPkgLen], &resMes)
	// if err != nil {
	// 	fmt.Println("json.Unmarshal err111", err)
	// 	return
	// }
	// fmt.Println("111")
	// 获取到mes.Data的数据,将其反序列化为resLoginMes
	var resLoginMes message.LoginResMes
	resMes, err := Cutils.ReadDataFromServer(this.conn)
	if err != nil {
		fmt.Println("ReadDataFormServer err", err)
	}
	err = json.Unmarshal([]byte(resMes.Data), &resLoginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err222", err)
		return
	}
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
	// fmt.Println("resMes.Type=", resMes.Type)
	// fmt.Println("resMes.Data=", resMes.Data)
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
