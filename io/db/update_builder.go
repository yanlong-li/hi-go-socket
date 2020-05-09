package db

type updateBuilder struct {
	builder builder
}

func (_updateBuilder *updateBuilder) Update() (uint64, OrmError) {

	result, err := db.Exec(_updateBuilder.Sql(), _updateBuilder.builder.args...)
	if err != nil {
		return 0, OrmError{err: err}
	}
	ra, err := result.RowsAffected()
	return uint64(ra), OrmError{err: err}
}
