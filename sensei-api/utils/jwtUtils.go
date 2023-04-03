package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"vilacorp.com/sensei/config"
)

type JWTClaim struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Username string    `json:"username"`
	Mail     string    `json:"mail"`
	Csrf     string    `json:"csrf"`
	jwt.StandardClaims
}

func GenerateJWT(id string, mail string, username string, csrf string) (tokenString string, err error) {
	expirationTime := time.Now().Add(6 * time.Hour)
	uuidId, _ := uuid.FromString(id)
	claims := &JWTClaim{
		Id:       uuidId,
		Mail:     mail,
		Username: username,
		Csrf:     csrf,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(config.AppConfig.JwtKey))
	return
}
func ValidateToken(signedToken string) (tokenClaims JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return *claims, nil
}
