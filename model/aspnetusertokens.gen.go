// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAspNetUserToken = "AspNetUserTokens"

// AspNetUserToken mapped from table <AspNetUserTokens>
type AspNetUserToken struct {
	UserID        string `gorm:"column:UserId;primaryKey" json:"UserId"`
	LoginProvider string `gorm:"column:LoginProvider;primaryKey" json:"LoginProvider"`
	Name          string `gorm:"column:Name;primaryKey" json:"Name"`
	Value         string `gorm:"column:Value" json:"Value"`
}

// TableName AspNetUserToken's table name
func (*AspNetUserToken) TableName() string {
	return TableNameAspNetUserToken
}
