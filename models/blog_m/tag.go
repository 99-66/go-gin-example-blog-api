package blog_m

type Tag struct {
	Id int `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Title string `json:"title" gorm:"column:title;type:varchar(75);not null"`
	MetaTitle string `json:"metaTitle" gorm:"column:metaTitle;type:varchar(100);default:null"`
	Slug string `json:"slug" gorm:"column:slug;type:varchar(100);not null"`
	Content string `json:"content" gorm:"column:content;type:text;default:null"`
}

func (Tag) TableName() string {
	return "tag"
}

