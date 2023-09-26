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

func newMapLayer(db *gorm.DB, opts ...gen.DOOption) mapLayer {
	_mapLayer := mapLayer{}

	_mapLayer.mapLayerDo.UseDB(db, opts...)
	_mapLayer.mapLayerDo.UseModel(&model.MapLayer{})

	tableName := _mapLayer.mapLayerDo.TableName()
	_mapLayer.ALL = field.NewAsterisk(tableName)
	_mapLayer.ID = field.NewInt32(tableName, "id")
	_mapLayer.DateCreate = field.NewTime(tableName, "date_create")
	_mapLayer.DateLastChange = field.NewTime(tableName, "date_last_change")
	_mapLayer.CreaterID = field.NewInt64(tableName, "creater_id")
	_mapLayer.ChangerID = field.NewInt64(tableName, "changer_id")
	_mapLayer.ProjectID = field.NewInt32(tableName, "project_id")
	_mapLayer.IsArchive = field.NewBool(tableName, "is_archive")
	_mapLayer.Title = field.NewString(tableName, "title")
	_mapLayer.PathLayer = field.NewString(tableName, "path_layer")
	_mapLayer.TypeLayerID = field.NewInt32(tableName, "type_layer_id")
	_mapLayer.IsVisible = field.NewBool(tableName, "is_visible")
	_mapLayer.ParentID = field.NewInt32(tableName, "parent_id")
	_mapLayer.TypeGeometry = field.NewString(tableName, "type_geometry")
	_mapLayer.ClassName = field.NewString(tableName, "class_name")
	_mapLayer.Style = field.NewString(tableName, "style")
	_mapLayer.ZoomViewFrom = field.NewInt16(tableName, "zoom_view_from")
	_mapLayer.ZoomViewTo = field.NewInt16(tableName, "zoom_view_to")
	_mapLayer.ZuluTypeID = field.NewInt32(tableName, "zulu_type_id")
	_mapLayer.ZuluModeID = field.NewInt32(tableName, "zulu_mode_id")
	_mapLayer.LibraryStyleID = field.NewInt32(tableName, "library_style_id")
	_mapLayer.Priority = field.NewInt32(tableName, "priority")
	_mapLayer.VectorMinZoom = field.NewInt32(tableName, "vector_min_zoom")
	_mapLayer.VectorMaxZoom = field.NewInt32(tableName, "vector_max_zoom")
	_mapLayer.IsTileVector = field.NewBool(tableName, "is_tile_vector")
	_mapLayer.ExtID = field.NewString(tableName, "ext_id")
	_mapLayer.VectorZoomFix = field.NewInt32(tableName, "vector_zoom_fix")
	_mapLayer.VectorZoomFixKoef = field.NewFloat64(tableName, "vector_zoom_fix_koef")
	_mapLayer.WmsZoomFix = field.NewInt32(tableName, "wms_zoom_fix")
	_mapLayer.WmsCacheSec = field.NewInt32(tableName, "wms_cache_sec")
	_mapLayer.HoverStyleID = field.NewInt32(tableName, "hover_style_id")
	_mapLayer.TemplateLayerID = field.NewInt32(tableName, "template_layer_id")

	_mapLayer.fillFieldMap()

	return _mapLayer
}

type mapLayer struct {
	mapLayerDo

	ALL               field.Asterisk
	ID                field.Int32
	DateCreate        field.Time
	DateLastChange    field.Time
	CreaterID         field.Int64
	ChangerID         field.Int64
	ProjectID         field.Int32
	IsArchive         field.Bool
	Title             field.String
	PathLayer         field.String
	TypeLayerID       field.Int32
	IsVisible         field.Bool
	ParentID          field.Int32
	TypeGeometry      field.String
	ClassName         field.String
	Style             field.String
	ZoomViewFrom      field.Int16
	ZoomViewTo        field.Int16
	ZuluTypeID        field.Int32
	ZuluModeID        field.Int32
	LibraryStyleID    field.Int32
	Priority          field.Int32
	VectorMinZoom     field.Int32
	VectorMaxZoom     field.Int32
	IsTileVector      field.Bool
	ExtID             field.String
	VectorZoomFix     field.Int32
	VectorZoomFixKoef field.Float64
	WmsZoomFix        field.Int32
	WmsCacheSec       field.Int32
	HoverStyleID      field.Int32
	TemplateLayerID   field.Int32

	fieldMap map[string]field.Expr
}

