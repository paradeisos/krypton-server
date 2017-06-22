package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	*jwt.StandardClaims
	UserID   string
	UserName string
}

const (
	TokenIssuer = "kryptoner"
)

var (
	key []byte = []byte("kryptoner@krypton.com")
)

func GenToken(userID, userName string, dt int64) string {
	claims := JwtClaims{
		&jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + dt,
			Issuer:    TokenIssuer,
		},
		userID,
		userName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		//TODO
	}
	return ss
}

func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		//TODO
		return false
	}

	return true
}
