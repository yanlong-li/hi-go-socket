package db

// 套一层更新的 where

// In
func (_updateBuilder *updateBuilder) WhereIn(field string, list []interface{}) *updateBuilder {
	_updateBuilder.builder.whereIn(field, list)
	return _updateBuilder
}

// Not In
func (_updateBuilder *updateBuilder) WhereNotIn(field string, list []interface{}) *updateBuilder {
	_updateBuilder.builder.whereNotIn(field, list)
	return _updateBuilder
}

// 区间
func (_updateBuilder *updateBuilder) WhereBetween(field string, value1, value2 interface{}) *updateBuilder {
	_updateBuilder.builder.whereBetween(field, value1, value2)
	return _updateBuilder
}

// 非区间
func (_updateBuilder *updateBuilder) WhereNotBetween(field string, value1, value2 interface{}) *updateBuilder {
	_updateBuilder.builder.whereNotBetween(field, value1, value2)
	return _updateBuilder
}

// or
func (_updateBuilder *updateBuilder) OrWhere(args ...interface{}) *updateBuilder {
	_updateBuilder.builder.orWhere(args...)
	return _updateBuilder
}

// and
func (_updateBuilder *updateBuilder) AndWhere(args ...interface{}) *updateBuilder {
	_updateBuilder.builder.andWhere(args...)
	return _updateBuilder
}

// 默认 and
func (_updateBuilder *updateBuilder) Where(args ...interface{}) *updateBuilder {
	_updateBuilder.builder.where(args...)
	return _updateBuilder
}
