package db

import "errors"

func (_deleteBuilder *deleteBuilder) Delete() (uint64, OrmError) {

	if len(_deleteBuilder.builder.getWhere()) == 0 {
		return 0, OrmError{
			Err: errors.New("删除条件不能为空"),
		}
	}

	result, err := db.Exec(_deleteBuilder.Sql(), _deleteBuilder.builder.args...)
	if err != nil {
		return 0, OrmError{
			Err: err,
		}
	}
	ra, err := result.RowsAffected()
	return uint64(ra), OrmError{
		Err: err,
	}
}
