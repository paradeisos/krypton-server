package models

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/astaxie/beego"
)

func TestMain(m *testing.M) {
	absPath, err := filepath.Abs("")
	if err != nil {
		beego.Error(err.Error())
		return
	}

	apppath := filepath.Dir(absPath)
	beego.TestBeegoInit(apppath)

	InitMongo()

	code := m.Run()

	DropDatabase()

	os.Exit(code)
}

func DropDatabase() {
	Mongo().Session().DB(mongo.Database()).DropDatabase()
}
