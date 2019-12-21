package db

import (
	"log"
	"reflect"
)

// 批量查询
func (orm *QueryBuilder) All() []interface{} {
	rows, err := db.Query(orm.Sql(), orm.whereArgs...)
	if err != nil {
		log.Panic(err)
	}
	refs := refs(orm.model)
	for rows.Next() {
		_ = rows.Scan(refs...)
		orm.rows(refs)
	}

	return orm.models
}

// 处理批量查询结果
func (orm *QueryBuilder) rows(refs []interface{}) {

	v9 := reflect.TypeOf(orm.model).Elem()
	v10 := reflect.Indirect(reflect.New(v9))
	for k, _ := range orm.fields {
		v11 := v10.Field(k)
		unmarshalConverter(v11, refs[k])
	}
	orm.models = append(orm.models, v10.Interface())
}
