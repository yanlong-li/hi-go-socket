package db

type QueryBuilder struct {
	// 模型
	model interface{}
	// 模型组
	models []interface{}
	// 查询条件
	where        string
	whereArgs    []interface{}
	fields       []string
	selectFields string
	// 表名
	table string
}

func Find(model interface{}) *QueryBuilder {
	orm := new(QueryBuilder)
	orm.model = model
	orm.initField()
	orm.tableNames()
	return orm
}
