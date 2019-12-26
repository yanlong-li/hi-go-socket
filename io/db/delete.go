package db

import "errors"

func (query *queryBuilder) Delete() (uint64, error) {

	if len(query.getWhere()) == 0 {
		return 0, errors.New("删除条件不能为空")
	}

	result, err := db.Exec(query.deleteSql(), query.args...)
	if err != nil {
		return 0, err
	}
	ra, err := result.RowsAffected()
	return uint64(ra), err
}
