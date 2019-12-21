package db

import "fmt"

// 返回 SQL
func (orm *QueryBuilder) Sql() string {

	return fmt.Sprintf("SELECT %s FROM %s %s", orm.selectFields, orm.table, orm.getWhere())
}
