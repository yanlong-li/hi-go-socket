package db

type queryBuilder struct {
	// 模型
	model interface{}
	// 模型组
	models []interface{}
	// 查询条件 或 插入数据 的占位符
	argsSql string
	args    []interface{}
	fields  []string
	// 表名
	table string

	limit string
}

func Model(model interface{}) *queryBuilder {
	orm := new(queryBuilder)
	orm.model = model
	orm.tableNames()
	return orm
}
