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

func newMapSetting(db *gorm.DB, opts ...gen.DOOption) mapSetting {
	_mapSetting := mapSetting{}

	_mapSetting.mapSettingDo.UseDB(db, opts...)
	_mapSetting.mapSettingDo.UseModel(&model.MapSetting{})

	tableName := _mapSetting.mapSettingDo.TableName()
	_mapSetting.ALL = field.NewAsterisk(tableName)
	_mapSetting.ID = field.NewInt32(tableName, "id")
	_mapSetting.DateCreate = field.NewTime(tableName, "date_create")
	_mapSetting.DateLastChange = field.NewTime(tableName, "date_last_change")
	_mapSetting.CreaterID = field.NewInt64(tableName, "creater_id")
	_mapSetting.ChangerID = field.NewInt64(tableName, "changer_id")
	_mapSetting.ProjectID = field.NewInt32(tableName, "project_id")
	_mapSetting.Title = field.NewString(tableName, "title")
	_mapSetting.Key = field.NewString(tableName, "key")
	_mapSetting.Value = field.NewString(tableName, "value")
	_mapSetting.JSONValue = field.NewString(tableName, "json_value")
	_mapSetting.Priority = field.NewInt32(tableName, "priority")

	_mapSetting.fillFieldMap()

	return _mapSetting
}

type mapSetting struct {
	mapSettingDo

	ALL            field.Asterisk
	ID             field.Int32
	DateCreate     field.Time
	DateLastChange field.Time
	CreaterID      field.Int64
	ChangerID      field.Int64
	ProjectID      field.Int32
	Title          field.String
	Key            field.String
	Value          field.String
	JSONValue      field.String
	Priority       field.Int32

	fieldMap map[string]field.Expr
}

func (m mapSetting) Table(newTableName string) *mapSetting {
	m.mapSettingDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m mapSetting) As(alias string) *mapSetting {
	m.mapSettingDo.DO = *(m.mapSettingDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *mapSetting) updateTableName(table string) *mapSetting {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.DateCreate = field.NewTime(table, "date_create")
	m.DateLastChange = field.NewTime(table, "date_last_change")
	m.CreaterID = field.NewInt64(table, "creater_id")
	m.ChangerID = field.NewInt64(table, "changer_id")
	m.ProjectID = field.NewInt32(table, "project_id")
	m.Title = field.NewString(table, "title")
	m.Key = field.NewString(table, "key")
	m.Value = field.NewString(table, "value")
	m.JSONValue = field.NewString(table, "json_value")
	m.Priority = field.NewInt32(table, "priority")

	m.fillFieldMap()

	return m
}

func (m *mapSetting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *mapSetting) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 11)
	m.fieldMap["id"] = m.ID
	m.fieldMap["date_create"] = m.DateCreate
	m.fieldMap["date_last_change"] = m.DateLastChange
	m.fieldMap["creater_id"] = m.CreaterID
	m.fieldMap["changer_id"] = m.ChangerID
	m.fieldMap["project_id"] = m.ProjectID
	m.fieldMap["title"] = m.Title
	m.fieldMap["key"] = m.Key
	m.fieldMap["value"] = m.Value
	m.fieldMap["json_value"] = m.JSONValue
	m.fieldMap["priority"] = m.Priority
}

func (m mapSetting) clone(db *gorm.DB) mapSetting {
	m.mapSettingDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m mapSetting) replaceDB(db *gorm.DB) mapSetting {
	m.mapSettingDo.ReplaceDB(db)
	return m
}

type mapSettingDo struct{ gen.DO }

