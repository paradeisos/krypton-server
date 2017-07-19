package controllers

import (
	"krypton-server/models"
	"krypton-server/options"
	"krypton-server/utils"
	"testing"
	"time"

	"net/http"

	"github.com/golib/assert"
	"gopkg.in/mgo.v2/bson"
)

func Test_Tomato(t *testing.T) {
	uid := bson.NewObjectId().Hex()
	opts := &options.CreateTomatoOpts{
		Uid:      uid,
		Start:    time.Now(),
		End:      time.Now().Add(25 * time.Minute),
		Desc:     "create first one tomato!",
		Finished: true,
	}

	// test for creating tomato
	r := utils.RequestJson(http.MethodPost, "/tomato", opts)
	response := mockRequest(r)
	assertion := assert.New(t)
	assertion.Equal(200, response.Status)

	tomatoes, err := models.Tomato.AllByUid(uid)
	assertion.Nil(err)
	assertion.Equal(len(tomatoes), 1)

	//test for getting tomato
	r = utils.RequestJson(http.MethodGet, "/tomato?id="+tomatoes[0].Id.Hex(), nil)
	response = mockRequest(r)
	assertion.Equal(200, response.Status)
	dataMap := response.Data.(map[string]interface{})
	assertion.Equal(dataMap["uid"], uid)

	// test for updating tomato
	newDesc := "new description"
	var updataOpts = options.UpdateTomatoOpts{
		ID:   tomatoes[0].Id.Hex(),
		Desc: newDesc,
	}
	r = utils.RequestJson(http.MethodPut, "/tomato", updataOpts)
	response = mockRequest(r)
	assertion.Equal(200, response.Status)
}
