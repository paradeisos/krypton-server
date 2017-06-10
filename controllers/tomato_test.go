package controllers

import (
	"krypton-server/utils"
	"testing"
	"time"

	"net/http"

	"github.com/golib/assert"
	"gopkg.in/mgo.v2/bson"
)

func Test_CreateTomato(t *testing.T) {
	assertion := assert.New(t)

	params := &CreateTomatoParams{
		Uid:      bson.NewObjectId().Hex(),
		Start:    time.Now(),
		End:      time.Now().Add(25 * time.Minute),
		Desc:     "create first one tomato!",
		Finished: true,
	}

	r := utils.RequestJson(http.MethodPost, "/tomato", params)
	response := mockRequest(r)

	assertion.Equal(200, response.Status)
}
