package util

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	er "myshop/client/errors"
	"time"
)

var screte = []byte("test")

type UserClaims struct {
	Id string
	jwt.StandardClaims
}
func CreateToken(id string) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,UserClaims{
		Id:             id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	})
	tokenString,err := token.SignedString(screte)
	if err != nil {
		return "",err
	}
	return tokenString,nil
}
func ParseToken(tokenString string) (interface{},error) {
	token,err := jwt.ParseWithClaims(tokenString,&UserClaims{},func(*jwt.Token) (interface{}, error){
		return screte,nil
	})
	fmt.Printf("%#v\n",token)
	fmt.Println(err)
	if err != nil {
		return nil,er.TokenParseError
	}
	if claim,ok := token.Claims.(*UserClaims);ok && token.Valid {
		fmt.Println(claim.StandardClaims.ExpiresAt)
		fmt.Println(claim.Id)
		return claim,nil
	}
	return nil,er.ObjectCheckError
}