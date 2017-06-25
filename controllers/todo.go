package controllers

import (
	"encoding/json"
	"krypton-server/errors"
	"krypton-server/models"
	"time"

	"net/http"

	"github.com/astaxie/beego"
)

type Todo struct {
	beego.Controller
}

type CreateTodoParams struct {
	Uid     string    `json:"uid"`
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

func (c *Todo) Get() {
	id := c.GetString("id")
	if id == "" {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	todo, err := models.Todo.Find(id)
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	c.Data["json"] = Newresponse(http.StatusOK, "", todo)
	c.ServeJSON()
}