func (m mapLayer) Table(newTableName string) *mapLayer {
	m.mapLayerDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m mapLayer) As(alias string) *mapLayer {
	m.mapLayerDo.DO = *(m.mapLayerDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *mapLayer) updateTableName(table string) *mapLayer {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.DateCreate = field.NewTime(table, "date_create")
	m.DateLastChange = field.NewTime(table, "date_last_change")
	m.CreaterID = field.NewInt64(table, "creater_id")
	m.ChangerID = field.NewInt64(table, "changer_id")
	m.ProjectID = field.NewInt32(table, "project_id")
	m.IsArchive = field.NewBool(table, "is_archive")
	m.Title = field.NewString(table, "title")
	m.PathLayer = field.NewString(table, "path_layer")
	m.TypeLayerID = field.NewInt32(table, "type_layer_id")
	m.IsVisible = field.NewBool(table, "is_visible")
	m.ParentID = field.NewInt32(table, "parent_id")
	m.TypeGeometry = field.NewString(table, "type_geometry")
	m.ClassName = field.NewString(table, "class_name")
	m.Style = field.NewString(table, "style")
	m.ZoomViewFrom = field.NewInt16(table, "zoom_view_from")
	m.ZoomViewTo = field.NewInt16(table, "zoom_view_to")
	m.ZuluTypeID = field.NewInt32(table, "zulu_type_id")
	m.ZuluModeID = field.NewInt32(table, "zulu_mode_id")
	m.LibraryStyleID = field.NewInt32(table, "library_style_id")
	m.Priority = field.NewInt32(table, "priority")
	m.VectorMinZoom = field.NewInt32(table, "vector_min_zoom")
	m.VectorMaxZoom = field.NewInt32(table, "vector_max_zoom")
	m.IsTileVector = field.NewBool(table, "is_tile_vector")
	m.ExtID = field.NewString(table, "ext_id")
	m.VectorZoomFix = field.NewInt32(table, "vector_zoom_fix")
	m.VectorZoomFixKoef = field.NewFloat64(table, "vector_zoom_fix_koef")
	m.WmsZoomFix = field.NewInt32(table, "wms_zoom_fix")
	m.WmsCacheSec = field.NewInt32(table, "wms_cache_sec")
	m.HoverStyleID = field.NewInt32(table, "hover_style_id")
	m.TemplateLayerID = field.NewInt32(table, "template_layer_id")

	m.fillFieldMap()

	return m
}

func (m *mapLayer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *mapLayer) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 31)
	m.fieldMap["id"] = m.ID
	m.fieldMap["date_create"] = m.DateCreate
	m.fieldMap["date_last_change"] = m.DateLastChange
	m.fieldMap["creater_id"] = m.CreaterID
	m.fieldMap["changer_id"] = m.ChangerID
	m.fieldMap["project_id"] = m.ProjectID
	m.fieldMap["is_archive"] = m.IsArchive
	m.fieldMap["title"] = m.Title
	m.fieldMap["path_layer"] = m.PathLayer
	m.fieldMap["type_layer_id"] = m.TypeLayerID
	m.fieldMap["is_visible"] = m.IsVisible
	m.fieldMap["parent_id"] = m.ParentID
	m.fieldMap["type_geometry"] = m.TypeGeometry
	m.fieldMap["class_name"] = m.ClassName
	m.fieldMap["style"] = m.Style
	m.fieldMap["zoom_view_from"] = m.ZoomViewFrom
	m.fieldMap["zoom_view_to"] = m.ZoomViewTo
	m.fieldMap["zulu_type_id"] = m.ZuluTypeID
	m.fieldMap["zulu_mode_id"] = m.ZuluModeID
	m.fieldMap["library_style_id"] = m.LibraryStyleID
	m.fieldMap["priority"] = m.Priority
	m.fieldMap["vector_min_zoom"] = m.VectorMinZoom
	m.fieldMap["vector_max_zoom"] = m.VectorMaxZoom
	m.fieldMap["is_tile_vector"] = m.IsTileVector
	m.fieldMap["ext_id"] = m.ExtID
	m.fieldMap["vector_zoom_fix"] = m.VectorZoomFix
	m.fieldMap["vector_zoom_fix_koef"] = m.VectorZoomFixKoef
	m.fieldMap["wms_zoom_fix"] = m.WmsZoomFix
	m.fieldMap["wms_cache_sec"] = m.WmsCacheSec
	m.fieldMap["hover_style_id"] = m.HoverStyleID
	m.fieldMap["template_layer_id"] = m.TemplateLayerID
}

func (m mapLayer) clone(db *gorm.DB) mapLayer {
	m.mapLayerDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m mapLayer) replaceDB(db *gorm.DB) mapLayer {
	m.mapLayerDo.ReplaceDB(db)
	return m
}

type mapLayerDo struct{ gen.DO }

