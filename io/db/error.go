package db

import "database/sql"

type OrmError struct {
	err error
}

func (_e OrmError) Empty() bool {

	if _e.err == sql.ErrNoRows {
		return true
	}
	return false
}
