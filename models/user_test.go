package models

import (
	"testing"

	"github.com/golib/assert"
	uuid "github.com/satori/go.uuid"
)

func Test_NewUserModel(t *testing.T) {
	assertion := assert.New(t)

	var (
		username = uuid.NewV4().String()
		email    = "tes1@test.com"
		password = uuid.NewV4().String()
		desc     = uuid.NewV4().String()
	)

	user := User.NewUserModel(username, email, password, desc)
	assertion.Equal(username, user.Username)
	assertion.Equal(email, user.Email)
	assertion.Equal(password, user.Password)
	assertion.Equal(UserStatusInactive, user.Status)
	assertion.Equal(desc, user.Description)
}

func Test_User_FindByEmail(t *testing.T) {
	assertion := assert.New(t)

	var (
		username = uuid.NewV4().String()
		email    = "test2@test.com"
		password = uuid.NewV4().String()
		desc     = uuid.NewV4().String()
	)

	user := User.NewUserModel(username, email, password, desc)
	err := user.Save()
	assertion.Nil(err)

	// duplicated save
	user.Save()
	assertion.Nil(err)

	userNew, err := User.FindByEmail(email)
	assertion.Nil(err)
	assertion.Equal(user.Id, userNew.Id)

}

func Test_User_FindByUsername(t *testing.T) {
	assertion := assert.New(t)

	var (
		username = uuid.NewV4().String()
		email    = "test3@test.com"
		password = uuid.NewV4().String()
		desc     = uuid.NewV4().String()
	)

	user := User.NewUserModel(username, email, password, desc)
	err := user.Save()
	assertion.Nil(err)

	userNew, err := User.FindByUsername(username)
	assertion.Nil(err)
	assertion.Equal(user.Id, userNew.Id)

}
