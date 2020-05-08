package db

// 套一层删除的 where

// In
func (_deleteBuilder *deleteBuilder) WhereIn(field string, list []interface{}) *deleteBuilder {
	_deleteBuilder.builder.whereIn(field, list)
	return _deleteBuilder
}

// Not In
func (_deleteBuilder *deleteBuilder) WhereNotIn(field string, list []interface{}) *deleteBuilder {
	_deleteBuilder.builder.whereNotIn(field, list)
	return _deleteBuilder
}

// 区间
func (_deleteBuilder *deleteBuilder) WhereBetween(field string, value1, value2 interface{}) *deleteBuilder {
	_deleteBuilder.builder.whereBetween(field, value1, value2)
	return _deleteBuilder
}

// 非区间
func (_deleteBuilder *deleteBuilder) WhereNotBetween(field string, value1, value2 interface{}) *deleteBuilder {
	_deleteBuilder.builder.whereNotBetween(field, value1, value2)
	return _deleteBuilder
}

// or
func (_deleteBuilder *deleteBuilder) OrWhere(args ...interface{}) *deleteBuilder {
	_deleteBuilder.builder.orWhere(args...)
	return _deleteBuilder
}

// and
func (_deleteBuilder *deleteBuilder) AndWhere(args ...interface{}) *deleteBuilder {
	_deleteBuilder.builder.andWhere(args...)
	return _deleteBuilder
}

// 默认 and
func (_deleteBuilder *deleteBuilder) Where(args ...interface{}) *deleteBuilder {
	_deleteBuilder.builder.where(args...)
	return _deleteBuilder
}
