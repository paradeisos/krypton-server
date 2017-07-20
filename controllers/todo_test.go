package controllers

import (
	"krypton-server/options"
	"krypton-server/utils"
	"net/http"
	"testing"
	"time"

	"github.com/golib/assert"
)

var (
	testTitle   = "testTitle"
	testContent = "testContent"
)

func Test_Toto(t *testing.T) {
	// creating todo item
	opts := &options.CreateTodoOpts{
		Title:   testTitle,
		Content: testContent,
		Due:     time.Now(),
	}

	r := utils.RequestJson(http.MethodPost, "/todo", opts)
	resp := mockRequest(r)
	assertion := assert.New(t)
	assertion.Equal(200, resp.Status)
}
