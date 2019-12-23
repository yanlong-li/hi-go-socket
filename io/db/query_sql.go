package db

import "fmt"

// 返回 SQL
func (query *queryBuilder) Sql() string {

	return fmt.Sprintf("SELECT %s FROM %s %s %s", query.getFields(), query.table, query.getWhere(), query.limit)
}
