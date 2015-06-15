package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"goblog/models"
)

// Operations about BlogCategory
type BlogCategoryController struct {
	beego.Controller
}

// @Title createBlogCategory
// @Description create BlogCategory
// @Param	body		body 	models.BlogCategory	true		"body for BlogCategory content"
// @Success 200 {int} models.BlogCategory.Id
// @Failure 403 body is empty
// @router / [post]
func (b *BlogCategoryController) Post() {
	var bc models.BlogCategory
	fmt.Println(string(b.Ctx.Input.RequestBody))
	json.Unmarshal(b.Ctx.Input.RequestBody, &bc)
	blogid := models.AddBlogCategory(bc)
	b.Data["json"] = map[string]string{"blogid": blogid}
	b.ServeJson()
}

// @Title GetAllBlogCategory
// @Description get all BlogCategory
// @Param  userId	query 	string	true		"The userId for get BlogCategory"
// @Success 200 {object} models.BlogCategory
// @Failure 403 :userId is empty
// @router /getUserBlogCategory [get]
func (u *BlogCategoryController) GetBlogCategory() {
	userId := u.GetString("userId")
	fmt.Println(userId)
	if userId != "" {
		blogCategorys := models.GetAllBlogCategory(userId)
		u.Data["json"] = blogCategorys
	}

	u.ServeJson()
}
