package models

const TableNameCategory = "categories"

// Category mapped from table <categories>
type Category struct {
	ID               int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID категории" json:"id"`                                          // ID категории
	Name             string `gorm:"column:name;not null;comment:Название категории" json:"name"`                                                     // Название категории
	Description      string `gorm:"column:description;comment:Описание категории" json:"description"`                                                // Описание категории
	ParentCategoryID int32  `gorm:"column:parent_category_id;comment:ID родительской категории (для вложенных категорий)" json:"parent_category_id"` // ID родительской категории (для вложенных категорий)
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
