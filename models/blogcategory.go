package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)
// User 	*User `orm:"rel(fk);null;on_delete(set_null)"`
type BlogCategory struct {
	Id 		int64  
    Title 	string 
    UserId  string  
    CreatedTime time.Time 
    Descri 	string 
}


func init() {

}

func AddBlogCategory(b BlogCategory) string {
	o := orm.NewOrm()
	orm.DefaultTimeLoc = time.Local
	b.CreatedTime = time.Now();
	o.Using("default")
	fmt.Println(b)
	res, err := o.Raw("insert into blogcategory(user_id,title,descri,created_time) values(?,?,?,?)", 
	b.UserId,b.Title,b.Descri,b.CreatedTime).Exec()
	if err == nil {
    num, _ := res.RowsAffected()
    fmt.Println("mysql row affected nums: ", num)
	}else{
		fmt.Println("insert error!")
	}
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

