package blog_m

type PostMeta struct {
	Id int `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	PostId int `json:"postId" gorm:"column:postId;index:idx_meta_post,sort:asc;uniqueIndex:uq_post_meta,sort:asc;not null"`
	Post Post `gorm:"foreignKey:PostId;references:Id;constraint:fk_meta_post,OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Key string `json:"key" gorm:"column:key;type:varchar(50);uniqueIndex:uq_post_meta,sort:asc;not null"`
	Content string `json:"content" gorm:"column:content;type:text;default:null"`
}

func (PostMeta) TableName() string {
	return "post_meta"
}
