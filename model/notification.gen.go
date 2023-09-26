// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameNotification = "notification"

// Notification mapped from table <notification>
type Notification struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	DateCreate time.Time `gorm:"column:date_create" json:"date_create"`
	ProjectID  int32     `gorm:"column:project_id" json:"project_id"`
	NotiType   string    `gorm:"column:noti_type" json:"noti_type"`
	Title      string    `gorm:"column:title" json:"title"`
	Note       string    `gorm:"column:note" json:"note"`
	ObjID      int64     `gorm:"column:obj_id" json:"obj_id"`
	ObjType    string    `gorm:"column:obj_type" json:"obj_type"`
	JSON       string    `gorm:"column:json" json:"json"`
	Idglm      int64     `gorm:"column:idglm" json:"idglm"`
	IsReadNoti bool      `gorm:"column:is_read_noti;not null;default:true" json:"is_read_noti"`
	ExtID      int64     `gorm:"column:ext_id" json:"ext_id"`
}

// TableName Notification's table name
func (*Notification) TableName() string {
	return TableNameNotification
}