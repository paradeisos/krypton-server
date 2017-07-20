package controllers

import (
	"encoding/json"
	"krypton-server/errors"
	"krypton-server/models"
	"krypton-server/options"

	"net/http"

	"github.com/astaxie/beego"
)

type Todo struct {
	beego.Controller
}

func (c *Todo) Post() {
	var opts *options.CreateTodoOpts
	resp := &Response{}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &opts)
	if err != nil {
		beego.Error(err)
	}

	// TODO uid should from session?
	uid := ""

	todo := models.Todo.NewModel(uid, opts.Title, opts.Content, opts.Due)

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
