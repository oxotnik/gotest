// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"glm-go-project/model"
)

func newMapObject(db *gorm.DB, opts ...gen.DOOption) mapObject {
	_mapObject := mapObject{}

	_mapObject.mapObjectDo.UseDB(db, opts...)
	_mapObject.mapObjectDo.UseModel(&model.MapObject{})

	tableName := _mapObject.mapObjectDo.TableName()
	_mapObject.ALL = field.NewAsterisk(tableName)
	_mapObject.ID = field.NewInt64(tableName, "id")
	_mapObject.ProjectID = field.NewInt32(tableName, "project_id")
	_mapObject.LayerID = field.NewInt32(tableName, "layer_id")
	_mapObject.ZuluTypeID = field.NewInt32(tableName, "zulu_type_id")
	_mapObject.ZuluModeID = field.NewInt32(tableName, "zulu_mode_id")
	_mapObject.Geometry = field.NewString(tableName, "geometry")
	_mapObject.Sys = field.NewInt64(tableName, "sys")
	_mapObject.Color = field.NewString(tableName, "color")
	_mapObject.Angle = field.NewFloat64(tableName, "angle")
	_mapObject.DateCreateObj = field.NewTime(tableName, "date_create_obj")
	_mapObject.DateDeleteObj = field.NewTime(tableName, "date_delete_obj")
	_mapObject.ParentID = field.NewInt64(tableName, "parent_id")
	_mapObject.OldObjID = field.NewInt64(tableName, "old_obj_id")
	_mapObject.ExtID = field.NewInt64(tableName, "ext_id")
	_mapObject.DateLastChange = field.NewTime(tableName, "date_last_change")

	_mapObject.fillFieldMap()

	return _mapObject
}

type mapObject struct {
	mapObjectDo

	ALL            field.Asterisk
	ID             field.Int64
	ProjectID      field.Int32
	LayerID        field.Int32
	ZuluTypeID     field.Int32
	ZuluModeID     field.Int32
	Geometry       field.String
	Sys            field.Int64
	Color          field.String
	Angle          field.Float64
	DateCreateObj  field.Time
	DateDeleteObj  field.Time
	ParentID       field.Int64
	OldObjID       field.Int64
	ExtID          field.Int64
	DateLastChange field.Time

	fieldMap map[string]field.Expr
}

func (m mapObject) Table(newTableName string) *mapObject {
	m.mapObjectDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m mapObject) As(alias string) *mapObject {
	m.mapObjectDo.DO = *(m.mapObjectDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *mapObject) updateTableName(table string) *mapObject {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.ProjectID = field.NewInt32(table, "project_id")
	m.LayerID = field.NewInt32(table, "layer_id")
	m.ZuluTypeID = field.NewInt32(table, "zulu_type_id")
	m.ZuluModeID = field.NewInt32(table, "zulu_mode_id")
	m.Geometry = field.NewString(table, "geometry")
	m.Sys = field.NewInt64(table, "sys")
	m.Color = field.NewString(table, "color")
	m.Angle = field.NewFloat64(table, "angle")
	m.DateCreateObj = field.NewTime(table, "date_create_obj")
	m.DateDeleteObj = field.NewTime(table, "date_delete_obj")
	m.ParentID = field.NewInt64(table, "parent_id")
	m.OldObjID = field.NewInt64(table, "old_obj_id")
	m.ExtID = field.NewInt64(table, "ext_id")
	m.DateLastChange = field.NewTime(table, "date_last_change")

	m.fillFieldMap()

	return m
}

func (m *mapObject) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *mapObject) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 15)
	m.fieldMap["id"] = m.ID
	m.fieldMap["project_id"] = m.ProjectID
	m.fieldMap["layer_id"] = m.LayerID
	m.fieldMap["zulu_type_id"] = m.ZuluTypeID
	m.fieldMap["zulu_mode_id"] = m.ZuluModeID
	m.fieldMap["geometry"] = m.Geometry
	m.fieldMap["sys"] = m.Sys
	m.fieldMap["color"] = m.Color
	m.fieldMap["angle"] = m.Angle
	m.fieldMap["date_create_obj"] = m.DateCreateObj
	m.fieldMap["date_delete_obj"] = m.DateDeleteObj
	m.fieldMap["parent_id"] = m.ParentID
	m.fieldMap["old_obj_id"] = m.OldObjID
	m.fieldMap["ext_id"] = m.ExtID
	m.fieldMap["date_last_change"] = m.DateLastChange
}

func (m mapObject) clone(db *gorm.DB) mapObject {
	m.mapObjectDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m mapObject) replaceDB(db *gorm.DB) mapObject {
	m.mapObjectDo.ReplaceDB(db)
	return m
}

type mapObjectDo struct{ gen.DO }

