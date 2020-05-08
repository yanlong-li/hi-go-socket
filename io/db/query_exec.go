package db

import (
	"fmt"
	"log"
	"reflect"
)

// 查询单条
func (_query *queryBuilder) One() error {
	// 准备查询字段
	row := db.QueryRow(_query.Sql(), _query.builder.args...)
	fmt.Println(_query.Sql(), _query.builder.args)

	refs := refs(_query.builder.model)
	err := row.Scan(refs...)
	_query.row(refs)
	return err
}

// 处理单条记录
func (_query *queryBuilder) row(refs []interface{}) {

	v9 := reflect.ValueOf(_query.builder.model)
	v10 := reflect.Indirect(v9)
	for k := range _query.builder.fields {
		v11 := v10.Field(k)
		unmarshalConverter(v11, refs[k])
	}
}

// 查询单条
func (_query *queryBuilder) Exists() bool {
	if _query.One() != nil {
		return false
	}
	return true
}

// 批量查询
func (_query *queryBuilder) All() []interface{} {
	rows, err := db.Query(_query.Sql(), _query.builder.args...)
	if err != nil {
		log.Panic(err)
	}
	refs := refs(_query.builder.model)
	for rows.Next() {
		_ = rows.Scan(refs...)
		_query.rows(refs)
	}

	return _query.builder.models
}

// 处理批量查询结果
func (_query *queryBuilder) rows(refs []interface{}) {

	v9 := reflect.TypeOf(_query.builder.model).Elem()
	v10 := reflect.Indirect(reflect.New(v9))
	for k := range _query.builder.fields {
		v11 := v10.Field(k)
		unmarshalConverter(v11, refs[k])
	}
	_query.builder.models = append(_query.builder.models, v10.Interface())
}