type IMapLayerDo interface {
	gen.SubQuery
	Debug() IMapLayerDo
	WithContext(ctx context.Context) IMapLayerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMapLayerDo
	WriteDB() IMapLayerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMapLayerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMapLayerDo
	Not(conds ...gen.Condition) IMapLayerDo
	Or(conds ...gen.Condition) IMapLayerDo
	Select(conds ...field.Expr) IMapLayerDo
	Where(conds ...gen.Condition) IMapLayerDo
	Order(conds ...field.Expr) IMapLayerDo
	Distinct(cols ...field.Expr) IMapLayerDo
	Omit(cols ...field.Expr) IMapLayerDo
	Join(table schema.Tabler, on ...field.Expr) IMapLayerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMapLayerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMapLayerDo
	Group(cols ...field.Expr) IMapLayerDo
	Having(conds ...gen.Condition) IMapLayerDo
	Limit(limit int) IMapLayerDo
	Offset(offset int) IMapLayerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMapLayerDo
	Unscoped() IMapLayerDo
	Create(values ...*model.MapLayer) error
	CreateInBatches(values []*model.MapLayer, batchSize int) error
	Save(values ...*model.MapLayer) error
	First() (*model.MapLayer, error)
	Take() (*model.MapLayer, error)
	Last() (*model.MapLayer, error)
	Find() ([]*model.MapLayer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MapLayer, err error)
	FindInBatches(result *[]*model.MapLayer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MapLayer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMapLayerDo
	Assign(attrs ...field.AssignExpr) IMapLayerDo
	Joins(fields ...field.RelationField) IMapLayerDo
	Preload(fields ...field.RelationField) IMapLayerDo
	FirstOrInit() (*model.MapLayer, error)
	FirstOrCreate() (*model.MapLayer, error)
	FindByPage(offset int, limit int) (result []*model.MapLayer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMapLayerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m mapLayerDo) Debug() IMapLayerDo {
	return m.withDO(m.DO.Debug())
}

func (m mapLayerDo) WithContext(ctx context.Context) IMapLayerDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m mapLayerDo) ReadDB() IMapLayerDo {
	return m.Clauses(dbresolver.Read)
}

func (m mapLayerDo) WriteDB() IMapLayerDo {
	return m.Clauses(dbresolver.Write)
}

func (m mapLayerDo) Session(config *gorm.Session) IMapLayerDo {
	return m.withDO(m.DO.Session(config))
}

func (m mapLayerDo) Clauses(conds ...clause.Expression) IMapLayerDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m mapLayerDo) Returning(value interface{}, columns ...string) IMapLayerDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m mapLayerDo) Not(conds ...gen.Condition) IMapLayerDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m mapLayerDo) Or(conds ...gen.Condition) IMapLayerDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m mapLayerDo) Select(conds ...field.Expr) IMapLayerDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m mapLayerDo) Where(conds ...gen.Condition) IMapLayerDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m mapLayerDo) Order(conds ...field.Expr) IMapLayerDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m mapLayerDo) Distinct(cols ...field.Expr) IMapLayerDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m mapLayerDo) Omit(cols ...field.Expr) IMapLayerDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m mapLayerDo) Join(table schema.Tabler, on ...field.Expr) IMapLayerDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m mapLayerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMapLayerDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m mapLayerDo) RightJoin(table schema.Tabler, on ...field.Expr) IMapLayerDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m mapLayerDo) Group(cols ...field.Expr) IMapLayerDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m mapLayerDo) Having(conds ...gen.Condition) IMapLayerDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m mapLayerDo) Limit(limit int) IMapLayerDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m mapLayerDo) Offset(offset int) IMapLayerDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m mapLayerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMapLayerDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m mapLayerDo) Unscoped() IMapLayerDo {
	return m.withDO(m.DO.Unscoped())
}

func (m mapLayerDo) Create(values ...*model.MapLayer) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m mapLayerDo) CreateInBatches(values []*model.MapLayer, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m mapLayerDo) Save(values ...*model.MapLayer) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m mapLayerDo) First() (*model.MapLayer, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapLayer), nil
	}
}

func (m mapLayerDo) Take() (*model.MapLayer, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapLayer), nil
	}
}

func (m mapLayerDo) Last() (*model.MapLayer, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapLayer), nil
	}
}

func (m mapLayerDo) Find() ([]*model.MapLayer, error) {
	result, err := m.DO.Find()
	return result.([]*model.MapLayer), err
}

func (m mapLayerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MapLayer, err error) {
	buf := make([]*model.MapLayer, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m mapLayerDo) FindInBatches(result *[]*model.MapLayer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m mapLayerDo) Attrs(attrs ...field.AssignExpr) IMapLayerDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m mapLayerDo) Assign(attrs ...field.AssignExpr) IMapLayerDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m mapLayerDo) Joins(fields ...field.RelationField) IMapLayerDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m mapLayerDo) Preload(fields ...field.RelationField) IMapLayerDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m mapLayerDo) FirstOrInit() (*model.MapLayer, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapLayer), nil
	}
}

func (m mapLayerDo) FirstOrCreate() (*model.MapLayer, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MapLayer), nil
	}
}

func (m mapLayerDo) FindByPage(offset int, limit int) (result []*model.MapLayer, count int64, err error) {
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

func (m mapLayerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m mapLayerDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m mapLayerDo) Delete(models ...*model.MapLayer) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *mapLayerDo) withDO(do gen.Dao) *mapLayerDo {
	m.DO = *do.(*gen.DO)
	return m
}
