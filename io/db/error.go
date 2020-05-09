package db

import "database/sql"

type OrmError struct {
	Err error
}

func (_e OrmError) Empty() bool {

	if _e.Err == sql.ErrNoRows {
		return true
	}
	return false
}

func (_e OrmError) Status() bool {
	if _e.Err == nil {
		return true
	}
	return false
}
func (_e OrmError) Error() string {
	if _e.Err != nil {
		return _e.Err.Error()
	}
	return ""
}
