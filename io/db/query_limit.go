package db

func (_query *queryBuilder) Limit(offset uint64, args ...uint64) *queryBuilder {

	_query.limit = "LIMIT ?"
	_query.builder.args = append(_query.builder.args, offset)
	if len(args) >= 1 {
		_query.limit += ",?"
		_query.builder.args = append(_query.builder.args, args[0])
	}
	_query.limit += " "
	return _query

}
