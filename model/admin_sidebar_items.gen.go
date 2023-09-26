// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAdminSidebarItem = "admin_sidebar_items"

// AdminSidebarItem mapped from table <admin_sidebar_items>
type AdminSidebarItem struct {
	ID           int32  `gorm:"column:id;primaryKey;default:nextval('"AdminSidebarItems_id_seq"'::regclass)" json:"id"`
	ParentID     int32  `gorm:"column:parent_id" json:"parent_id"`
	Icon         string `gorm:"column:icon" json:"icon"`
	Title        string `gorm:"column:title" json:"title"`
	URL          string `gorm:"column:url" json:"url"`
	IsDisable    bool   `gorm:"column:is_disable" json:"is_disable"`
	Priority     int32  `gorm:"column:priority" json:"priority"`
	IsCategory   bool   `gorm:"column:is_category" json:"is_category"`
	ModuleID     int32  `gorm:"column:module_id" json:"module_id"`
	Role         string `gorm:"column:role" json:"role"`
	IconCategory string `gorm:"column:icon_category" json:"icon_category"`
}

// TableName AdminSidebarItem's table name
func (*AdminSidebarItem) TableName() string {
	return TableNameAdminSidebarItem
}
