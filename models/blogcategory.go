package models

import (
	"fmt"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)
// User 	*User `orm:"rel(fk);null;on_delete(set_null)"`
type BlogCategory struct {
	Id 		int64  
    Title 	string 
    UserId string  
    CreatedTime time.Time 
    Descri 	string 
}

// BlogCategory *BlogCategory `orm:"rel(fk);null;on_delete(set_null)"`



func init() {

}

func AddBlogCategory(b BlogCategory) string {
	o := orm.NewOrm()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	orm.DefaultTimeLoc = time.Local
	o.Using("default")
	b.Id=time.Now().UnixNano()
	b.CreatedTime = time.Now();
	fmt.Println(b)
	fmt.Println(o.Insert(&b))
	return strconv.FormatInt(int64(b.Id),10)
}

