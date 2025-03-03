package utils

import (
	"errors"
	"server/global"
	"server/model/tables"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var mySecret = []byte("夏天夏天悄悄过去")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中

// CustomClaims 拿去生成token的信息
type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

// BaseClaims 基本信息
type BaseClaims struct {
	UUID        int64
	ID          uint
	Username    string
	FirstName   string
	LastName    string
	AuthorityId uint
}

func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		return 0
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.ID
	}
}

// GenToken 生成JWT
func GenToken(user tables.User) (string, error) {
	// 创建一个我们自己的声明的数据
	c := CustomClaims{
		BaseClaims{
			user.UUID,
			user.ID,
			user.Username,
			user.FirstName,
			user.LastName,
			user.AuthorityId,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(global.MPS_CONFIG.JWT.ExpiresTime) * time.Hour).Unix(), // 过期时间
			Issuer:    "bluebell",                                                                          // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret) // 加了自己写的夏天别人不知道
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var mc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
