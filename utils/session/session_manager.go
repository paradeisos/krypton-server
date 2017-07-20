package session

import (
	"strings"
	"time"
	"math/rand"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenHeader = "Bearer "
)

type SessionManger struct {
	Secret      string
	Issuer      string
	ExpiredTime int64 // expired time
}

func (s *SessionManger) NewSession(userID, username string) *Session {
	return &Session{
		&jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + s.ExpiredTime,
			Issuer:    s.Issuer,
		},
		rand.Float64(),
		userID,
		username,
		s.Secret,
	}
}

func (s *SessionManger) NewSessionByToken(token string) (session *Session, err error) {
	token = s.removeTokenHeader(token)
	token, userID, err := s.parseToken(token)
	if err != nil {
		return nil, err
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Session{}, func(*jwt.Token) (interface{}, error) {
		key := s.key(userID)
		return key, nil
	})

	if err != nil {
		return session, err
	}

	return jwtToken.Claims.(*Session), nil
}

func (s *SessionManger) key(userID string) []byte {
	return key(userID, s.Secret)
}

func (s *SessionManger) removeTokenHeader(token string) string {
	return strings.TrimPrefix(token, TokenHeader)
}

func (s *SessionManger) parseToken(t string) (token, userId string, err error) {
	arr := strings.Split(t, "|")
	if len(arr) != 2 {
		return t, "", errors.New("session token has no user info")
	}

	return arr[0], arr[1], nil
}
