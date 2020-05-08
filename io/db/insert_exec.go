package db

import "reflect"

func (query *insertBuilder) Insert() error {

	result, err := db.Exec(query.Sql(), query.builder.args...)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		return err
	}
	p := reflect.ValueOf(query.builder.model).Elem()
	f := p.Field(0)
	f.SetUint(uint64(lastInsertID))
	return nil
}
