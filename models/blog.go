package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

// var (
// 	UserList map[string]*User
// )

func init() {

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/go_developer?charset=utf8")
	orm.RegisterModel(new(BlogCategory),new(Blog),new(User))
	createBlogTable()
}
type BlogCategory struct{
	Id 		int  `orm:"pk"`
    Title 	string `orm:"size(100)"`
    User 	*User `orm:"rel(fk);null;on_delete(set_null)"`
    CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
    Descri 	string `orm:size(100)`
}

type Blog struct{
	Id	string    `orm:"pk"`
	BlogCategory *BlogCategory `orm:"rel(fk);null;on_delete(set_null)"`
	User *User `orm:"rel(fk);null;on_delete(set_null)"`
	BlogTitle    string    `orm:"size(20)"`
	Content    string    `orm:"size(8000)"`
	ImageUrl      string    `orm:"size(100)"`
	Tags string    `orm:"size(100)"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	Public      int       `orm:"int"`
}


/**
* open create table auto
**/
func createBlogTable() {
	name := "default"
	force := false
	verbose := true                            //true :show the sql (create table) ;false not show
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
}

func AddBlog(b Blog) string {
	o := orm.NewOrm()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	orm.DefaultTimeLoc = time.Local
	o.Using("default")
	b.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	b.CreatedTime = time.Now()
	fmt.Println(b)
	fmt.Println(o.Insert(&b))
	return b.Id
}

