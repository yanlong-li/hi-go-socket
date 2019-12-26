package db

import (
	"reflect"
)

func (query *queryBuilder) Insert() error {
	query.modelFillArgs()
	result, err := db.Exec(query.InsertSql(), query.args...)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		return err
	}
	p := reflect.ValueOf(query.model).Elem()
	f := p.Field(0)
	f.SetUint(uint64(lastInsertID))
	return nil
}

func (query *queryBuilder) modelFillArgs() {
	p := reflect.ValueOf(query.model).Elem()
	for i := 1; i < p.NumField(); i++ {
		f := p.Field(i)
		field2 := f.Interface()
		query.args = append(query.args, field2)
		query.argsSql += "?,"
	}
	query.argsSql = query.argsSql[0 : len(query.argsSql)-1]
}
