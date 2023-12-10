package jwt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Options 统一的鉴权jwt数据
type Options struct {
	UID  uint32 `json:"uid"`  //鉴权用户的UID
	Acct string `json:"acct"` //登录用户名或者鉴权 Identifier
	Ut   string `json:"ut"`   //用户类型。//需要区分等级的时候才需要用到
	Role string `json:"role"` //身份角色。"身份1,状态1|身份2,状态2"
	Pt   string `json:"pt"`   //patch
	//jwtClaims jwt.StandardClaims
	jwt.RegisteredClaims
}

// Option Option
type Option func(o *Options)

// UID set uid
func UID(uid uint32) Option {
	return func(o *Options) {
		o.UID = uid
	}
}
func (o JwtNumber) String() string {
	return string(strconv.Itoa(int(o)))
}
func (o JwtNumber) Uint32() uint32 {
	return uint32(o)
}

// Account account
func Acct(account string) Option {
	return func(o *Options) {
		o.Acct = account
	}
}

// OwnType usertype
func Ut(ut string) Option {
	return func(o *Options) {
		o.Ut = ut
	}
}

// OwnType usertype
func Pt(pt string) Option {
	return func(o *Options) {
		o.Pt = pt
	}
}

// Role roles
func Role(role string) Option {
	return func(o *Options) {
		o.Role = role
	}
}

// Expire expire
func Expire(expire int64) Option {
	return func(o *Options) {
		o.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expire)))
		//o.ExpiresAt = expire
	}
}

// Issuer Issuer
func Issuer(issuer string) Option {
	return func(o *Options) {
		o.Issuer = issuer
	}
}

// func App(app string) Option {
// 	return func(o *Options) {
// 		o.App = app
// 	}
// }

// DefaultOptions default options
func DefaultOptions() Options {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 30 * time.Hour)

	o := Options{
		Role: "",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "go-crypto",
		},
	}
	return o
}
