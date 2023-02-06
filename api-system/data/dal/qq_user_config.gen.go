// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/axiangcoding/antonstar-bot/data/table"
)

func newQQUserConfig(db *gorm.DB, opts ...gen.DOOption) qQUserConfig {
	_qQUserConfig := qQUserConfig{}

	_qQUserConfig.qQUserConfigDo.UseDB(db, opts...)
	_qQUserConfig.qQUserConfigDo.UseModel(&table.QQUserConfig{})

	tableName := _qQUserConfig.qQUserConfigDo.TableName()
	_qQUserConfig.ALL = field.NewAsterisk(tableName)
	_qQUserConfig.ID = field.NewUint(tableName, "id")
	_qQUserConfig.CreatedAt = field.NewTime(tableName, "created_at")
	_qQUserConfig.UpdatedAt = field.NewTime(tableName, "updated_at")
	_qQUserConfig.DeletedAt = field.NewField(tableName, "deleted_at")
	_qQUserConfig.UserId = field.NewInt64(tableName, "user_id")
	_qQUserConfig.Banned = field.NewBool(tableName, "banned")
	_qQUserConfig.TodayQueryCount = field.NewInt(tableName, "today_query_count")
	_qQUserConfig.OneDayQueryLimit = field.NewInt(tableName, "one_day_query_limit")
	_qQUserConfig.TotalQueryCount = field.NewInt(tableName, "total_query_count")
	_qQUserConfig.TodayUsageCount = field.NewInt(tableName, "today_usage_count")
	_qQUserConfig.OneDayUsageLimit = field.NewInt(tableName, "one_day_usage_limit")
	_qQUserConfig.TotalUsageCount = field.NewInt(tableName, "total_usage_count")
	_qQUserConfig.Admin = field.NewBool(tableName, "admin")
	_qQUserConfig.SuperAdmin = field.NewBool(tableName, "super_admin")
	_qQUserConfig.BindingGameNick = field.NewString(tableName, "binding_game_nick")

	_qQUserConfig.fillFieldMap()

	return _qQUserConfig
}

type qQUserConfig struct {
	qQUserConfigDo

	ALL              field.Asterisk
	ID               field.Uint
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field
	UserId           field.Int64
	Banned           field.Bool
	TodayQueryCount  field.Int
	OneDayQueryLimit field.Int
	TotalQueryCount  field.Int
	TodayUsageCount  field.Int
	OneDayUsageLimit field.Int
	TotalUsageCount  field.Int
	Admin            field.Bool
	SuperAdmin       field.Bool
	BindingGameNick  field.String

	fieldMap map[string]field.Expr
}

func (q qQUserConfig) Table(newTableName string) *qQUserConfig {
	q.qQUserConfigDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q qQUserConfig) As(alias string) *qQUserConfig {
	q.qQUserConfigDo.DO = *(q.qQUserConfigDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *qQUserConfig) updateTableName(table string) *qQUserConfig {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewUint(table, "id")
	q.CreatedAt = field.NewTime(table, "created_at")
	q.UpdatedAt = field.NewTime(table, "updated_at")
	q.DeletedAt = field.NewField(table, "deleted_at")
	q.UserId = field.NewInt64(table, "user_id")
	q.Banned = field.NewBool(table, "banned")
	q.TodayQueryCount = field.NewInt(table, "today_query_count")
	q.OneDayQueryLimit = field.NewInt(table, "one_day_query_limit")
	q.TotalQueryCount = field.NewInt(table, "total_query_count")
	q.TodayUsageCount = field.NewInt(table, "today_usage_count")
	q.OneDayUsageLimit = field.NewInt(table, "one_day_usage_limit")
	q.TotalUsageCount = field.NewInt(table, "total_usage_count")
	q.Admin = field.NewBool(table, "admin")
	q.SuperAdmin = field.NewBool(table, "super_admin")
	q.BindingGameNick = field.NewString(table, "binding_game_nick")

	q.fillFieldMap()

	return q
}

func (q *qQUserConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *qQUserConfig) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 15)
	q.fieldMap["id"] = q.ID
	q.fieldMap["created_at"] = q.CreatedAt
	q.fieldMap["updated_at"] = q.UpdatedAt
	q.fieldMap["deleted_at"] = q.DeletedAt
	q.fieldMap["user_id"] = q.UserId
	q.fieldMap["banned"] = q.Banned
	q.fieldMap["today_query_count"] = q.TodayQueryCount
	q.fieldMap["one_day_query_limit"] = q.OneDayQueryLimit
	q.fieldMap["total_query_count"] = q.TotalQueryCount
	q.fieldMap["today_usage_count"] = q.TodayUsageCount
	q.fieldMap["one_day_usage_limit"] = q.OneDayUsageLimit
	q.fieldMap["total_usage_count"] = q.TotalUsageCount
	q.fieldMap["admin"] = q.Admin
	q.fieldMap["super_admin"] = q.SuperAdmin
	q.fieldMap["binding_game_nick"] = q.BindingGameNick
}

