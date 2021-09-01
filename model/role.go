package model

type Role struct {
	Id       int64  `json:"id"`
	Account  int64  `json:"account"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
