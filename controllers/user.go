package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"krypton-server/errors"
	"krypton-server/models"

	"github.com/astaxie/beego"
)

type User struct {
	beego.Controller
}

//TODO: session
func (c *User) Login() {
	var params *UserLoginParams
	if json.Unmarshal(c.Ctx.Input.RequestBody, &params) != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	var (
		user *models.UserModel
		err  error
	)

	if strings.Contains(params.Name, "@") {
		user, err = models.User.FindByEmail(params.Name)
	} else {
		user, err = models.User.FindByUsername(params.Name)
	}
	if err != nil || user.Password != params.Password {
		c.Data["json"] = errors.NewErrorResponse(errors.InvalidParameter)
		c.ServeJSON()
		return
	}

	token, err := sessionManager.NewSession(user.Id.Hex(), user.Username).Token()
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InternalError)
		c.ServeJSON()
		return
	}

	cookie := http.Cookie{
		Name:   "Authorization",
		Value:  token,
		Path:   "/",
		MaxAge: 86400,
	}

	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	c.Data["json"] = Newresponse(http.StatusOK, "", user)
	c.ServeJSON()
}

func (c *User) Logout() {
	cookie := http.Cookie{Name: "Authorization", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	c.Data["json"] = Newresponse(http.StatusOK, "", nil)
	c.ServeJSON()
}

// register
func (c *User) Register() {
	var params *UserRegisterParams
	resp := &Response{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InternalError)
		c.ServeJSON()
		return

	}

	user := models.User.NewUserModel(params.Username, params.Email, params.Password, "")
	err = user.Save()
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InternalError)
		c.ServeJSON()
		return
	}

	session := sessionManager.NewSession(user.Id.Hex(), params.Username)

	token, err := session.Token()
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InternalError)
		c.ServeJSON()
		return
	}

	mailer.SendRegisterMail(params.Email, fmt.Sprintf("%s%s?activeCode=%s", beego.AppConfig.String("host")+"/user/active", token))

	resp.Status = http.StatusOK
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *User) Active() {
	token := c.GetString("activeCode")
	resp := &Response{}
	session, err := sessionManager.NewSessionByToken(token)
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.AccessForbidden)
		c.ServeJSON()
		return
	}

	username := session.UserName
	user, err := models.User.FindByUsername(username)
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InternalError)
		c.ServeJSON()
		return
	}

	user.Status = models.UserStatusActive
	err = user.Save()
	if err != nil {
		c.Data["json"] = errors.NewErrorResponse(errors.InternalError)
		c.ServeJSON()
		return
	}

	resp.Status = http.StatusOK
	c.Data["json"] = resp
	c.ServeJSON()
}
