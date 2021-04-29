package blog_m

import "time"

type Post struct {
	Id int `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	AuthorId int `json:"authorId" gorm:"column:authorId;index:idx_post_user,sort:asc;not null"`
	User User `gorm:"foreignKey:AuthorId;references:Id;constraint:fk_post_user,OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	ParentId *int `json:"parentId" gorm:"column:parentId;index:idx_post_parent,sort:asc;null"`
	Parent []Post `gorm:"foreignKey:ParentId;references:Id;constraint:fk_post_parent,OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Title string `json:"title" gorm:"column:title;type:varchar(75);not null"`
	MetaTitle string `json:"metaTitle" gorm:"column:metaTitle;type:varchar(100)"`
	Slug string `json:"slug" gorm:"column:slug;type:varchar(100);uniqueIndex:ug_slug;not null"`
	Summary string `json:"summary" gorm:"column:summary;type:tinytext"`
	Published int `json:"published" gorm:"column:published;type:tinyint(1);not null;default:0"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;type:datetime;not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt;type:datetime"`
	PublishedAt time.Time `json:"publishedAt" gorm:"column:publishedAt;type:datetime"`
	Content string `json:"content" gorm:"column:content;type:text"`
}

func (Post) TableName() string {
	return "post"
}