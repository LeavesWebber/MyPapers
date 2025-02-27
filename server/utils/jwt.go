package utils

import (
	"errors"
	"server/global"
	"server/model/tables"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// BaseClaims 基本信息
type BaseClaims struct {
	UUID        int64
	ID          uint
	Username    string
	FirstName   string
	LastName    string
	AuthorityId uint
}

// CustomClaims 自定义JWT声明
type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

var mySecret = []byte("mypapers-secret-key")

// GetUserID 从上下文中获取用户ID
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
	c := CustomClaims{
		BaseClaims: BaseClaims{
			UUID:        user.UUID,
			ID:          user.ID,
			Username:    user.Username,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			AuthorityId: user.AuthorityId,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(global.MPS_CONFIG.JWT.ExpiresTime) * time.Hour).Unix(),
			Issuer:    "mypapers",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	var mc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
