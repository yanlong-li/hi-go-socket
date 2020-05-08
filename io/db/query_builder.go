package db

import "fmt"

type queryBuilder struct {
	builder builder

	limit string
	order string
}

const (
	SortAsc  = "ASC"
	SortDesc = "DESC"
)

func (_queryBuilder *queryBuilder) OrderBy(field, sort string) *queryBuilder {

	_queryBuilder.order = fmt.Sprintf("ORDER BY `%s` %s ", field, sort)
	return _queryBuilder

}

func (_queryBuilder *queryBuilder) OrderByMap(args map[string]string) *queryBuilder {

	_orderByArgs := ""
	for k, v := range args {
		_orderByArgs += fmt.Sprintf("`%s` %s,", k, v)
	}
	_orderByArgs = _orderByArgs[:len(_orderByArgs)-1]

	_queryBuilder.order = fmt.Sprintf("ORDER BY %s ", _orderByArgs)
	return _queryBuilder

}
