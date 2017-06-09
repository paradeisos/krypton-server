package controllers

import (
	"encoding/json"
	"krypton-server/models"
	"time"

	"net/http"

	"github.com/astaxie/beego"
)

type Todo struct {
	beego.Controller
}

type CreateTodoParams struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Due     time.Time `json:"due"`
}

func (c *Todo) Post() {
	var params *CreateTodoParams
	resp := &Response{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if err != nil {
		beego.Error(err)
	}

	// TODO uid should from session?
	uid := ""

	todo := models.Todo.NewModel(uid, params.Title, params.Content, params.Due)

	if err := todo.Save(); err != nil {
		beego.Error(err)
	}

	resp.Status = http.StatusOK
	c.Data["json"] = resp
	c.ServeJSON()
}
