package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtNumber uint32

// New 产生token
func New(JwtKey string, opts ...Option) (string, error) {
	options := DefaultOptions()
	for _, o := range opts {
		o(&options)
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, options)
	token, err := claims.SignedString([]byte(JwtKey))
	return token, err

}

// Parse 分析token
func Parse(JwtKey, token string) (*Options, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Options{}, func(token *jwt.Token) (interface{}, error) {
		//返回秘钥
		return []byte(JwtKey), nil
	})

	if tokenClaims != nil {
		//检查token的有效性tokenClaims.Valid,类型断言
		if claims, ok := tokenClaims.Claims.(*Options); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
