package models

import (
	"fmt"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	// "strconv"
)

// BlogCategory *BlogCategory `orm:"rel(fk);null;on_delete(set_null)"`

type Blog struct {
	Blogid         int64
	BlogCategoryId int64
	UserId         string
	BlogTitle      string
	Content        string
	ImageUrl       string
	Tags           string
	CreatedTime    string
	Public         int
}

func init() {

}

func AddBlog(b Blog) int64 {
	var affectedNum int64
	affectedNum = 0
	o := orm.NewOrm()
	orm.DefaultTimeLoc = time.Local
	fmt.Println(b)
	b.CreatedTime = time.Now().Format("2006-01-02 15:04:05")
	o.Using("default")
	fmt.Println(b.BlogCategoryId)
	fmt.Println(b.BlogTitle)
	res, err := o.Raw("insert into blog(blog_category_id,user_id,blog_title,content,imageurl,tags,created_time,public) values(?,?,?,?,?,?,?,?)",
		b.BlogCategoryId, b.UserId, b.BlogTitle, b.Content, b.ImageUrl, b.Tags, b.CreatedTime, b.Public).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		affectedNum = num
	} else {
		fmt.Println("insert error!")
	}
	return affectedNum
}

func GetAllBlogs(userid string) []Blog {
	var blogs []Blog
	o = orm.NewOrm()
	fmt.Println(userid)
	rs = o.Raw("select * from blog where user_id=?", userid)
	num, err := rs.QueryRows(&blogs)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Queried ", num, "blogs")
		for _, blog := range blogs {
			fmt.Println(blog)
		}
	}
	return blogs
}

func GetRecentBlogs(userid string) []Blog {
	var blogs []Blog
	o = orm.NewOrm()
	fmt.Println(userid)
	rs = o.Raw("select * from blog where user_id=? order by blogid desc limit 2", userid)
	num, err := rs.QueryRows(&blogs)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Queried ", num, "blogs")
		for _, blog := range blogs {
			fmt.Println(blog)
		}
	}
	return blogs
}

func GetBlogbyId(blogId string) (Blog, error) {
	var blog Blog
	o = orm.NewOrm()
	fmt.Println(blogId)
	rs = o.Raw("select * from blog where blogid=?", blogId)
	err := rs.QueryRow(&blog)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(blog)
	}
	return blog, err
}

func ChangeBlog(b Blog) (int64, error) {
	o = orm.NewOrm()
	o.Using("default")
	fmt.Println(b.Blogid)
	res, err := o.Raw("update blog set blog_category_id=?,blog_title=?,content=?,imageurl=?,tags=?,created_time=?,public=? where blogid=?",
		b.BlogCategoryId, b.BlogTitle, b.Content, b.ImageUrl, b.Tags, b.CreatedTime, b.Public, b.Blogid).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	} else {
		fmt.Println("update error!")
	}

	return b.Blogid, err
}

func DeleteBlogbyId(blogId string) int64 {
	var affectNum int64
	affectNum = 0
	o = orm.NewOrm()
	fmt.Println(blogId)
	res, err := o.Raw("delete from blog where blogid=?", blogId).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		affectNum = num
		fmt.Println("mysql row affected nums: ", num)
	}
	return affectNum
}
