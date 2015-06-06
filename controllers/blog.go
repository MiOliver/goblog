package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"goblog/models"
)

// Operations about Users
type BlogController struct {
	beego.Controller
}

// @Title createBlog
// @Description create Blogs
// @Param	body		body 	models.Blog	true		"body for blog content"
// @Success 200 {int} models.Blog.Id
// @Failure 403 body is empty
// @router / [post]
func (b *BlogController) Post() {
	var blog models.Blog
	fmt.Println(string(b.Ctx.Input.RequestBody))
	json.Unmarshal(b.Ctx.Input.RequestBody, &blog)
	blogid := models.AddBlog(blog)
	b.Data["json"] = map[string]string{"blogid": blogid}
	b.ServeJson()
}
