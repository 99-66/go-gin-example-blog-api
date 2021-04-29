package blog_m

import "time"

type PostComment struct {
	Id int `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	PostId int `json:"postId" gorm:"column:postId;index:idx_comment_post,sort:asc;not null"`
	Post Post `gorm:"foreignKey:PostId;references:Id;constraint:fk_comment_post,OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	ParentId int `json:"parentId" gorm:"columns:parentId;index:idx_comment_parent,sort:asc;default:null"`
	Parent []PostComment `gorm:"foreignKey:ParentId;references:Id;constraint:fk_comment_parent;OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Title string `json:"title" gorm:"column:title;type:varchar(100);not null"`
	Published int `json:"published" gorm:"column:published;type:tinyint(1);not null;default:0"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;type:datetime;not null"`
	PublishedAt time.Time `json:"publishedAt" gorm:"column:publishedAt;type:datetime;default:null"`
	Content string `json:"content" gorm:"column:content;type:text;default:null"`
}

func (PostComment) TableName() string {
	return "post_comment"
}
