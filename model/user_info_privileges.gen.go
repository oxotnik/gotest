// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserInfoPrivilege = "user_info_privileges"

// UserInfoPrivilege mapped from table <user_info_privileges>
type UserInfoPrivilege struct {
	ID             int32     `gorm:"column:id;primaryKey;default:nextval('"UserInfoPrivileges_Id_seq"'::regclass)" json:"id"`
	UserID         int64     `gorm:"column:user_id;not null" json:"user_id"`
	DateCreate     time.Time `gorm:"column:date_create" json:"date_create"`
	DateLastChange time.Time `gorm:"column:date_last_change" json:"date_last_change"`
	CreaterID      int64     `gorm:"column:creater_id" json:"creater_id"`
	ChangerID      int64     `gorm:"column:changer_id" json:"changer_id"`
	PrivilegesID   int32     `gorm:"column:privileges_id;not null" json:"privileges_id"`
}

// TableName UserInfoPrivilege's table name
func (*UserInfoPrivilege) TableName() string {
	return TableNameUserInfoPrivilege
}
