package db

import "errors"

func (_delete *deleteBuilder) Delete() (uint64, error) {

	if len(_delete.builder.getWhere()) == 0 {
		return 0, errors.New("删除条件不能为空")
	}

	result, err := db.Exec(_delete.Sql(), _delete.builder.args...)
	if err != nil {
		return 0, err
	}
	ra, err := result.RowsAffected()
	return uint64(ra), err
}
