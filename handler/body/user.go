package body

type UserLoginReq struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Type     int    `json:"type"`
}

type Res struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type RegisterReq struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Type     int    `json:"type"`
}
