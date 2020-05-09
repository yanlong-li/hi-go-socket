package db

import (
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"reflect"
)

// 查询单条
func (_queryBuilder *queryBuilder) One() OrmError {
	// 准备查询字段
	row := db.QueryRow(_queryBuilder.Sql(), _queryBuilder.builder.args...)
	logger.Debug("执行SQL:"+_queryBuilder.Sql(), 0, _queryBuilder.builder.args)

	refs := refs(_queryBuilder.builder.model)
	err := row.Scan(refs...)
	_queryBuilder.row(refs)
	return OrmError{Err: err}
}

// 处理单条记录
func (_queryBuilder *queryBuilder) row(refs []interface{}) {

	v9 := reflect.ValueOf(_queryBuilder.builder.model)
	v10 := reflect.Indirect(v9)
	for k := range _queryBuilder.builder.fields {
		v11 := v10.Field(k)
		unmarshalConverter(v11, refs[k])
	}
}

// 查询单条
func (_queryBuilder *queryBuilder) Exists() bool {
	if _queryBuilder.One().Empty() {
		return false
	}
	return true
}

// 批量查询
func (_queryBuilder *queryBuilder) All() []interface{} {
	rows, err := db.Query(_queryBuilder.Sql(), _queryBuilder.builder.args...)
	logger.Debug("执行SQL:"+_queryBuilder.Sql(), 0, _queryBuilder.builder.args)
	if err != nil {
		logger.Fatal("database Find ALL error", 0, err.Error())
	}
	refs := refs(_queryBuilder.builder.model)
	for rows.Next() {
		err = rows.Scan(refs...)
		if err != nil {
			logger.Fatal("database Scan ALL error", 0, err.Error())
		}
		_queryBuilder.rows(refs)
	}

	return _queryBuilder.builder.models
}

// 处理批量查询结果
func (_queryBuilder *queryBuilder) rows(refs []interface{}) {

	v9 := reflect.TypeOf(_queryBuilder.builder.model).Elem()
	v10 := reflect.Indirect(reflect.New(v9))
	for k := range _queryBuilder.builder.fields {
		v11 := v10.Field(k)
		unmarshalConverter(v11, refs[k])
	}
	_queryBuilder.builder.models = append(_queryBuilder.builder.models, v10.Interface())
}
