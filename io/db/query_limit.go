package db

func (query *queryBuilder) Limit(offset uint64, args ...uint64) *queryBuilder {

	query.limit = "LIMIT ?"
	query.queryArgs = append(query.queryArgs, offset)
	if len(args) >= 1 {
		query.limit += ",?"
		query.queryArgs = append(query.queryArgs, args[0])
	}
	query.limit += " "
	return query

}
