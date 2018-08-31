// @APIVersion 1.0.0
// @Title 测试使用swagger
// @Description 测试使用swagger的演示样例
// @Contact tingshage@163.com
package routers

import (
	"demo_swagger/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/demo",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
