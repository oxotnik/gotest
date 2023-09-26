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

func newLibraryIcon(db *gorm.DB, opts ...gen.DOOption) libraryIcon {
	_libraryIcon := libraryIcon{}

	_libraryIcon.libraryIconDo.UseDB(db, opts...)
	_libraryIcon.libraryIconDo.UseModel(&model.LibraryIcon{})

	tableName := _libraryIcon.libraryIconDo.TableName()
	_libraryIcon.ALL = field.NewAsterisk(tableName)
	_libraryIcon.ID = field.NewInt32(tableName, "id")
	_libraryIcon.DateCreate = field.NewTime(tableName, "date_create")
	_libraryIcon.DateLastChange = field.NewTime(tableName, "date_last_change")
	_libraryIcon.CreaterID = field.NewInt64(tableName, "creater_id")
	_libraryIcon.ChangerID = field.NewInt64(tableName, "changer_id")
	_libraryIcon.Title = field.NewString(tableName, "title")
	_libraryIcon.HashTag = field.NewString(tableName, "hash_tag")
	_libraryIcon.ObjType = field.NewString(tableName, "obj_type")
	_libraryIcon.FullPath = field.NewString(tableName, "full_path")
	_libraryIcon.FileName = field.NewString(tableName, "file_name")
	_libraryIcon.Extension = field.NewString(tableName, "extension")
	_libraryIcon.Size = field.NewInt64(tableName, "size")
	_libraryIcon.FileType = field.NewString(tableName, "file_type")
	_libraryIcon.ThumbnailURL = field.NewString(tableName, "thumbnail_url")
	_libraryIcon.URL = field.NewString(tableName, "url")
	_libraryIcon.IsArchive = field.NewBool(tableName, "is_archive")

	_libraryIcon.fillFieldMap()

	return _libraryIcon
}

type libraryIcon struct {
	libraryIconDo

	ALL            field.Asterisk
	ID             field.Int32
	DateCreate     field.Time
	DateLastChange field.Time
	CreaterID      field.Int64
	ChangerID      field.Int64
	Title          field.String
	HashTag        field.String
	ObjType        field.String
	FullPath       field.String
	FileName       field.String
	Extension      field.String
	Size           field.Int64
	FileType       field.String
	ThumbnailURL   field.String
	URL            field.String
	IsArchive      field.Bool

	fieldMap map[string]field.Expr
}

func (l libraryIcon) Table(newTableName string) *libraryIcon {
	l.libraryIconDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l libraryIcon) As(alias string) *libraryIcon {
	l.libraryIconDo.DO = *(l.libraryIconDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *libraryIcon) updateTableName(table string) *libraryIcon {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt32(table, "id")
	l.DateCreate = field.NewTime(table, "date_create")
	l.DateLastChange = field.NewTime(table, "date_last_change")
	l.CreaterID = field.NewInt64(table, "creater_id")
	l.ChangerID = field.NewInt64(table, "changer_id")
	l.Title = field.NewString(table, "title")
	l.HashTag = field.NewString(table, "hash_tag")
	l.ObjType = field.NewString(table, "obj_type")
	l.FullPath = field.NewString(table, "full_path")
	l.FileName = field.NewString(table, "file_name")
	l.Extension = field.NewString(table, "extension")
	l.Size = field.NewInt64(table, "size")
	l.FileType = field.NewString(table, "file_type")
	l.ThumbnailURL = field.NewString(table, "thumbnail_url")
	l.URL = field.NewString(table, "url")
	l.IsArchive = field.NewBool(table, "is_archive")

	l.fillFieldMap()

	return l
}

func (l *libraryIcon) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *libraryIcon) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 16)
	l.fieldMap["id"] = l.ID
	l.fieldMap["date_create"] = l.DateCreate
	l.fieldMap["date_last_change"] = l.DateLastChange
	l.fieldMap["creater_id"] = l.CreaterID
	l.fieldMap["changer_id"] = l.ChangerID
	l.fieldMap["title"] = l.Title
	l.fieldMap["hash_tag"] = l.HashTag
	l.fieldMap["obj_type"] = l.ObjType
	l.fieldMap["full_path"] = l.FullPath
	l.fieldMap["file_name"] = l.FileName
	l.fieldMap["extension"] = l.Extension
	l.fieldMap["size"] = l.Size
	l.fieldMap["file_type"] = l.FileType
	l.fieldMap["thumbnail_url"] = l.ThumbnailURL
	l.fieldMap["url"] = l.URL
	l.fieldMap["is_archive"] = l.IsArchive
}

func (l libraryIcon) clone(db *gorm.DB) libraryIcon {
	l.libraryIconDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l libraryIcon) replaceDB(db *gorm.DB) libraryIcon {
	l.libraryIconDo.ReplaceDB(db)
	return l
}

type libraryIconDo struct{ gen.DO }