type IMapObjectDo interface {
	gen.SubQuery
	Debug() IMapObjectDo
	WithContext(ctx context.Context) IMapObjectDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMapObjectDo
	WriteDB() IMapObjectDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMapObjectDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMapObjectDo
	Not(conds ...gen.Condition) IMapObjectDo
	Or(conds ...gen.Condition) IMapObjectDo
	Select(conds ...field.Expr) IMapObjectDo
	Where(conds ...gen.Condition) IMapObjectDo
	Order(conds ...field.Expr) IMapObjectDo
	Distinct(cols ...field.Expr) IMapObjectDo
	Omit(cols ...field.Expr) IMapObjectDo
	Join(table schema.Tabler, on ...field.Expr) IMapObjectDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMapObjectDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMapObjectDo
	Group(cols ...field.Expr) IMapObjectDo
	Having(conds ...gen.Condition) IMapObjectDo
	Limit(limit int) IMapObjectDo
	Offset(offset int) IMapObjectDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMapObjectDo
	Unscoped() IMapObjectDo
	Create(values ...*model.MapObject) error
	CreateInBatches(values []*model.MapObject, batchSize int) error
	Save(values ...*model.MapObject) error
	First() (*model.MapObject, error)
	Take() (*model.MapObject, error)
	Last() (*model.MapObject, error)
	Find() ([]*model.MapObject, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MapObject, err error)
	FindInBatches(result *[]*model.MapObject, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MapObject) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMapObjectDo
	Assign(attrs ...field.AssignExpr) IMapObjectDo
	Joins(fields ...field.RelationField) IMapObjectDo
	Preload(fields ...field.RelationField) IMapObjectDo
	FirstOrInit() (*model.MapObject, error)
	FirstOrCreate() (*model.MapObject, error)
	FindByPage(offset int, limit int) (result []*model.MapObject, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMapObjectDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m mapObjectDo) Debug() IMapObjectDo {
	return m.withDO(m.DO.Debug())
}

func (m mapObjectDo) WithContext(ctx context.Context) IMapObjectDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m mapObjectDo) ReadDB() IMapObjectDo {
	return m.Clauses(dbresolver.Read)
}

func (m mapObjectDo) WriteDB() IMapObjectDo {
	return m.Clauses(dbresolver.Write)
}

func (m mapObjectDo) Session(config *gorm.Session) IMapObjectDo {
	return m.withDO(m.DO.Session(config))
}

func (m mapObjectDo) Clauses(conds ...clause.Expression) IMapObjectDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m mapObjectDo) Returning(value interface{}, columns ...string) IMapObjectDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m mapObjectDo) Not(conds ...gen.Condition) IMapObjectDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m mapObjectDo) Or(conds ...gen.Condition) IMapObjectDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m mapObjectDo) Select(conds ...field.Expr) IMapObjectDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m mapObjectDo) Where(conds ...gen.Condition) IMapObjectDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m mapObjectDo) Order(conds ...field.Expr) IMapObjectDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m mapObjectDo) Distinct(cols ...field.Expr) IMapObjectDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m mapObjectDo) Omit(cols ...field.Expr) IMapObjectDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m mapObjectDo) Join(table schema.Tabler, on ...field.Expr) IMapObjectDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m mapObjectDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMapObjectDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m mapObjectDo) RightJoin(table schema.Tabler, on ...field.Expr) IMapObjectDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m mapObjectDo) Group(cols ...field.Expr) IMapObjectDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m mapObjectDo) Having(conds ...gen.Condition) IMapObjectDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m mapObjectDo) Limit(limit int) IMapObjectDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m mapObjectDo) Offset(offset int) IMapObjectDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m mapObjectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMapObjectDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m mapObjectDo) Unscoped() IMapObjectDo {
	return m.withDO(m.DO.Unscoped())
}

func (m mapObjectDo) Create(values ...*model.MapObject) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m mapObjectDo) CreateInBatches(values []*model.MapObject, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m mapObjectDo) Save(values ...*model.MapObject) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m mapObjectDo) First() (*model.MapObject, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapObject), nil
	}
}

func (m mapObjectDo) Take() (*model.MapObject, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapObject), nil
	}
}

func (m mapObjectDo) Last() (*model.MapObject, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapObject), nil
	}
}

func (m mapObjectDo) Find() ([]*model.MapObject, error) {
	result, err := m.DO.Find()
	return result.([]*model.MapObject), err
}

func (m mapObjectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MapObject, err error) {
	buf := make([]*model.MapObject, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m mapObjectDo) FindInBatches(result *[]*model.MapObject, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m mapObjectDo) Attrs(attrs ...field.AssignExpr) IMapObjectDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m mapObjectDo) Assign(attrs ...field.AssignExpr) IMapObjectDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m mapObjectDo) Joins(fields ...field.RelationField) IMapObjectDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m mapObjectDo) Preload(fields ...field.RelationField) IMapObjectDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m mapObjectDo) FirstOrInit() (*model.MapObject, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapObject), nil
	}
}

func (m mapObjectDo) FirstOrCreate() (*model.MapObject, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapObject), nil
	}
}

func (m mapObjectDo) FindByPage(offset int, limit int) (result []*model.MapObject, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m mapObjectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m mapObjectDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m mapObjectDo) Delete(models ...*model.MapObject) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *mapObjectDo) withDO(do gen.Dao) *mapObjectDo {
	m.DO = *do.(*gen.DO)
	return m
}
