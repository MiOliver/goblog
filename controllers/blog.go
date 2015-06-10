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
// @Param  userId	query 	string	true		"The userId for get Blogs"
// @Success 200 {object} models.Blog
// @Failure 403 :userId is empty
// @router /getUserBlog [get]
func (u *BlogController) GetAllBlogs() {
	userId := u.GetString("userId")
	fmt.Println(userId)
	if userId != "" {
		blogs := models.GetAllBlogs(userId)
		u.Data["json"] = blogs
	}

	u.ServeJson()
}

// @Title GetBlog
// @Description get blog by blogId
// @Param	blogId	 query 	string	true		"The blogId for get Blogs"
// @Success 200 {object} models.Blog
// @Failure 403 :blogId is empty
// @router /getOneBlog [get]
func (u *BlogController) GetBlog() {
	// blogId := u.GetString(":blogId")  //param type difference path
	blogId := u.GetString("blogId") //param type difference query
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

// @Title SetBlog
// @Description update blog
// @Param	body	body 	models.Blog	true		"The blog you want to update"
// @Success 200 {int} models.Blog.Id
// @Failure 403 :blogId is empty
// @router /resetBlog [post]
func (b *BlogController) SetBlog() {
	var blog models.Blog
	fmt.Println(string(b.Ctx.Input.RequestBody))
	json.Unmarshal(b.Ctx.Input.RequestBody, &blog)
	fmt.Println(blog)
	fmt.Println(blog.Blogid)
	if blog.Blogid != "" {
		blogId, err := models.ChangeBlog(blog)
		if err != nil {
			b.Data["json"] = err
		} else {
			b.Data["json"] = blogId
		}
	}
	b.ServeJson()
}

// @Title DeleteBlog
// @Description delete blog by blogId
// @Param	blogId	 query 	string	true		"The blogId for delete Blogs"
// @Success 200 {string} delete success!
// @Failure 403 :blogId is empty
// @router /deleteBlog [get]
func (u *BlogController) DeleteBlog() {
	blogId := u.GetString("blogId") //param type difference query
	fmt.Println(blogId)
	if blogId != "" {
		num := models.DeleteBlogbyId(blogId)
		if num > 0 {
			u.Data["json"] = "delete success!"
		} else {
			u.Data["json"] = "delete failure!"
		}
	}
	u.ServeJson()
}
