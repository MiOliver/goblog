package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var (
	UserList map[string]*User
)

func init() {

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/go_developer?charset=utf8")
	orm.RegisterModel(new(User))
	UserList = make(map[string]*User)
	// u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	// UserList["user_11111"] = &u
	createTable()
}

type User struct {
	Id          string    `orm:"pk"`
	Username    string    `orm:"size(20)"`
	Password    string    `orm:"size(100)"`
	Gender      string    `orm:"size(2)"`
	Age         int       `orm:"int"`
	Address     string    `orm:"size(100)"`
	Email       string    `orm:"size(30)"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	Weight      int       `orm:"int"`
}

/**
* open create table auto
**/
func createTable() {
	name := "default"
	force := false
	verbose := true                            //true :show the sql (create table) ;false not show
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
}

func AddUser(u User) string {
	o := orm.NewOrm()
	orm.DefaultTimeLoc = time.Local
	o.Using("default")
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	u.CreatedTime=time.Now()
	fmt.Println(u)
	fmt.Println(o.Insert(&u))
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (num int64, err error) {
	o := orm.NewOrm()
	if uu != nil {
		num, err := o.Update(&uu)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("update id is: %d", num)
		}
	}
	return num, err
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
