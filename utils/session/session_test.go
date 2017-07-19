package session

import (
	"testing"

	"github.com/golib/assert"
	"github.com/satori/go.uuid"
)

var (
	testUserID   = "testUserID"
	testUserName = "testUserName"
	testTime     = int64(86400)
)

func Test_Token(t *testing.T) {
	assertion := assert.New(t)
	manger := &SessionManger{
		Secret:      uuid.NewV4().String(),
		Issuer:      uuid.NewV4().String(),
		ExpiredTime: testTime,
	}

	session := manger.NewSession(testUserID, testUserName)
	token, err := session.Token()
	assertion.Nil(err)

	sessionRes, err := manger.NewSessionByToken(token)
	assertion.Nil(err)

	assertion.Equal(session.UserID, sessionRes.UserID)

}
