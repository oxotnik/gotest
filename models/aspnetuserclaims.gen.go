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

func newAspNetUserClaim(db *gorm.DB, opts ...gen.DOOption) aspNetUserClaim {
	_aspNetUserClaim := aspNetUserClaim{}

	_aspNetUserClaim.aspNetUserClaimDo.UseDB(db, opts...)
	_aspNetUserClaim.aspNetUserClaimDo.UseModel(&model.AspNetUserClaim{})

	tableName := _aspNetUserClaim.aspNetUserClaimDo.TableName()
	_aspNetUserClaim.ALL = field.NewAsterisk(tableName)
	_aspNetUserClaim.ID = field.NewInt32(tableName, "Id")
	_aspNetUserClaim.UserID = field.NewString(tableName, "UserId")
	_aspNetUserClaim.ClaimType = field.NewString(tableName, "ClaimType")
	_aspNetUserClaim.ClaimValue = field.NewString(tableName, "ClaimValue")

	_aspNetUserClaim.fillFieldMap()

	return _aspNetUserClaim
}

type aspNetUserClaim struct {
	aspNetUserClaimDo

	ALL        field.Asterisk
	ID         field.Int32
	UserID     field.String
	ClaimType  field.String
	ClaimValue field.String

	fieldMap map[string]field.Expr
}

func (a aspNetUserClaim) Table(newTableName string) *aspNetUserClaim {
	a.aspNetUserClaimDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a aspNetUserClaim) As(alias string) *aspNetUserClaim {
	a.aspNetUserClaimDo.DO = *(a.aspNetUserClaimDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *aspNetUserClaim) updateTableName(table string) *aspNetUserClaim {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "Id")
	a.UserID = field.NewString(table, "UserId")
	a.ClaimType = field.NewString(table, "ClaimType")
	a.ClaimValue = field.NewString(table, "ClaimValue")

	a.fillFieldMap()

	return a
}

func (a *aspNetUserClaim) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *aspNetUserClaim) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["Id"] = a.ID
	a.fieldMap["UserId"] = a.UserID
	a.fieldMap["ClaimType"] = a.ClaimType
	a.fieldMap["ClaimValue"] = a.ClaimValue
}

func (a aspNetUserClaim) clone(db *gorm.DB) aspNetUserClaim {
	a.aspNetUserClaimDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a aspNetUserClaim) replaceDB(db *gorm.DB) aspNetUserClaim {
	a.aspNetUserClaimDo.ReplaceDB(db)
	return a
}

type aspNetUserClaimDo struct{ gen.DO }

type IAspNetUserClaimDo interface {
	gen.SubQuery
	Debug() IAspNetUserClaimDo
	WithContext(ctx context.Context) IAspNetUserClaimDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAspNetUserClaimDo
	WriteDB() IAspNetUserClaimDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAspNetUserClaimDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAspNetUserClaimDo
	Not(conds ...gen.Condition) IAspNetUserClaimDo
	Or(conds ...gen.Condition) IAspNetUserClaimDo
	Select(conds ...field.Expr) IAspNetUserClaimDo
	Where(conds ...gen.Condition) IAspNetUserClaimDo
	Order(conds ...field.Expr) IAspNetUserClaimDo
	Distinct(cols ...field.Expr) IAspNetUserClaimDo
	Omit(cols ...field.Expr) IAspNetUserClaimDo
	Join(table schema.Tabler, on ...field.Expr) IAspNetUserClaimDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAspNetUserClaimDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAspNetUserClaimDo
	Group(cols ...field.Expr) IAspNetUserClaimDo
	Having(conds ...gen.Condition) IAspNetUserClaimDo
	Limit(limit int) IAspNetUserClaimDo
	Offset(offset int) IAspNetUserClaimDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAspNetUserClaimDo
	Unscoped() IAspNetUserClaimDo
	Create(values ...*model.AspNetUserClaim) error
	CreateInBatches(values []*model.AspNetUserClaim, batchSize int) error
	Save(values ...*model.AspNetUserClaim) error
	First() (*model.AspNetUserClaim, error)
	Take() (*model.AspNetUserClaim, error)
	Last() (*model.AspNetUserClaim, error)
	Find() ([]*model.AspNetUserClaim, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AspNetUserClaim, err error)
	FindInBatches(result *[]*model.AspNetUserClaim, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AspNetUserClaim) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAspNetUserClaimDo
	Assign(attrs ...field.AssignExpr) IAspNetUserClaimDo
	Joins(fields ...field.RelationField) IAspNetUserClaimDo
	Preload(fields ...field.RelationField) IAspNetUserClaimDo
	FirstOrInit() (*model.AspNetUserClaim, error)
	FirstOrCreate() (*model.AspNetUserClaim, error)
	FindByPage(offset int, limit int) (result []*model.AspNetUserClaim, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAspNetUserClaimDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a aspNetUserClaimDo) Debug() IAspNetUserClaimDo {
	return a.withDO(a.DO.Debug())
}

func (a aspNetUserClaimDo) WithContext(ctx context.Context) IAspNetUserClaimDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aspNetUserClaimDo) ReadDB() IAspNetUserClaimDo {
	return a.Clauses(dbresolver.Read)
}

