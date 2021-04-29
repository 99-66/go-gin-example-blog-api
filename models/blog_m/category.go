package blog_m

type Category struct {
	Id int `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ParentId int `json:"parentId" gorm:"column:parentId;index:idx_category_parent,sort:asc;default:null"`
	Parent []Category `gorm:"foreignKey:ParentId;references:Id;constraint:fk_category_parent,OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Title string `json:"title" gorm:"column:title;type:varchar(75);not null"`
	MetaTitle string `json:"metaTitle" gorm:"column:metaTitle;type:varchar(100);default:null"`
	Slug string `json:"slug" gorm:"column:slug;type:varchar(100);not null"`
	Content string `json:"content" gorm:"column:content;type:text;default:null"`
}

func (Category) TableName() string {
	return "category"
}
