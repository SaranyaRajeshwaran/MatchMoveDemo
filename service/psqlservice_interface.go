package service

import (
	model "MatchMove/model"
	"time"
)

//IPSQLService ...
type IPSQLService interface {
	CheckUserInformation(login *model.Login) (*model.LoginInfoResponse, error)
	CreateToken(username string, validduration time.Duration) (string, error)
	ValidateToken(token string) (*model.Token, error)
	DisableToken(token string, username string) (bool, error)
	NewJWT(key string) (*model.JWTCreator, error)
}
