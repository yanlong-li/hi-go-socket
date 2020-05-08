package db

import (
	"reflect"
)

type insertBuilder struct {
	builder builder
}

func (query *insertBuilder) modelFillArgs() {
	p := reflect.ValueOf(query.builder.model).Elem()
	for i := 1; i < p.NumField(); i++ {
		f := p.Field(i)
		field2 := f.Interface()
		query.builder.args = append(query.builder.args, field2)
		query.builder.argsSql += "?,"
	}
	query.builder.argsSql = query.builder.argsSql[0 : len(query.builder.argsSql)-1]
}
