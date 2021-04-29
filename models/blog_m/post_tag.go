package blog_m

type PostTag struct {
	PostId int `json:"postId" gorm:"column:postId;primaryKey;index:idx_pt_post,sort:asc;not null"`
	Post Post `gorm:"foreignKey:PostId;references:Id;constraint:fk_pt_post,OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	TagId int `json:"tagId" gorm:"column:tagId;primaryKey;index:idx_pt_tag,sort:asc;not null"`
	Tag Tag `gorm:"foreignKey:TagId;references:Id;constraint:fk_pt_tag,OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}

func (PostTag) TableName() string {
	return "post_tag"
}
