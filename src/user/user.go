package user

type TbUser struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

