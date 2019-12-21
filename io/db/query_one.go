package db

import (
	"reflect"
)

// 查询单条
func (orm *QueryBuilder) One() {
	// 准备查询字段
	//todo 准备查询条件
	row := db.QueryRow(orm.Sql(), orm.whereArgs...)

	refs := refs(orm.model)
	_ = row.Scan(refs...)
	orm.row(refs)
}

// 处理单条记录
func (orm *QueryBuilder) row(refs []interface{}) {

	v9 := reflect.ValueOf(orm.model)
	v10 := reflect.Indirect(v9)
	for k, _ := range orm.fields {
		v11 := v10.Field(k)
		unmarshalConverter(v11, refs[k])
	}
}
