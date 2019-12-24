package db

func (query *queryBuilder) Limit(offset uint64, args ...uint64) *queryBuilder {

	query.limit = "LIMIT ?"
	query.args = append(query.args, offset)
	if len(args) >= 1 {
		query.limit += ",?"
		query.args = append(query.args, args[0])
	}
	query.limit += " "
	return query

}
