package db

import (
	"fmt"
)

// 返回 SQL
func (_queryBuilder *queryBuilder) Sql() string {

	return fmt.Sprintf("SELECT %s FROM `%s` %s %s %s", _queryBuilder.builder.getFields(), _queryBuilder.builder.table, _queryBuilder.builder.getWhere(), _queryBuilder.order, _queryBuilder.limit)
}

// 返回 SQL
func (_insertBuilder *insertBuilder) Sql() string {
	_value2StrArgs := value2StrArgs{Link: ",", Left: "`", Right: "`"}
	return fmt.Sprintf("INSERT INTO `%s`(%s) value(%s)", _insertBuilder.builder.table, _insertBuilder.builder.getNotListFields(_value2StrArgs, _insertBuilder.builder.pk), _insertBuilder.builder.argsSql)
}

// 返回 SQL
func (_deleteBuilder *deleteBuilder) Sql() string {

	return fmt.Sprintf("DELETE FROM `%s` %s", _deleteBuilder.builder.table, _deleteBuilder.builder.getWhere())
}

// 返回 SQL
func (_updateBuilder *updateBuilder) Sql() string {
	// update {table} {set} {where}

	if len(_updateBuilder.builder.getWhere()) == 0 {
		// 主键的值
		_updateBuilder.Where("=", _updateBuilder.builder.pk, _updateBuilder.builder.pkValue)
	}

	_value2StrArgs := value2StrArgs{Link: ",", Left: "`", Right: "` = ?"}
	return fmt.Sprintf("UPDATE `%s` SET %s %s", _updateBuilder.builder.table, _updateBuilder.builder.getNotListFields(_value2StrArgs, _updateBuilder.builder.pk), _updateBuilder.builder.getWhere())
}
