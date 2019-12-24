package db

import (
	"fmt"
	"reflect"
)

// 查询单条
func (query *queryBuilder) One() error {
	// 准备查询字段
	row := db.QueryRow(query.SelectSql(), query.args...)
	fmt.Println(query.SelectSql(), query.args)

	refs := refs(query.model)
	err := row.Scan(refs...)
	query.row(refs)
	return err
}

// 处理单条记录
func (query *queryBuilder) row(refs []interface{}) {

	v9 := reflect.ValueOf(query.model)
	v10 := reflect.Indirect(v9)
	for k, _ := range query.fields {
		v11 := v10.Field(k)
		unmarshalConverter(v11, refs[k])
	}
}

// 查询单条
func (query *queryBuilder) Exists() bool {
	if query.One() != nil {
		return false
	}
	return true
}
