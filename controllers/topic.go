package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
	c.TplName = "topic.html"
	topics, err := models.GetAllTopics("", false)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
}

func (c *TopicController) Post() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	tid := c.Input().Get("tid")
	category := c.Input().Get("category")
	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, content)

	} else {
		err = models.ModifyTopic(tid, category, title, content)
	}
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)
}

func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
}

func (c *TopicController) View() {
	c.TplName = "topic_view.html"
	topic, err := models.GetTopic(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Tid"] = c.Ctx.Input.Param("0")
	replies, err := models.GetAllReplies(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Replies"] = replies
}

func (c *TopicController) Modify() {
	c.TplName = "topic_modify.html"
	tid := c.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
}

func (c *TopicController) Delete() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	tid := c.Input().Get("tid")
	err := models.DeleteTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/topic", 302)
		return
	}
	c.Redirect("/topic", 302)

}
