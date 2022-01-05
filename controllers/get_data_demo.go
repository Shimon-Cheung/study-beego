package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type UriDemo struct {
	beego.Controller
}

func (d *UriDemo) Get() {
	getString := d.GetString(":id")
	d.Ctx.WriteString(getString)
}

type QueryDemo struct {
	beego.Controller
}

func (q *QueryDemo) Get() {
	getString := q.GetString("id")
	q.Ctx.WriteString(getString)
}

type PostDemo struct {
	beego.Controller
}

func (c *PostDemo) Post() {
	name := c.GetString("name")
	age := c.GetString("age")
	c.Ctx.WriteString(name + ":" + age)
}

type ParseForm struct {
	beego.Controller
}

func (c *ParseForm) Post() {
	type my struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}
	m := new(my)
	c.ParseForm(m)
	fmt.Println(m)

	c.Ctx.WriteString("解析成功")
}

type ParseJson struct {
	beego.Controller
}

func (c *ParseJson) Post() {
	type HuYang struct {
		Name string `json:"name"`
	}
	ob := new(HuYang)
	json.Unmarshal(c.Ctx.Input.RequestBody, ob)
	fmt.Println(ob)
	c.Ctx.WriteString("解析成功")
}
