package db

// 套一层查询的 where

// In
func (_queryBuilder *queryBuilder) WhereIn(field string, list []interface{}) *queryBuilder {
	_queryBuilder.builder.whereIn(field, list)
	return _queryBuilder
}

// Not In
func (_queryBuilder *queryBuilder) WhereNotIn(field string, list []interface{}) *queryBuilder {
	_queryBuilder.builder.whereNotIn(field, list)
	return _queryBuilder
}

// 区间
func (_queryBuilder *queryBuilder) WhereBetween(field string, value1, value2 interface{}) *queryBuilder {
	_queryBuilder.builder.whereBetween(field, value1, value2)
	return _queryBuilder
}

// 非区间
func (_queryBuilder *queryBuilder) WhereNotBetween(field string, value1, value2 interface{}) *queryBuilder {
	_queryBuilder.builder.whereNotBetween(field, value1, value2)
	return _queryBuilder
}

// or
func (_queryBuilder *queryBuilder) OrWhere(args ...interface{}) *queryBuilder {
	_queryBuilder.builder.orWhere(args...)
	return _queryBuilder
}

// and
func (_queryBuilder *queryBuilder) AndWhere(args ...interface{}) *queryBuilder {
	_queryBuilder.builder.andWhere(args...)
	return _queryBuilder
}

// 默认 and
func (_queryBuilder *queryBuilder) Where(args ...interface{}) *queryBuilder {
	_queryBuilder.builder.where(args...)
	return _queryBuilder
}
