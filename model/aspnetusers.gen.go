// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAspNetUser = "AspNetUsers"

// AspNetUser mapped from table <AspNetUsers>
type AspNetUser struct {
	ID                   string    `gorm:"column:Id;primaryKey" json:"Id"`
	UserName             string    `gorm:"column:UserName" json:"UserName"`
	NormalizedUserName   string    `gorm:"column:NormalizedUserName" json:"NormalizedUserName"`
	Email                string    `gorm:"column:Email" json:"Email"`
	NormalizedEmail      string    `gorm:"column:NormalizedEmail" json:"NormalizedEmail"`
	EmailConfirmed       bool      `gorm:"column:EmailConfirmed;not null" json:"EmailConfirmed"`
	PasswordHash         string    `gorm:"column:PasswordHash" json:"PasswordHash"`
	SecurityStamp        string    `gorm:"column:SecurityStamp" json:"SecurityStamp"`
	ConcurrencyStamp     string    `gorm:"column:ConcurrencyStamp" json:"ConcurrencyStamp"`
	PhoneNumber          string    `gorm:"column:PhoneNumber" json:"PhoneNumber"`
	PhoneNumberConfirmed bool      `gorm:"column:PhoneNumberConfirmed;not null" json:"PhoneNumberConfirmed"`
	TwoFactorEnabled     bool      `gorm:"column:TwoFactorEnabled;not null" json:"TwoFactorEnabled"`
	LockoutEnd           time.Time `gorm:"column:LockoutEnd" json:"LockoutEnd"`
	LockoutEnabled       bool      `gorm:"column:LockoutEnabled;not null" json:"LockoutEnabled"`
	AccessFailedCount    int32     `gorm:"column:AccessFailedCount;not null" json:"AccessFailedCount"`
}

// TableName AspNetUser's table name
func (*AspNetUser) TableName() string {
	return TableNameAspNetUser
}