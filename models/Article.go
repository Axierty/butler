package models

type Article struct {
	Id        int `gorm:"primary_key"`
	Title     string
	Author    string
	Content   string
	CreatedAt int
	UpdatedAt int
	IsDel     int
	Sort      int
	IsTop     int
	AdminId   int
}


//获取文章详情
func (Article) GetArticleDetail(id int) Article {

	var result Article
	DB.Where("id = ?", id).Find(&result)
	return result
}


func (Article) GetArticleList(id int) []Article {

	var result []Article
	DB.Where("id = ?", id).Find(&result)

	return result
}