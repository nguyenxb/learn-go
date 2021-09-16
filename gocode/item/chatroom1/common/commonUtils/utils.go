package commonUtils

// import (
// 	"encoding/binary"
// 	"encoding/json"
// 	"fmt"
// 	"gocode/item/chatroom/common/message"
// 	"net"
// )

// type Utils struct {
// 	Conn net.Conn
// }

// func (this Utils) WriteDataToServer(mes *message.Message) {
// 	fmt.Println("conn=", this.Conn)
// 	// 将mes 序列化
// 	data, err := json.Marshal(mes)
// 	if err != nil {
// 		err = message.ErrorSerializationFailure
// 	}
// 	// 获取序列化后的数据长度
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	// 向服务端发送数据长度
// 	_, err = this.Conn.Write(buf[:4])
// 	if err != nil {
// 		return
// 	}

// 	// 向服务端发送数据
// 	n, err := this.Conn.Write(data)
// 	if err != nil {
// 		fmt.Println("conn.Write err", err)
// 		return
// 	}
// 	fmt.Println("接收到的数据包长度，", n)
// 	fmt.Println("发送成功")

// }
// func (this *Utils) ReadDataFromServer() (resMes message.Message, err error) {
// 	var resBuf [4096]byte

// 	// 获取服务端发回的数据包长度字节
// 	_, err = this.Conn.Read(resBuf[:4])
// 	if err != nil {
// 		fmt.Println("conn.Read err", err)
// 		return
// 	}
// 	// 获取数据包长度
// 	resPkgLen := binary.BigEndian.Uint32(resBuf[:4])

// 	n, err := this.Conn.Read(resBuf[:resPkgLen])
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
