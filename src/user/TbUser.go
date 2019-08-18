package user

type TbUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
