package Models

//分类
type Category struct {
	Model
	//分类名
	Name string `json:"name" gorm:"type:varchar(50)"`
	//图片
	Image string `json:"image" gorm:"type:varchar(255)"`
	//排序
	SortNo int64 `json:"sort_no"`
	//启用(1是2否)
	Enable int8 `json:"enable" gorm:"default:1"`

}
func (m *Category) TableName() string {
	return "category"
}
