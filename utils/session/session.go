package session

import (
	"crypto/md5"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

type Session struct {
	*jwt.StandardClaims
	R        float64
	UserID   string
	UserName string
	Secret   string
}

func (s *Session) Token() (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, s)
	token, err = jwtToken.SignedString(s.key())
	if err != nil {
		return
	}

	return fmt.Sprintf("%s%s|%s", TokenHeader, token, s.UserID), nil
}

func (s *Session) key() []byte {
	return key(s.UserID, s.Secret)
}

func key(userID, secret string) []byte {
	s := fmt.Sprintf("kryptpon %s %s", userID, secret)
	return []byte(fmt.Sprintf("%x", md5.Sum([]byte(s))))
}
