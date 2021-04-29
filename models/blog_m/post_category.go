package blog_m

type PostCategory struct {
	PostId int `json:"postId" gorm:"column:postId;primaryKey;index:idx_pc_post,sort:asc;not null"`
	Post Post `gorm:"foreignKey:PostId;references:Id;constraint:fk_pc_post,OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CategoryId int `json:"categoryId" gorm:"column:categoryId;primaryKey;index:idx_pc_category,sort:asc;not null"`
	Category Category `gorm:"foreignKey:CategoryId;references:Id;constraint:fk_pc_category,OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}

func (PostCategory) TableName() string {
	return "post_category"
}