func (q qQUserConfig) clone(db *gorm.DB) qQUserConfig {
	q.qQUserConfigDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q qQUserConfig) replaceDB(db *gorm.DB) qQUserConfig {
	q.qQUserConfigDo.ReplaceDB(db)
	return q
}

type qQUserConfigDo struct{ gen.DO }

type IQQUserConfigDo interface {
	gen.SubQuery
	Debug() IQQUserConfigDo
	WithContext(ctx context.Context) IQQUserConfigDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IQQUserConfigDo
	WriteDB() IQQUserConfigDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IQQUserConfigDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IQQUserConfigDo
	Not(conds ...gen.Condition) IQQUserConfigDo
	Or(conds ...gen.Condition) IQQUserConfigDo
	Select(conds ...field.Expr) IQQUserConfigDo
	Where(conds ...gen.Condition) IQQUserConfigDo
	Order(conds ...field.Expr) IQQUserConfigDo
	Distinct(cols ...field.Expr) IQQUserConfigDo
	Omit(cols ...field.Expr) IQQUserConfigDo
	Join(table schema.Tabler, on ...field.Expr) IQQUserConfigDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IQQUserConfigDo
	RightJoin(table schema.Tabler, on ...field.Expr) IQQUserConfigDo
	Group(cols ...field.Expr) IQQUserConfigDo
	Having(conds ...gen.Condition) IQQUserConfigDo
	Limit(limit int) IQQUserConfigDo
	Offset(offset int) IQQUserConfigDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IQQUserConfigDo
	Unscoped() IQQUserConfigDo
	Create(values ...*table.QQUserConfig) error
	CreateInBatches(values []*table.QQUserConfig, batchSize int) error
	Save(values ...*table.QQUserConfig) error
	First() (*table.QQUserConfig, error)
	Take() (*table.QQUserConfig, error)
	Last() (*table.QQUserConfig, error)
	Find() ([]*table.QQUserConfig, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.QQUserConfig, err error)
	FindInBatches(result *[]*table.QQUserConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.QQUserConfig) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IQQUserConfigDo
	Assign(attrs ...field.AssignExpr) IQQUserConfigDo
	Joins(fields ...field.RelationField) IQQUserConfigDo
	Preload(fields ...field.RelationField) IQQUserConfigDo
	FirstOrInit() (*table.QQUserConfig, error)
	FirstOrCreate() (*table.QQUserConfig, error)
	FindByPage(offset int, limit int) (result []*table.QQUserConfig, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IQQUserConfigDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (q qQUserConfigDo) Debug() IQQUserConfigDo {
	return q.withDO(q.DO.Debug())
}

func (q qQUserConfigDo) WithContext(ctx context.Context) IQQUserConfigDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q qQUserConfigDo) ReadDB() IQQUserConfigDo {
	return q.Clauses(dbresolver.Read)
}

func (q qQUserConfigDo) WriteDB() IQQUserConfigDo {
	return q.Clauses(dbresolver.Write)
}

func (q qQUserConfigDo) Session(config *gorm.Session) IQQUserConfigDo {
	return q.withDO(q.DO.Session(config))
}

func (q qQUserConfigDo) Clauses(conds ...clause.Expression) IQQUserConfigDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q qQUserConfigDo) Returning(value interface{}, columns ...string) IQQUserConfigDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q qQUserConfigDo) Not(conds ...gen.Condition) IQQUserConfigDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q qQUserConfigDo) Or(conds ...gen.Condition) IQQUserConfigDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q qQUserConfigDo) Select(conds ...field.Expr) IQQUserConfigDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q qQUserConfigDo) Where(conds ...gen.Condition) IQQUserConfigDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q qQUserConfigDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IQQUserConfigDo {
	return q.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (q qQUserConfigDo) Order(conds ...field.Expr) IQQUserConfigDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q qQUserConfigDo) Distinct(cols ...field.Expr) IQQUserConfigDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q qQUserConfigDo) Omit(cols ...field.Expr) IQQUserConfigDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q qQUserConfigDo) Join(table schema.Tabler, on ...field.Expr) IQQUserConfigDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q qQUserConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) IQQUserConfigDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q qQUserConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) IQQUserConfigDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q qQUserConfigDo) Group(cols ...field.Expr) IQQUserConfigDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q qQUserConfigDo) Having(conds ...gen.Condition) IQQUserConfigDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q qQUserConfigDo) Limit(limit int) IQQUserConfigDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q qQUserConfigDo) Offset(offset int) IQQUserConfigDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q qQUserConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IQQUserConfigDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q qQUserConfigDo) Unscoped() IQQUserConfigDo {
	return q.withDO(q.DO.Unscoped())
}

