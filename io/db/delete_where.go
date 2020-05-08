package db

// 套一层删除的 where

// In
func (_delete *deleteBuilder) WhereIn(field string, list []interface{}) *deleteBuilder {
	_delete.builder.whereIn(field, list)
	return _delete
}

// Not In
func (_delete *deleteBuilder) WhereNotIn(field string, list []interface{}) *deleteBuilder {
	_delete.builder.whereNotIn(field, list)
	return _delete
}

// 区间
func (_delete *deleteBuilder) WhereBetween(field string, value1, value2 interface{}) *deleteBuilder {
	_delete.builder.whereBetween(field, value1, value2)
	return _delete
}

// 非区间
func (_delete *deleteBuilder) WhereNotBetween(field string, value1, value2 interface{}) *deleteBuilder {
	_delete.builder.whereNotBetween(field, value1, value2)
	return _delete
}

// or
func (_delete *deleteBuilder) OrWhere(args ...interface{}) *deleteBuilder {
	_delete.builder.orWhere(args...)
	return _delete
}

// and
func (_delete *deleteBuilder) AndWhere(args ...interface{}) *deleteBuilder {
	_delete.builder.andWhere(args...)
	return _delete
}

// 默认 and
func (_delete *deleteBuilder) Where(args ...interface{}) *deleteBuilder {
	_delete.builder.where(args...)
	return _delete
}
