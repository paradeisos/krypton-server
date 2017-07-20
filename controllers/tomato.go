package controllers

import (
	"encoding/json"
	"krypton-server/errors"
	"krypton-server/models"
	"krypton-server/options"

	"net/http"

	"github.com/astaxie/beego"
)

type Tomato struct {
	beego.Controller
}

func (c *Tomato) Post() {
	var opts *options.CreateTomatoOpts

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &opts)
	if err != nil {
		beego.Error(err)
	}

	tomato := models.Tomato.NewTomatoModel(opts.Uid, opts.Start, opts.End, opts.Desc, opts.Finished)

	if err := tomato.Save(); err != nil {
		beego.Error(err)
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
	var opts *options.UpdateTomatoOpts

	if json.Unmarshal(c.Ctx.Input.RequestBody, &opts) != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	tomato, err := models.Tomato.Find(opts.ID)
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	tomato.Start = opts.Start
	tomato.End = opts.End
	tomato.Desc = opts.Desc
	tomato.Finished = opts.Finished

	if err := tomato.Update(); err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InternalError)
		c.ServeJSON()
		return
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
