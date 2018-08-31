package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type MainController struct {
	beego.Controller
}

// @Title demoGet
// @Description 测试get方法
// @Param   name     query    string  true        "测试name"
// @Param   addr     query    string  true        "测试addr"
// @Success 200 sucees
// @Failure 400 为上传参数或错误参数
// @Failure 404 not found
// @router /info [get]
func (c *MainController) Get() {
	name := c.GetString("name")
	addr := c.GetString("addr")
	if strings.EqualFold(name, "") {
		c.Abort("400")
	}
	if strings.EqualFold(addr, "") {
		c.Abort("400")
	}
	c.Data["json"] = "test"
	c.ServeJSON()
}

// @Title getAge
// @Description 测试自定义get方法
// @Success 200 sucees
// @Failure 404 not found
// @router /age [get]
func (c *MainController) GetAge() {
	c.Data["json"] = "13"
	c.ServeJSON()
}
