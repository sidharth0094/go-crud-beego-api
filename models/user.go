package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User represents a person in the system
type User struct {
	Id        int
	FirstName string    `orm:"null"`
	LastName  string    `orm:"null"`
	Email     string    `orm:"null;unique"`
	Password  string    `orm:"null"`
	RegDate   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// GetAll function gets all users
func GetAllUsers() []*User {
	o := orm.NewOrm()
	var users []*User
	o.QueryTable(new(User)).All(&users)

	return users
}