func (a aspNetUserClaimDo) WriteDB() IAspNetUserClaimDo {
	return a.Clauses(dbresolver.Write)
}

func (a aspNetUserClaimDo) Session(config *gorm.Session) IAspNetUserClaimDo {
	return a.withDO(a.DO.Session(config))
}

func (a aspNetUserClaimDo) Clauses(conds ...clause.Expression) IAspNetUserClaimDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aspNetUserClaimDo) Returning(value interface{}, columns ...string) IAspNetUserClaimDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aspNetUserClaimDo) Not(conds ...gen.Condition) IAspNetUserClaimDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aspNetUserClaimDo) Or(conds ...gen.Condition) IAspNetUserClaimDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aspNetUserClaimDo) Select(conds ...field.Expr) IAspNetUserClaimDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aspNetUserClaimDo) Where(conds ...gen.Condition) IAspNetUserClaimDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aspNetUserClaimDo) Order(conds ...field.Expr) IAspNetUserClaimDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aspNetUserClaimDo) Distinct(cols ...field.Expr) IAspNetUserClaimDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aspNetUserClaimDo) Omit(cols ...field.Expr) IAspNetUserClaimDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aspNetUserClaimDo) Join(table schema.Tabler, on ...field.Expr) IAspNetUserClaimDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aspNetUserClaimDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAspNetUserClaimDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aspNetUserClaimDo) RightJoin(table schema.Tabler, on ...field.Expr) IAspNetUserClaimDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aspNetUserClaimDo) Group(cols ...field.Expr) IAspNetUserClaimDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aspNetUserClaimDo) Having(conds ...gen.Condition) IAspNetUserClaimDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aspNetUserClaimDo) Limit(limit int) IAspNetUserClaimDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aspNetUserClaimDo) Offset(offset int) IAspNetUserClaimDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aspNetUserClaimDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAspNetUserClaimDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aspNetUserClaimDo) Unscoped() IAspNetUserClaimDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aspNetUserClaimDo) Create(values ...*model.AspNetUserClaim) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aspNetUserClaimDo) CreateInBatches(values []*model.AspNetUserClaim, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aspNetUserClaimDo) Save(values ...*model.AspNetUserClaim) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aspNetUserClaimDo) First() (*model.AspNetUserClaim, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AspNetUserClaim), nil
	}
}

func (a aspNetUserClaimDo) Take() (*model.AspNetUserClaim, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AspNetUserClaim), nil
	}
}

func (a aspNetUserClaimDo) Last() (*model.AspNetUserClaim, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AspNetUserClaim), nil
	}
}

func (a aspNetUserClaimDo) Find() ([]*model.AspNetUserClaim, error) {
	result, err := a.DO.Find()
	return result.([]*model.AspNetUserClaim), err
}

func (a aspNetUserClaimDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AspNetUserClaim, err error) {
	buf := make([]*model.AspNetUserClaim, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aspNetUserClaimDo) FindInBatches(result *[]*model.AspNetUserClaim, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aspNetUserClaimDo) Attrs(attrs ...field.AssignExpr) IAspNetUserClaimDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aspNetUserClaimDo) Assign(attrs ...field.AssignExpr) IAspNetUserClaimDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aspNetUserClaimDo) Joins(fields ...field.RelationField) IAspNetUserClaimDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aspNetUserClaimDo) Preload(fields ...field.RelationField) IAspNetUserClaimDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aspNetUserClaimDo) FirstOrInit() (*model.AspNetUserClaim, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AspNetUserClaim), nil
	}
}

func (a aspNetUserClaimDo) FirstOrCreate() (*model.AspNetUserClaim, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AspNetUserClaim), nil
	}
}

func (a aspNetUserClaimDo) FindByPage(offset int, limit int) (result []*model.AspNetUserClaim, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a aspNetUserClaimDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aspNetUserClaimDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aspNetUserClaimDo) Delete(models ...*model.AspNetUserClaim) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aspNetUserClaimDo) withDO(do gen.Dao) *aspNetUserClaimDo {
	a.DO = *do.(*gen.DO)
	return a
}
