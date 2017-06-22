package controllers

import (
	"encoding/json"
	"krypton-server/errors"
	"krypton-server/models"

	"net/http"

	"github.com/astaxie/beego"
)

type Tomato struct {
	beego.Controller
}

func (c *Tomato) Post() {
	var params *CreateTomatoParams

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if err != nil {
		beego.Error(err)
	}

	tomato := models.Tomato.NewTomatoModel(params.Uid, params.Start, params.End, params.Desc, params.Finished)

	if err := tomato.Save(); err != nil {
		beego.Error(err)
	}

	c.Data["json"] = Newresponse(http.StatusOK, "", nil)
	c.ServeJSON()
}

func (c *Tomato) Delete() {
	id := c.GetString("id")
	if id == "" {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	tomato, err := models.Tomato.Find(id)
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	if tomato.Delete() != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InternalError)
		c.ServeJSON()
		return
	}

	c.Data["json"] = Newresponse(http.StatusOK, "", nil)
	c.ServeJSON()
}

func (c *Tomato) Get() {
	id := c.GetString("id")
	if id == "" {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	tomato, err := models.Tomato.Find(id)
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	c.Data["json"] = Newresponse(http.StatusOK, "", tomato)
	c.ServeJSON()
}

func (c *Tomato) Put() {
	var params *UpdateTomatoParams

	if json.Unmarshal(c.Ctx.Input.RequestBody, &params) != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	tomato, err := models.Tomato.Find(params.ID)
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	tomato.Start = params.Start
	tomato.End = params.End
	tomato.Desc = params.Desc
	tomato.Finished = params.Finished

	if err := tomato.Update(); err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InternalError)
		c.ServeJSON()
		return
	}

	c.Data["json"] = Newresponse(http.StatusOK, "", nil)
	c.ServeJSON()
}
