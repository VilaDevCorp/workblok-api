package utils

import (
	"errors"
	"time"
	"workblok/conf"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTClaim struct {
	SessionId  uuid.UUID `json:"id"`
	UserAgent  string    `json:"userAgent"`
	UserId     uuid.UUID `json:"userId"`
	Username   string    `json:"username"`
	Mail       string    `json:"mail"`
	RememberMe bool      `json:"rememberMe"`
	jwt.RegisteredClaims
}

func GenerateJWT(
	userId uuid.UUID,
	mail string,
	username string,
	lifeTime time.Duration,
	rememberMe bool,
	sessionId uuid.UUID,
	userAgent string,
) (tokenString string, err error) {
	conf := conf.Get()
	issuer := conf.Dev.CookieHost
	if conf.Env == "prod" {
		issuer = conf.Prod.CookieHost
	}
	expirationTime := time.Now().Add(lifeTime)

	claims := &JWTClaim{
		SessionId:  sessionId,
		UserId:     userId,
		Mail:       mail,
		Username:   username,
		UserAgent:  userAgent,
		RememberMe: rememberMe,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: issuer,
			ExpiresAt: jwt.NewNumericDate(
				expirationTime,
			),
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	tokenString, err = token.SignedString(
		[]byte(conf.JwtKey),
	)
	return
}

func ValidateToken(
	signedToken string,
) (tokenClaims JWTClaim, err error) {
	conf := conf.Get()

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.JwtKey), nil
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
	if int64(
		claims.ExpiresAt.Compare(
			time.Now().Local(),
		),
	) <= 0 {
		err = errors.New("token expired")
		return
	}
	return *claims, nil
}
