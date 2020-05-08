package db

// 套一层查询的 where

// In
func (_query *queryBuilder) WhereIn(field string, list []interface{}) *queryBuilder {
	_query.builder.whereIn(field, list)
	return _query
}

// Not In
func (_query *queryBuilder) WhereNotIn(field string, list []interface{}) *queryBuilder {
	_query.builder.whereNotIn(field, list)
	return _query
}

// 区间
func (_query *queryBuilder) WhereBetween(field string, value1, value2 interface{}) *queryBuilder {
	_query.builder.whereBetween(field, value1, value2)
	return _query
}

// 非区间
func (_query *queryBuilder) WhereNotBetween(field string, value1, value2 interface{}) *queryBuilder {
	_query.builder.whereNotBetween(field, value1, value2)
	return _query
}

// or
func (_query *queryBuilder) OrWhere(args ...interface{}) *queryBuilder {
	_query.builder.orWhere(args...)
	return _query
}

// and
func (_query *queryBuilder) AndWhere(args ...interface{}) *queryBuilder {
	_query.builder.andWhere(args...)
	return _query
}

// 默认 and
func (_query *queryBuilder) Where(args ...interface{}) *queryBuilder {
	_query.builder.where(args...)
	return _query
}
