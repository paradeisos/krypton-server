package controllers

import (
	"krypton-server/utils"
	"testing"

	"github.com/golib/assert"
	"github.com/satori/go.uuid"
	"net/http"
)

func TestUser_Register(t *testing.T) {
	assertion := assert.New(t)

	params := &UserRegisterParams{
		Username: uuid.NewV4().String(),
		Email:    uuid.NewV4().String() + "@test.com",
		Password: uuid.NewV4().String(),
	}
	r := utils.RequestJson(http.MethodPost, "/user", params)
	response := mockRequest(r)

	assertion.Equal(200, response.Status)
}
