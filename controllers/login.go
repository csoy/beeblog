package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"	
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
/*	isExit := c.Input().Get("exit")=="true"
	if isExit{
		c.Ctx.SetCookie("name", "", -1, "/")
                c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 301)
		return
	}*/
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	uname := c.Input().Get("uname")
	upwd := c.Input().Get("upwd")
	autologin := c.Input().Get("autologin") == "on"
	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("upwd")==upwd{
		maxAge := 0
		if autologin{
			maxAge = 1<<31 -1
		}
		c.Ctx.SetCookie("name", uname, maxAge, "/")
        	c.Ctx.SetCookie("pwd", upwd, maxAge, "/")
	}	
	c.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool{
	ck, err := ctx.Request.Cookie("name")
	if err != nil {
    		return false
	}
	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
    		return false
	}
	upwd := ck.Value
	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("upwd")==upwd
}
