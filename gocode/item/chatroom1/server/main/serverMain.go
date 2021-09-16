// server
package main

import (
	"fmt"
	"gocode/item/chatroom/common/message"
	"gocode/item/chatroom/server/Smodel"
	"gocode/item/chatroom/server/Sprocess"
	"gocode/item/chatroom/server/Sutils"
	"net"
	"time"
)

func CreatProcess(conn net.Conn) (err error) {
	// 延时关闭连接
	defer conn.Close()
	fmt.Println("客户端连接。。")
	util := &Sutils.Utils{}
	// // 消息结构体
	// var mes message.Message
	// // 接收数据包的缓冲区
	// var buf [4096]byte
	// // 接收数据包的长度
	// _, err := conn.Read(buf[0:4])
	// if err != nil {
	// 	fmt.Println("conn.Read err1=", err)
	// 	return
	// }
	// var pkgLen uint32
	// // 读取数据包的长度
	// pkgLen = binary.BigEndian.Uint32(buf[:4])
	// var dataPkgLen int
	// dataPkgLen = int(pkgLen)
	// // 接收数据包
	// n, err := conn.Read(buf[:dataPkgLen])
	// if err != nil {
	// 	fmt.Println("conn.Read err==", err)
	// 	return
	// }
	// if n != dataPkgLen {
	// 	fmt.Println("数据丢包了")
	// 	return
	// }
	// // 将数据包反序列化得到 Message
	// err = json.Unmarshal(buf[:pkgLen], &mes)
	// if err != nil {
	// 	fmt.Println("json.Unmarshal err", err)
	// 	return
	// }

	mes, err := util.ReadDataFromClient(conn)
	if err != nil {
		err = message.ErrorReceiveData
		return
	}
	// fmt.Println("messsss=", mes)
	// fmt.Println(mes.Data) // "{\"userId\":200,\"userPwd\":\"123456\",\"userName\":\"\"}"
	switch mes.Type {
	case message.LoginMesType:
		up := &Sprocess.UserProcess{}
		err = up.LoginCheck(conn, mes.Data)

	case message.RegisterMesType:
		up := &Sprocess.UserProcess{}
		err = up.RegisterCheck(conn, mes.Data)

	default:
		return message.ErrorUnknown
	}
	return
}

// 初始化线程池
func init() {
	Smodel.InitPool(16, 0, 300*time.Second)
}
func main() {

	// 1 监听端口 9999
	fmt.Println("服务器在端口 9999 开始监听")
	listen, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}

	// 2 等待获取来自于客户端的连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err", err)
			return
		}
		// 每连接到一个客户端就启动一个协程为其服务
		go CreatProcess(conn)

	}

}
