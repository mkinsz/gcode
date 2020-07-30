package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//设置全局表名禁用复数
	db.SingularTable(true)
}

type User struct {
	Id    int
	Name  string
	Age   int
	Sex   byte
	Phone string
}

//插入数据
func (user *User) Insert() {
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	db.Table("user").Create(user)
}

func Update() {
	//使用struct的方式
	// user := User{Id: 1, Name: "xiaoming"}
	// db.Model(&user).Update(user)

	//Map的方式
	// db.Model(&User{}).Where("sex = ?", 1).Update("name", "xiaohong")

	// 如果你想手动将某个字段set为空值, 可以使用单独选定某些字段的方式来更新：
	// user := User{Id: 1}
	// db.Model(&user).Select("name").Update(map[string]interface{}{"name": "", "age": 0})

	//
	user := User{Id: 1, Name: "xiaoming", Age: 12}
	db.Model(&user).Omit("name").Update(&user)
}
