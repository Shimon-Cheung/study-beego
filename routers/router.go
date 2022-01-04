package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"study-beego/controllers"
)

func init() {
	// 默认的beego根路由页面
	beego.Router("/", &controllers.MainController{})
	// 渲染样例学习路由
	beego.Router("/render", &controllers.RenderController{})

}
