package model

type Login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginInfoResponse struct {
	UserName    string `json:"userName"`
	IsAdmin     bool   `json:"isAdmin"`
	IsValidUser bool   `json:"isValidUser"`
	Token       string `json:"token"`
}
