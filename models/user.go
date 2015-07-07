package models

import (
	// "errors"
	"fmt"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var (
	UserList map[string]*User
)
var rs orm.RawSeter
var o orm.Ormer

func init() {

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/go_developer?charset=utf8")
	UserList = make(map[string]*User)
}

type User struct {
	Id          string
	Username    string
	Password    string
	Gender      string
	Age         int
	Address     string
	Email       string
	CreatedTime string
	Weight      int
}

func AddUser(u User) string {
	o := orm.NewOrm()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	orm.DefaultTimeLoc = time.Local
	o.Using("default")
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	u.CreatedTime = time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(u)
	res, err := o.Raw("insert into user(id,username,password,gender,age,address,email,created_time,weight) values(?,?,?,?,?,?,?,?,?)",
		u.Id, u.Username, u.Password, u.Gender, u.Age, u.Address, u.Email, u.CreatedTime, u.Weight).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	} else {
		fmt.Println("insert error!")
	}
	return u.Id
}

func GetUser(uid string) (User, error) {
	var user User
	err := orm.NewOrm().QueryTable("user").Filter("Id", uid).One(&user)
	if err == nil {
		fmt.Println(user)
	}
	fmt.Println()
	return user, err
}

func GetAllUsers() []User {

	var users []User
	o = orm.NewOrm()
	rs = o.Raw("SELECT * FROM user")

	num, err := rs.QueryRows(&users)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Queried ", num, "users")
		for _, user := range users {
			fmt.Println(user)
		}
	}

	return users
}

func UpdateUser(uu *User) (num int64, err error) {
	o := orm.NewOrm()
	if uu != nil {
		num, err := o.Update(uu)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("update id is: %d", num)
		}
	}
	return num, err
}

func Login(username, password string) (bool, User) {
	var user User
	o = orm.NewOrm()
	rs = o.Raw("SELECT * FROM user where username=? and password=?", username, password)
	err := rs.QueryRow(&user)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
		// DefaultTimeLoc := time.Local
		// var createTime string
		// createTime = user.CreatedTime
		// user.CreatedTime, err := time.ParseInLocation("2006-01-02 15:04:05",user.CreatedTime, DefaultTimeLoc)
		// fmt.Println(user)
		return true, user
	}
	return false, user
}

func DeleteUser(uid string) {
	var user User
	o := orm.NewOrm()
	err := orm.NewOrm().QueryTable("user").Filter("Id", uid).One(&user)
	if err == nil {
		fmt.Println(user)
	}

	num, err := o.Delete(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