func (q qQUserConfigDo) Create(values ...*table.QQUserConfig) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q qQUserConfigDo) CreateInBatches(values []*table.QQUserConfig, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q qQUserConfigDo) Save(values ...*table.QQUserConfig) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q qQUserConfigDo) First() (*table.QQUserConfig, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.QQUserConfig), nil
	}
}

func (q qQUserConfigDo) Take() (*table.QQUserConfig, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.QQUserConfig), nil
	}
}

func (q qQUserConfigDo) Last() (*table.QQUserConfig, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.QQUserConfig), nil
	}
}

func (q qQUserConfigDo) Find() ([]*table.QQUserConfig, error) {
	result, err := q.DO.Find()
	return result.([]*table.QQUserConfig), err
}

func (q qQUserConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.QQUserConfig, err error) {
	buf := make([]*table.QQUserConfig, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q qQUserConfigDo) FindInBatches(result *[]*table.QQUserConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q qQUserConfigDo) Attrs(attrs ...field.AssignExpr) IQQUserConfigDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q qQUserConfigDo) Assign(attrs ...field.AssignExpr) IQQUserConfigDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q qQUserConfigDo) Joins(fields ...field.RelationField) IQQUserConfigDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q qQUserConfigDo) Preload(fields ...field.RelationField) IQQUserConfigDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q qQUserConfigDo) FirstOrInit() (*table.QQUserConfig, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.QQUserConfig), nil
	}
}

func (q qQUserConfigDo) FirstOrCreate() (*table.QQUserConfig, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.QQUserConfig), nil
	}
}

func (q qQUserConfigDo) FindByPage(offset int, limit int) (result []*table.QQUserConfig, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q qQUserConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q qQUserConfigDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q qQUserConfigDo) Delete(models ...*table.QQUserConfig) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *qQUserConfigDo) withDO(do gen.Dao) *qQUserConfigDo {
	q.DO = *do.(*gen.DO)
	return q
}
