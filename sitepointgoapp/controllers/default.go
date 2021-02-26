package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) HelloSitepoint() {
	c.Data["Website"] = "My Website"
	c.Data["Email"] = "942091885@qq.com"
	c.Data["EmailName"] = "cyanxu"
	c.TplName = "hello-sitepoint.tpl"
}

func MakeExcelByStructs(sheetName string)int{
	if sheetName != ""{
		return 1
	}
	return 0
}