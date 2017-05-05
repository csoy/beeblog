package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Add() {
	tid := c.Input().Get("tid")
	name := c.Input().Get("nickname")
	content := c.Input().Get("content")
	err := models.AddReply(tid, name, content)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	err = models.UpdTopic(tid, true)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Redirect("/topic/view/"+tid, 302)
}

func (c *ReplyController) Delete() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	tid := c.Input().Get("tid")
	rid := c.Input().Get("rid")
	err := models.DelReply(rid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	err = models.UpdTopic(tid, false)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Redirect("/topic/view/"+tid, 302)

}
