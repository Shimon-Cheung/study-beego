package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type RenderController struct {
	beego.Controller
}

func (c *RenderController) Get() {
	// 结构体渲染
	type People struct {
		Name string
		Age  int
	}
	c.Data["People"] = People{Name: "xmzhang", Age: 24}
	// 数组渲染
	arr := [5]int{1, 2, 3, 4, 5}
	c.Data["arr"] = arr
	// map渲染
	m := map[string]interface{}{
		"name": "xmzhang",
	}
	c.Data["m"] = m
	// 指定渲染模版文件
	c.TplName = "render.tpl"

}
