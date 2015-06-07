package models

import (
	"fmt"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

// BlogCategory *BlogCategory `orm:"rel(fk);null;on_delete(set_null)"`

type Blog struct {
	Blogid	string    
	BlogCategoryId int64
	UserId  string  
	BlogTitle    string    
	Content    string   
	ImageUrl      string    
	Tags string    
	CreatedTime time.Time 
	Public      int      
}

func init() {

}

func AddBlog(b Blog) string {
	o := orm.NewOrm()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	orm.DefaultTimeLoc = time.Local
	o.Using("default")
	b.Blogid = "blog_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	b.CreatedTime = time.Now()
	fmt.Println(b)
	fmt.Println(o.Insert(&b))
	return b.Blogid
}


