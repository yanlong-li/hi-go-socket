package db

func (_queryBuilder *queryBuilder) Limit(offset uint64, args ...uint64) *queryBuilder {

	_queryBuilder.limit = "LIMIT ?"
	_queryBuilder.builder.args = append(_queryBuilder.builder.args, offset)
	if len(args) >= 1 {
		_queryBuilder.limit += ",?"
		_queryBuilder.builder.args = append(_queryBuilder.builder.args, args[0])
	}
	_queryBuilder.limit += " "
	return _queryBuilder

}
