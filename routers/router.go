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
	// 获取uri参数
	beego.Router("/uri/:id", &controllers.UriDemo{})
	// 获取parmes参数
	beego.Router("/uri", &controllers.QueryDemo{})
	// 获取post中的表单参数
	beego.Router("/user", &controllers.PostDemo{})
	// 获取表单数据解析到结构体中
	beego.Router("/parse", &controllers.ParseForm{})
	// 获取json数据
	beego.Router("/json", &controllers.ParseJson{})

}