type ILibraryIconDo interface {
	gen.SubQuery
	Debug() ILibraryIconDo
	WithContext(ctx context.Context) ILibraryIconDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILibraryIconDo
	WriteDB() ILibraryIconDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILibraryIconDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILibraryIconDo
	Not(conds ...gen.Condition) ILibraryIconDo
	Or(conds ...gen.Condition) ILibraryIconDo
	Select(conds ...field.Expr) ILibraryIconDo
	Where(conds ...gen.Condition) ILibraryIconDo
	Order(conds ...field.Expr) ILibraryIconDo
	Distinct(cols ...field.Expr) ILibraryIconDo
	Omit(cols ...field.Expr) ILibraryIconDo
	Join(table schema.Tabler, on ...field.Expr) ILibraryIconDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILibraryIconDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILibraryIconDo
	Group(cols ...field.Expr) ILibraryIconDo
	Having(conds ...gen.Condition) ILibraryIconDo
	Limit(limit int) ILibraryIconDo
	Offset(offset int) ILibraryIconDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILibraryIconDo
	Unscoped() ILibraryIconDo
	Create(values ...*model.LibraryIcon) error
	CreateInBatches(values []*model.LibraryIcon, batchSize int) error
	Save(values ...*model.LibraryIcon) error
	First() (*model.LibraryIcon, error)
	Take() (*model.LibraryIcon, error)
	Last() (*model.LibraryIcon, error)
	Find() ([]*model.LibraryIcon, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LibraryIcon, err error)
	FindInBatches(result *[]*model.LibraryIcon, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.LibraryIcon) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILibraryIconDo
	Assign(attrs ...field.AssignExpr) ILibraryIconDo
	Joins(fields ...field.RelationField) ILibraryIconDo
	Preload(fields ...field.RelationField) ILibraryIconDo
	FirstOrInit() (*model.LibraryIcon, error)
	FirstOrCreate() (*model.LibraryIcon, error)
	FindByPage(offset int, limit int) (result []*model.LibraryIcon, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILibraryIconDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l libraryIconDo) Debug() ILibraryIconDo {
	return l.withDO(l.DO.Debug())
}

func (l libraryIconDo) WithContext(ctx context.Context) ILibraryIconDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l libraryIconDo) ReadDB() ILibraryIconDo {
	return l.Clauses(dbresolver.Read)
}

func (l libraryIconDo) WriteDB() ILibraryIconDo {
	return l.Clauses(dbresolver.Write)
}

func (l libraryIconDo) Session(config *gorm.Session) ILibraryIconDo {
	return l.withDO(l.DO.Session(config))
}

func (l libraryIconDo) Clauses(conds ...clause.Expression) ILibraryIconDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l libraryIconDo) Returning(value interface{}, columns ...string) ILibraryIconDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l libraryIconDo) Not(conds ...gen.Condition) ILibraryIconDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l libraryIconDo) Or(conds ...gen.Condition) ILibraryIconDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l libraryIconDo) Select(conds ...field.Expr) ILibraryIconDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l libraryIconDo) Where(conds ...gen.Condition) ILibraryIconDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l libraryIconDo) Order(conds ...field.Expr) ILibraryIconDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l libraryIconDo) Distinct(cols ...field.Expr) ILibraryIconDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l libraryIconDo) Omit(cols ...field.Expr) ILibraryIconDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l libraryIconDo) Join(table schema.Tabler, on ...field.Expr) ILibraryIconDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l libraryIconDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILibraryIconDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l libraryIconDo) RightJoin(table schema.Tabler, on ...field.Expr) ILibraryIconDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l libraryIconDo) Group(cols ...field.Expr) ILibraryIconDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l libraryIconDo) Having(conds ...gen.Condition) ILibraryIconDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l libraryIconDo) Limit(limit int) ILibraryIconDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l libraryIconDo) Offset(offset int) ILibraryIconDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l libraryIconDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILibraryIconDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l libraryIconDo) Unscoped() ILibraryIconDo {
	return l.withDO(l.DO.Unscoped())
}

func (l libraryIconDo) Create(values ...*model.LibraryIcon) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l libraryIconDo) CreateInBatches(values []*model.LibraryIcon, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l libraryIconDo) Save(values ...*model.LibraryIcon) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l libraryIconDo) First() (*model.LibraryIcon, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LibraryIcon), nil
	}
}

func (l libraryIconDo) Take() (*model.LibraryIcon, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LibraryIcon), nil
	}
}

func (l libraryIconDo) Last() (*model.LibraryIcon, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LibraryIcon), nil
	}
}

func (l libraryIconDo) Find() ([]*model.LibraryIcon, error) {
	result, err := l.DO.Find()
	return result.([]*model.LibraryIcon), err
}

func (l libraryIconDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LibraryIcon, err error) {
	buf := make([]*model.LibraryIcon, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l libraryIconDo) FindInBatches(result *[]*model.LibraryIcon, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l libraryIconDo) Attrs(attrs ...field.AssignExpr) ILibraryIconDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l libraryIconDo) Assign(attrs ...field.AssignExpr) ILibraryIconDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l libraryIconDo) Joins(fields ...field.RelationField) ILibraryIconDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l libraryIconDo) Preload(fields ...field.RelationField) ILibraryIconDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l libraryIconDo) FirstOrInit() (*model.LibraryIcon, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LibraryIcon), nil
	}
}

func (l libraryIconDo) FirstOrCreate() (*model.LibraryIcon, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LibraryIcon), nil
	}
}

func (l libraryIconDo) FindByPage(offset int, limit int) (result []*model.LibraryIcon, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l libraryIconDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l libraryIconDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l libraryIconDo) Delete(models ...*model.LibraryIcon) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *libraryIconDo) withDO(do gen.Dao) *libraryIconDo {
	l.DO = *do.(*gen.DO)
	return l
}