package controllers

import (
	"encoding/json"
	"krypton-server/models"

	"net/http"

	"github.com/astaxie/beego"
)

type Tomato struct {
	beego.Controller
}

func (c *Tomato) Post() {
	var params *CreateTomatoParams
	resp := &Response{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if err != nil {
		beego.Error(err)
	}

	tomato := models.Tomato.NewTomatoModel(params.Uid, params.Start, params.End, params.Desc, params.Finished)

	if err := tomato.Save(); err != nil {
		beego.Error(err)
	}

	resp.Status = http.StatusOK
	c.Data["json"] = resp
	c.ServeJSON()
}
