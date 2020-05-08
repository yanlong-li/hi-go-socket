package db

import "fmt"

// 返回 SQL
func (_query *queryBuilder) Sql() string {

	return fmt.Sprintf("SELECT %s FROM %s %s %s", _query.builder.getFields(), _query.builder.table, _query.builder.getWhere(), _query.limit)
}

// 返回 SQL
func (query *insertBuilder) Sql() string {

	return fmt.Sprintf("INSERT INTO %s(%s) value(%s)", query.builder.table, query.builder.getNotListFields("id"), query.builder.argsSql)
}

// 返回 SQL
func (_delete *deleteBuilder) Sql() string {

	return fmt.Sprintf("DELETE FROM %s %s", _delete.builder.table, _delete.builder.getWhere())
}
