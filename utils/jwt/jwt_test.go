package jwt

import (
	"testing"

	"github.com/golib/assert"
)

var (
	testUserID   = "testUserID"
	testUserName = "testUserName"
	testTime     = int64(86400)
)

func Test_Token(t *testing.T) {
	//Generate Token
	token := GenToken(testUserID, testUserName, testTime)
	assertion := assert.New(t)
	assertion.NotNil(token)

	// Check Token
	assertion.Equal(true, CheckToken(token))
}
