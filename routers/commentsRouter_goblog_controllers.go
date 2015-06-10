package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["goblog/controllers:ObjectController"] = append(beego.GlobalControllerRouter["goblog/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:ObjectController"] = append(beego.GlobalControllerRouter["goblog/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:ObjectController"] = append(beego.GlobalControllerRouter["goblog/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:ObjectController"] = append(beego.GlobalControllerRouter["goblog/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:ObjectController"] = append(beego.GlobalControllerRouter["goblog/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:BlogCategoryController"] = append(beego.GlobalControllerRouter["goblog/controllers:BlogCategoryController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:UserController"] = append(beego.GlobalControllerRouter["goblog/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:UserController"] = append(beego.GlobalControllerRouter["goblog/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:UserController"] = append(beego.GlobalControllerRouter["goblog/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:UserController"] = append(beego.GlobalControllerRouter["goblog/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:UserController"] = append(beego.GlobalControllerRouter["goblog/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:UserController"] = append(beego.GlobalControllerRouter["goblog/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:UserController"] = append(beego.GlobalControllerRouter["goblog/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:BlogController"] = append(beego.GlobalControllerRouter["goblog/controllers:BlogController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:BlogController"] = append(beego.GlobalControllerRouter["goblog/controllers:BlogController"],
		beego.ControllerComments{
			"GetAllBlogs",
			`/getUserBlog`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:BlogController"] = append(beego.GlobalControllerRouter["goblog/controllers:BlogController"],
		beego.ControllerComments{
			"GetBlog",
			`/getOneBlog`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goblog/controllers:BlogController"] = append(beego.GlobalControllerRouter["goblog/controllers:BlogController"],
		beego.ControllerComments{
			"SetBlog",
			`/resetBlog`,
			[]string{"post"},
			nil})

}
