package service

import (
	"MatchMove/database"
	model "MatchMove/model"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

//DatabaseService ...
type DatabaseService struct {
	DatabaseService database.DbIface
}

//NewDatabaseServicesInstance instantiates the struct
func NewDatabaseServicesInstance() *DatabaseService {
	return &DatabaseService{
		DatabaseService: database.DBNewHandler(),
	}
}

const Minkeysize = 24
const SecretKey = "secret"

func (service *DatabaseService) CheckUserInformation(login *model.Login) (*model.LoginInfoResponse, error) {

	defer service.DatabaseService.DbClose()
	log.Println(login.UserName)
	log.Println(login.Password)
	query := "select isadmin, user_name from Demo.UserInformation where user_name='" + login.UserName + "' AND password='" + login.Password + "'"
	log.Println(query)
	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	txResult := model.LoginInfoResponse{
		IsAdmin:     false,
		IsValidUser: false,
	}
	if err != nil {
		log.Println(err)
		return &txResult, err
	}

	if rowsAffected.Next() {
		var isadmin bool
		var userName string
		log.Println("rowsAffected")

		if err := rowsAffected.Scan(&isadmin, &userName); err != nil {
			log.Println("err")
			log.Println(err)
			return &txResult, nil

		}
		txResult.IsAdmin = isadmin
		txResult.UserName = userName
		txResult.IsValidUser = true
		var token string = ""
		if isadmin {
			token, err = service.CreateToken(userName, 3000000)
			log.Println("token")
			log.Println(token)
			if err != nil {
				token = ""
			}
			query := "insert into DEMO.token(user_name, token) values ('" + userName + "','" + token + "')"
			log.Println("log query")
			log.Println(query)
			_, err = service.DatabaseService.DbExecuteScalar(query)
			if err != nil {
				log.Println("txq")
				token = ""
			}
			txResult.Token = token
		}

		log.Println(txResult)

		return &txResult, nil

	}

	return &txResult, errors.New("No Rows Exists")
}

func (service *DatabaseService) ValidateToken(token string) (*model.Token, error) {
	keyFunc := func(tokem *jwt.Token) (interface{}, error) {

		_, ok := tokem.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SecretKey), nil
	}

	jwttoken, err := jwt.ParseWithClaims(token, &model.Token{}, keyFunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errors.New("Token Expired")) {
			return &model.Token{}, errors.New("token has expired")
		}
		log.Println("token1")
		log.Println(verr.Inner)
		return &model.Token{}, errors.New("token is Invalid")
	}
	payload, ok := jwttoken.Claims.(*model.Token)
	if !ok {
		log.Println("token2")
		return payload, errors.New("token is Invalid")
	}
	return payload, nil
}

func (service *DatabaseService) DisableToken(username string, token string) (bool, error) {

	query := "delete from DEMO.token where user_name ='" + username + "'AND token ='" + token + "'"
	log.Println("log query")
	log.Println(query)
	_, err := service.DatabaseService.DbExecuteScalar(query)
	if err != nil {
		log.Println("err")
		return false, nil

	}

	return true, nil
}

func (service *DatabaseService) CreateToken(username string, validduration time.Duration) (string, error) {
	log.Println("create token")
	defer service.DatabaseService.DbClose()
	log.Println("create token")
	log.Println(username)

	tokenID, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return "", err
	}

	token := &model.Token{
		Id:          tokenID,
		UserName:    username,
		IssuedTime:  time.Now(),
		ExpiredTime: time.Now().Add(validduration),
	}

	jwttoken := jwt.NewWithClaims(jwt.SigningMethodHS256, token)

	temp, err := service.NewJWT(SecretKey)

	if err != nil {
		log.Println("create token2")
		log.Println(err)
		return "", err
	}
	signedkey, err := jwttoken.SignedString([]byte(temp.SecretKey))

	if err != nil {
		log.Println(err)
		return "", err
	}
	return signedkey, nil

}

func (service *DatabaseService) NewJWT(key string) (*model.JWTCreator, error) {
	log.Println("len(key)")
	log.Println(len(key) < Minkeysize)
	if len(key) > Minkeysize {
		return nil, errors.New("Size Exceeds")
	}
	return &model.JWTCreator{key}, nil
}
