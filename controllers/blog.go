package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"goblog/models"
)

// Operations about Blogs
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
	b.Data["json"] = map[string]int64{"blogid": blogid}
	b.ServeJson()
}

// @Title GetAllBlogs
// @Description get all Blogs
// @Param userId		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Blog
// @Failure 403 :userId is empty
// @router /:userId [get]
func (u *BlogController) GetAllBlogs() {
	userId := u.GetString(":userId")
	fmt.Println(userId)
	if userId != "" {
		blogs := models.GetAllBlogs(userId)
		u.Data["json"] = blogs
	}

	u.ServeJson()
}

// @Title GetBlog
// @Description get blog by blogId
// @Param	blogId		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Blog
// @Failure 403 :blogId is empty
// @router /:blogId [get]
func (u *BlogController) GetBlog() {
	blogId := u.GetString(":blogId")
	fmt.Println(blogId)
	if blogId != "" {
		blog, err := models.GetBlogbyId(blogId)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = blog
		}
	}
	u.ServeJson()
}
