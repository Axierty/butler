package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"errors"
	"strconv"
	"time"
)

type User struct {
	Id        int `gorm:"primary_key"`
	Name        string
	Mobile      string
	Password    string
	CreatedAt     int
	UpdatedAt     int
}


//生成token
func (User) GenerateToken() string {

	h := md5.New()
	curtime := time.Now().Unix()

	h.Write([]byte(strconv.FormatInt(curtime, 10)))
	token := hex.EncodeToString(h.Sum(nil))

	return token
}

//检查用户密码是否正确
func (User) CheckPassword(mobile string,password string) (error,User) {
	var user User
	DB.Where("mobile = ?", mobile).Find(&user)

	if user.Id == 0 {
		return errors.New("未知的用户手机号"),user
	}

	h := md5.New()
	h.Write([]byte(password))
	passwordNew := hex.EncodeToString(h.Sum(nil))
	if passwordNew != user.Password {
		return errors.New("账号密码错误"),user
	}

	return nil,user
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
}