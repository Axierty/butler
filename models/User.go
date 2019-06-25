package models

import (
	"fmt"
	"reflect"
	"errors"
)

type User struct {
	Id        int `gorm:"primary_key"`
	Name        string
	Mobile      string
	CreatedAt     int
	UpdatedAt     int
}

func (User) FindAll() []User {
	var users []User
	DB.Find(&users)
	return users
}


func (User) Find(id int) User {
	var result User
	DB.Where("id = ?", id).Find(&result)
	return result
}

func (User) UpdateById(id int,data map[string]interface{}) error {
	var user User
	fmt.Println("默认初始化=>",user)
	DB.Find(&user, id)

	//通过反射判断是否为空
	if reflect.ValueOf(user).FieldByName("Id").Int() == 0 {
		return errors.New("修改数据为空")
	}

	DB.Model(&user).Updates(data)
	return nil
	//反射拿到user
	//userReflect := reflect.ValueOf(user)
	//
	//for k, v := range data {
	//
	//	isExists := userReflect.FieldByName(k).Int()
	//	if isExists == 0 {
	//		return nil
	//	}
	//	user.k = v
	//}
	//db.Model(&user).Updates()

}