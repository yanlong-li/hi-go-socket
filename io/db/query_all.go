package db

import (
	"log"
	"reflect"
)

// 批量查询
func (query *queryBuilder) All() []interface{} {
	rows, err := db.Query(query.Sql(), query.queryArgs...)
	if err != nil {
		log.Panic(err)
	}
	refs := refs(query.model)
	for rows.Next() {
		_ = rows.Scan(refs...)
		query.rows(refs)
	}

	return query.models
}

// 处理批量查询结果
func (query *queryBuilder) rows(refs []interface{}) {

	v9 := reflect.TypeOf(query.model).Elem()
	v10 := reflect.Indirect(reflect.New(v9))
	for k, _ := range query.fields {
		v11 := v10.Field(k)
		unmarshalConverter(v11, refs[k])
	}
	query.models = append(query.models, v10.Interface())
}
