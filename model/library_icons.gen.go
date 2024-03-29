// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLibraryIcon = "library_icons"

// LibraryIcon mapped from table <library_icons>
type LibraryIcon struct {
	ID             int32     `gorm:"column:id;primaryKey;default:nextval('"LibraryIcons_Id_seq"'::regclass)" json:"id"`
	DateCreate     time.Time `gorm:"column:date_create" json:"date_create"`
	DateLastChange time.Time `gorm:"column:date_last_change" json:"date_last_change"`
	CreaterID      int64     `gorm:"column:creater_id" json:"creater_id"`
	ChangerID      int64     `gorm:"column:changer_id" json:"changer_id"`
	Title          string    `gorm:"column:title" json:"title"`
	HashTag        string    `gorm:"column:hash_tag" json:"hash_tag"`
	ObjType        string    `gorm:"column:obj_type;not null" json:"obj_type"`
	FullPath       string    `gorm:"column:full_path" json:"full_path"`
	FileName       string    `gorm:"column:file_name" json:"file_name"`
	Extension      string    `gorm:"column:extension" json:"extension"`
	Size           int64     `gorm:"column:size" json:"size"`
	FileType       string    `gorm:"column:file_type" json:"file_type"`
	ThumbnailURL   string    `gorm:"column:thumbnail_url" json:"thumbnail_url"`
	URL            string    `gorm:"column:url" json:"url"`
	IsArchive      bool      `gorm:"column:is_archive" json:"is_archive"`
}

// TableName LibraryIcon's table name
func (*LibraryIcon) TableName() string {
	return TableNameLibraryIcon
}
