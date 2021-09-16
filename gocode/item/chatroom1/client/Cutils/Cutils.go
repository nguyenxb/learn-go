package Cutils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/item/chatroom/common/message"
	"net"
)

type Utils struct {
}

func (this *Utils) WriteDataToServer(conn net.Conn, mes *message.Message) (err error) {
	// defer conn.Close()
	// fmt.Println("客户端正在发送数据conn=", this.Conn)
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
	fmt.Println("客户端向服务端发送的数据包长度字节", buf)
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
	fmt.Println("发送到的数据包长度，", n)
	fmt.Println("发送成功")
	return
}
func (this *Utils) ReadDataFromServer(conn net.Conn) (resMes message.Message, err error) {
	// defer conn.Close()
	var resBuf [4096]byte

	// 获取服务端发回的数据包长度字节
	_, err = conn.Read(resBuf[:4])
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}
	fmt.Println("数据包返回成功", resBuf[:4])
	// 获取数据包长度
	resPkgLen := binary.BigEndian.Uint32(resBuf[:4])
	fmt.Println("resPkgLen=", resPkgLen)
	n, err := conn.Read(resBuf[:resPkgLen])
	if int(resPkgLen) != n || err != nil {
		fmt.Println("conn.Read err11=", err)
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
