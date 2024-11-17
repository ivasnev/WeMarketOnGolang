package models

const TableNameCategory = "categories"

// Category mapped from table <categories>
type Category struct {
	ID               int32   `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:ID категории" json:"id"`                                          // ID категории
	Name             string  `gorm:"column:name;type:character varying(100);not null;comment:Название категории" json:"name"`                                      // Название категории
	Description      *string `gorm:"column:description;type:text;comment:Описание категории" json:"description"`                                                   // Описание категории
	ParentCategoryID *int32  `gorm:"column:parent_category_id;type:integer;comment:ID родительской категории (для вложенных категорий)" json:"parent_category_id"` // ID родительской категории (для вложенных категорий)
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
