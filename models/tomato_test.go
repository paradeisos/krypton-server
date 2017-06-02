package models

import (
	"testing"
	"time"

	"github.com/golib/assert"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func Test_NewTomatoModel(t *testing.T) {
	assertion := assert.New(t)
	start := time.Now()
	end := start.Add(25 * time.Minute)
	uid := uuid.NewV4().String()

	tomato := Tomato.NewTomatoModel(uid, start, end, "test", true)
	assertion.Equal(start, tomato.Start)
	assertion.Equal(end, tomato.End)
	assertion.Equal(true, tomato.Finished)
}

func Test_FindByUid(t *testing.T) {
	assertion := assert.New(t)
	start := time.Now()
	end := start.Add(25 * time.Minute)
	uid := bson.NewObjectId().Hex()

	tomato := Tomato.NewTomatoModel(uid, start, end, "test", true)
	assertion.Nil(tomato.Save())

	// duplicated save
	assertion.Nil(tomato.Save())

	tomatoes, err := Tomato.AllByUid(uid)
	assertion.Nil(err)
	assertion.Equal(1, len(tomatoes))
	assertion.Equal(tomato.Id, tomatoes[0].Id)
}
