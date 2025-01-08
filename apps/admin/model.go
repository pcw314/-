package admin

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (LoginRequest) TableName() string {
	return "store_menus_authority"
}
