package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func init() {

}
func (c *HomeController) Get() {
	log := logs.NewLogger(10000)
	log.SetLogger("console", "")
	c.Data["IsHome"] = true
	c.TplName = "home.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics, err := models.GetAllTopics(c.Input().Get("cate"), true)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
		//log.Trace("test", topics[0].Id)
	}
	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Caregories"] = categories
	}
}
