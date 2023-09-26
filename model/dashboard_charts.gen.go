// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDashboardChart = "dashboard_charts"

// DashboardChart mapped from table <dashboard_charts>
type DashboardChart struct {
	ID             int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	DateCreate     time.Time `gorm:"column:date_create" json:"date_create"`
	DateLastChange time.Time `gorm:"column:date_last_change" json:"date_last_change"`
	CreaterID      int64     `gorm:"column:creater_id" json:"creater_id"`
	ChangerID      int64     `gorm:"column:changer_id" json:"changer_id"`
	ProjectID      int32     `gorm:"column:project_id" json:"project_id"`
	Title          string    `gorm:"column:title" json:"title"`
	SqlFunction    string    `gorm:"column:sql_function" json:"sql_function"`
	SettingsJSON   string    `gorm:"column:settings_json" json:"settings_json"`
	Priority       int32     `gorm:"column:priority" json:"priority"`
	Type           string    `gorm:"column:type" json:"type"`
}

// TableName DashboardChart's table name
func (*DashboardChart) TableName() string {
	return TableNameDashboardChart
}
