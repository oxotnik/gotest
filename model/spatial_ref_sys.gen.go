// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSpatialRefSy = "spatial_ref_sys"

// SpatialRefSy mapped from table <spatial_ref_sys>
type SpatialRefSy struct {
	Srid      int32  `gorm:"column:srid;primaryKey" json:"srid"`
	AuthName  string `gorm:"column:auth_name" json:"auth_name"`
	AuthSrid  int32  `gorm:"column:auth_srid" json:"auth_srid"`
	Srtext    string `gorm:"column:srtext" json:"srtext"`
	Proj4text string `gorm:"column:proj4text" json:"proj4text"`
}

// TableName SpatialRefSy's table name
func (*SpatialRefSy) TableName() string {
	return TableNameSpatialRefSy
}
