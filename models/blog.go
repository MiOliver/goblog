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
	Blogid         string
	Blogcategoryid int64
	UserId         string
	BlogTitle      string
	Content        string
	ImageUrl       string
	Tags           string
	CreatedTime    time.Time
	Public         int
}

func init() {

}

func AddBlog(b Blog) int64 {
	var affectedNum int64
	affectedNum = 0
	o := orm.NewOrm()
	orm.DefaultTimeLoc = time.Local
	b.CreatedTime = time.Now()
	o.Using("default")
	fmt.Println(b)
	fmt.Println(b.Blogcategoryid)
	fmt.Println(b.BlogTitle)
	res, err := o.Raw("insert into blog(blog_category_id,user_id,blog_title,content,imageurl,tags,created_time,public) values(?,?,?,?,?,?,?,?)",
		b.Blogcategoryid, b.UserId, b.BlogTitle, b.Content, b.ImageUrl, b.Tags, b.CreatedTime, b.Public).Exec()
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
