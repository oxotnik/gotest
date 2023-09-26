// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMapLabel = "map_labels"

// MapLabel mapped from table <map_labels>
type MapLabel struct {
	ID             int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	DateCreate     time.Time `gorm:"column:date_create" json:"date_create"`
	DateLastChange time.Time `gorm:"column:date_last_change" json:"date_last_change"`
	CreaterID      int64     `gorm:"column:creater_id" json:"creater_id"`
	ChangerID      int64     `gorm:"column:changer_id" json:"changer_id"`
	IsArchive      bool      `gorm:"column:is_archive" json:"is_archive"`
	Title          string    `gorm:"column:title" json:"title"`
	ProjectID      int32     `gorm:"column:project_id" json:"project_id"`
	LayerID        int32     `gorm:"column:layer_id" json:"layer_id"`
	StyleJSON      string    `gorm:"column:style_json" json:"style_json"`
}

// TableName MapLabel's table name
func (*MapLabel) TableName() string {
	return TableNameMapLabel
}