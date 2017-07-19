package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"krypton-server/models"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	absPath, err := filepath.Abs("")
	if err != nil {
		beego.Error(err.Error())
		return
	}

	apppath := filepath.Dir(absPath)

	beego.TestBeegoInit(apppath)
}

func mockRequest(request *http.Request) (response *Response) {
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, request)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		beego.Error(string(w.Body.Bytes()))
		beego.Error(err)
	}
	return
}

func TestMain(m *testing.M) {

	InitEnv()
	// init models
	models.InitMongo()

	// running all test suites
	result := m.Run()

	// output test result
	os.Exit(result)
}
