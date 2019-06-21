package models

type User struct {
	Id        int `gorm:"primary_key"`
	Name        string
	Mobile      string
	CreatedAt     int
	UpdatedAt     int
}

func (User) Find(id int) User {
	var result User
	DB.Where("id = ?", id).Find(&result)
	return result
}