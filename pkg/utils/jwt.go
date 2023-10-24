package utils

import (
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaim struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenToken(userId int64) (string, error) {
	nowTime := jwt.NewNumericDate(time.Now())
	expireTime := jwt.NewNumericDate(nowTime.Add(time.Hour * 24))
	claims := UserClaim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expireTime,
			IssuedAt:  nowTime,
			Issuer:    "tiktok_jwt",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(constants.JwtSecrete))
	return token, err
}
