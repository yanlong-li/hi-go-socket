package db

type queryBuilder struct {
	// 模型
	model interface{}
	// 模型组
	models []interface{}
	// 查询条件
	where     string
	queryArgs []interface{}
	fields    []string
	// 表名
	table string

	limit string
}

func Find(model interface{}) *queryBuilder {
	orm := new(queryBuilder)
	orm.model = model
	orm.tableNames()
	return orm
}
