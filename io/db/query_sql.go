package db

import "fmt"

// 返回 SQL
func (query *queryBuilder) SelectSql() string {

	return fmt.Sprintf("SELECT %s FROM %s %s %s", query.getFields(), query.table, query.getWhere(), query.limit)
}

// 返回 SQL
func (query *queryBuilder) InsertSql() string {

	return fmt.Sprintf("INSERT INTO %s(%s) value(%s)", query.table, query.getNotListFields("id"), query.argsSql)
}

// 返回 SQL
func (query *queryBuilder) deleteSql() string {

	return fmt.Sprintf("DELETE FROM %s %s", query.table, query.getWhere())
}
