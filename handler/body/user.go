package body

type UserLoginReq struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Type     int    `json:"type"`
}

type RegisterReq struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Type     int    `json:"type"`
}

type UpdateUserReq struct {
	ID          string `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
