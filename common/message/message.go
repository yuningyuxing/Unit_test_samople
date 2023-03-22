package message

const (
	//这个表示登陆类信息
	LoginMesType = "LoginMes"
	//这个表示登陆后返回的登陆情况的信息
	LoginReMesType = "LoginResMes"
	//表示注册类信息
	RegisterMesType = "RegisterMes"
	//表示该信息为返回注册类信息情况的信息
	RegisterResMesType = "RegisterResMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息本身
}
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}
type LoginResMes struct {
	Code  int    `json:"code"'` //返回状态码 500表示用户未错误，200表示登陆成功
	Error string `json:"error"` //返回错误信息
}

// 定义一下注册的消息类型
type RegisterMes struct {
	User User `json:"user"`
}

// 定义一下注册后返回的信息的结构体
type RegisterResMes struct {
	Code  int    `json:"code"` //状态码400表示该用户已经存在  200表示注册成功
	Error string `json:"error"`
}
