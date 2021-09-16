package Sutils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/item/chatroom/common/message"
	"net"
)

type Utils struct {
}

func (this *Utils) WriteDataToClient(conn net.Conn, mes *message.Message) {

	// 将mes 序列化
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}
	// 获取序列化后的数据长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 向服务端发送数据长度
	_, err = conn.Write(buf[:4])
	if err != nil {
		return
	}

	// 向服务端发送数据
	n, err := conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}
	fmt.Println("接收到的数据包长度，", n)
	fmt.Println("发送成功")

}
func (this *Utils) ReadDataFromClient(conn net.Conn) (resMes message.Message, err error) {
	var resBuf [4096]byte
	// 获取服务端发给的数据包长度字节
	_, err = conn.Read(resBuf[:4])
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}
	fmt.Println("数据包长度读取成功", resBuf[:4])
	// 获取数据包长度
	resPkgLen := binary.BigEndian.Uint32(resBuf[:4])
	fmt.Println("数据包长度resPkglen", resPkgLen)
	// 读取来自客户端得的数据包
	n, err := conn.Read(resBuf[:resPkgLen])
	if err != nil {
		fmt.Println("conn.Read err11=", err)
		return
	}
	if int(resPkgLen) != n {
		fmt.Println("接收来自客户端的数据不一致")
		return
	}

	// 将数据包进行发序列化resMes
	err = json.Unmarshal(resBuf[:resPkgLen], &resMes)
	if err != nil {
		fmt.Println("json.Unmarshal err111", err)
		return
	}
	return

}