type IMapSettingDo interface {
	gen.SubQuery
	Debug() IMapSettingDo
	WithContext(ctx context.Context) IMapSettingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMapSettingDo
	WriteDB() IMapSettingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMapSettingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMapSettingDo
	Not(conds ...gen.Condition) IMapSettingDo
	Or(conds ...gen.Condition) IMapSettingDo
	Select(conds ...field.Expr) IMapSettingDo
	Where(conds ...gen.Condition) IMapSettingDo
	Order(conds ...field.Expr) IMapSettingDo
	Distinct(cols ...field.Expr) IMapSettingDo
	Omit(cols ...field.Expr) IMapSettingDo
	Join(table schema.Tabler, on ...field.Expr) IMapSettingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMapSettingDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMapSettingDo
	Group(cols ...field.Expr) IMapSettingDo
	Having(conds ...gen.Condition) IMapSettingDo
	Limit(limit int) IMapSettingDo
	Offset(offset int) IMapSettingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMapSettingDo
	Unscoped() IMapSettingDo
	Create(values ...*model.MapSetting) error
	CreateInBatches(values []*model.MapSetting, batchSize int) error
	Save(values ...*model.MapSetting) error
	First() (*model.MapSetting, error)
	Take() (*model.MapSetting, error)
	Last() (*model.MapSetting, error)
	Find() ([]*model.MapSetting, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MapSetting, err error)
	FindInBatches(result *[]*model.MapSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MapSetting) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMapSettingDo
	Assign(attrs ...field.AssignExpr) IMapSettingDo
	Joins(fields ...field.RelationField) IMapSettingDo
	Preload(fields ...field.RelationField) IMapSettingDo
	FirstOrInit() (*model.MapSetting, error)
	FirstOrCreate() (*model.MapSetting, error)
	FindByPage(offset int, limit int) (result []*model.MapSetting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMapSettingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m mapSettingDo) Debug() IMapSettingDo {
	return m.withDO(m.DO.Debug())
}

func (m mapSettingDo) WithContext(ctx context.Context) IMapSettingDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m mapSettingDo) ReadDB() IMapSettingDo {
	return m.Clauses(dbresolver.Read)
}

func (m mapSettingDo) WriteDB() IMapSettingDo {
	return m.Clauses(dbresolver.Write)
}

func (m mapSettingDo) Session(config *gorm.Session) IMapSettingDo {
	return m.withDO(m.DO.Session(config))
}

func (m mapSettingDo) Clauses(conds ...clause.Expression) IMapSettingDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m mapSettingDo) Returning(value interface{}, columns ...string) IMapSettingDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m mapSettingDo) Not(conds ...gen.Condition) IMapSettingDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m mapSettingDo) Or(conds ...gen.Condition) IMapSettingDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m mapSettingDo) Select(conds ...field.Expr) IMapSettingDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m mapSettingDo) Where(conds ...gen.Condition) IMapSettingDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m mapSettingDo) Order(conds ...field.Expr) IMapSettingDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m mapSettingDo) Distinct(cols ...field.Expr) IMapSettingDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m mapSettingDo) Omit(cols ...field.Expr) IMapSettingDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m mapSettingDo) Join(table schema.Tabler, on ...field.Expr) IMapSettingDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m mapSettingDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMapSettingDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m mapSettingDo) RightJoin(table schema.Tabler, on ...field.Expr) IMapSettingDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m mapSettingDo) Group(cols ...field.Expr) IMapSettingDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m mapSettingDo) Having(conds ...gen.Condition) IMapSettingDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m mapSettingDo) Limit(limit int) IMapSettingDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m mapSettingDo) Offset(offset int) IMapSettingDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m mapSettingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMapSettingDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m mapSettingDo) Unscoped() IMapSettingDo {
	return m.withDO(m.DO.Unscoped())
}

func (m mapSettingDo) Create(values ...*model.MapSetting) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m mapSettingDo) CreateInBatches(values []*model.MapSetting, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m mapSettingDo) Save(values ...*model.MapSetting) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m mapSettingDo) First() (*model.MapSetting, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapSetting), nil
	}
}

func (m mapSettingDo) Take() (*model.MapSetting, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapSetting), nil
	}
}

func (m mapSettingDo) Last() (*model.MapSetting, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapSetting), nil
	}
}

func (m mapSettingDo) Find() ([]*model.MapSetting, error) {
	result, err := m.DO.Find()
	return result.([]*model.MapSetting), err
}

func (m mapSettingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MapSetting, err error) {
	buf := make([]*model.MapSetting, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m mapSettingDo) FindInBatches(result *[]*model.MapSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m mapSettingDo) Attrs(attrs ...field.AssignExpr) IMapSettingDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m mapSettingDo) Assign(attrs ...field.AssignExpr) IMapSettingDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m mapSettingDo) Joins(fields ...field.RelationField) IMapSettingDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m mapSettingDo) Preload(fields ...field.RelationField) IMapSettingDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m mapSettingDo) FirstOrInit() (*model.MapSetting, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapSetting), nil
	}
}

func (m mapSettingDo) FirstOrCreate() (*model.MapSetting, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapSetting), nil
	}
}

func (m mapSettingDo) FindByPage(offset int, limit int) (result []*model.MapSetting, count int64, err error) {
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

func (m mapSettingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m mapSettingDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m mapSettingDo) Delete(models ...*model.MapSetting) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *mapSettingDo) withDO(do gen.Dao) *mapSettingDo {
	m.DO = *do.(*gen.DO)
	return m
}