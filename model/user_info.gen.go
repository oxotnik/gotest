// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserInfo = "user_info"

// UserInfo mapped from table <user_info>
type UserInfo struct {
	ID             int64     `gorm:"column:id;primaryKey;" json:"id"`
	UserName       string    `gorm:"column:user_name" json:"user_name"`
	LastName       string    `gorm:"column:last_name" json:"last_name"`
	FirstName      string    `gorm:"column:first_name" json:"first_name"`
	MiddleName     string    `gorm:"column:middle_name" json:"middle_name"`
	Email          string    `gorm:"column:email" json:"email"`
	Phone          string    `gorm:"column:phone" json:"phone"`
	IP             string    `gorm:"column:IP" json:"IP"`
	DateLastOnline time.Time `gorm:"column:date_last_online" json:"date_last_online"`
	DateOfBirth    time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`
	Photo          string    `gorm:"column:photo" json:"photo"`
	PhotoFull      string    `gorm:"column:photo_full" json:"photo_full"`
	Sex            bool      `gorm:"column:sex;not null" json:"sex"`
	IsBlocked      bool      `gorm:"column:is_blocked" json:"is_blocked"`
	IsArchive      bool      `gorm:"column:is_archive" json:"is_archive"`
	IsOnline       bool      `gorm:"column:is_online" json:"is_online"`
	IDStatus       int32     `gorm:"column:id_status" json:"id_status"`
	DateStarted    time.Time `gorm:"column:date_started" json:"date_started"`
	IsNewPassword  bool      `gorm:"column:is_new_password;not null" json:"is_new_password"`
	IsUser         bool      `gorm:"column:is_user;not null" json:"is_user"`
	DateCreate     time.Time `gorm:"column:date_create" json:"date_create"`
	DateLastChange time.Time `gorm:"column:date_last_change" json:"date_last_change"`
	CreaterID      int64     `gorm:"column:creater_id" json:"creater_id"`
	ChangerID      int64     `gorm:"column:changer_id" json:"changer_id"`
	OrgStructureID int32     `gorm:"column:org_structure_id" json:"org_structure_id"`
	PositionID     int32     `gorm:"column:position_id" json:"position_id"`
	DateRegister   time.Time `gorm:"column:date_register" json:"date_register"`
}

// TableName UserInfo's table name
func (*UserInfo) TableName() string {
	return TableNameUserInfo
}
