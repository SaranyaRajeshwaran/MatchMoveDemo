package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Token struct {
	Id          uuid.UUID `json:Id`
	UserName    string    `json:"userName"`
	IssuedTime  time.Time `json:"issuedTime"`
	ExpiredTime time.Time `json:"expiredTime"`
}

type ValidateToken struct {
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

type JWTCreator struct {
	SecretKey string `json:"secretkey"`
}

func (token *Token) Valid() error {
	if time.Now().After(token.ExpiredTime) {
		return errors.New("Token Expired")
	}
	return nil

}
