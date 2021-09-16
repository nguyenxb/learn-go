package message

import "errors"

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

//根据业务逻辑需要，自定义一些错误.
var (
	ErrorUserExists             = errors.New("用户已经存在...")
	ErrorUserNotexists          = errors.New("用户不存在...")
	ErrorUserPwd                = errors.New("密码不正确...")
	ErrorSerializationFailure   = errors.New("序列化失败...")
	ErrorDeserializationFailure = errors.New("反序列化失败...")
	ErrorTypeAssertionFailure   = errors.New("类型断言错误...")
	ErrorSendData               = errors.New("数据发送错误...")
	ErrorReceiveData            = errors.New("数据接收错误...")
	ErrorSelect                 = errors.New("选项错误")
	ErrorDatabaseConnection     = errors.New("数据库连接错误...")
	ErrorDatabaseStorageFailure = errors.New("数据库存储失败...")
	ErrorDatabaseSearchFailure  = errors.New("数据库查询失败...")
	ErrorUnknown                = errors.New("未知错误...")
)

// 定义消息类型
type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` //数据的内容
}

type LoginMes struct {
	User
}
type LoginResMes struct {
	Code  int    `json:"code"`  // 返回状态码，200 表示登录成功，300 表示用户未注册
	Error string `json:"error"` // 返回错误信息
}

type User struct {
	UserId   int    `json:"userId"`   // 用户id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}
type RegisterMes struct {
	User
}

type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回状态码，200 表示注册成功，300 表示用户未注册
	Error string `json:"error"` // 返回错误信息
}
