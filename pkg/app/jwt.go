package app

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"nickygin.com/global"
	"nickygin.com/pkg/util"
)

type Claims struct {
	Uid            string `json:"app_key"`
	NickName       string `json:"name"`
	*jwt.MapClaims `json:"app_claims"`
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(tokenClaims Claims) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	uid, error := util.AESEncrypt(tokenClaims.Uid, global.JWTSetting.Secret2)
	if error != nil {
		return tokenClaims.Uid, error
	}
	tokenClaims.Uid = uid
	tokenClaims.MapClaims = &jwt.MapClaims{"exp": expireTime.Unix(), "iat": global.JWTSetting.Issuer}
	newJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	s, err := newJwt.SignedString(GetJWTSecret())
	return s, err

}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
