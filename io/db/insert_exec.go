package db

import (
	"reflect"
)

func (_insertBuilder *insertBuilder) Insert() OrmError {

	result, err := db.Exec(_insertBuilder.Sql(), _insertBuilder.builder.args...)
	if err != nil {
		return OrmError{err: err}
	}

	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		return OrmError{err: err}
	}
	p := reflect.ValueOf(_insertBuilder.builder.model).Elem()
	f := p.Field(0)
	f.SetUint(uint64(lastInsertID))
	return OrmError{err: nil}
}
