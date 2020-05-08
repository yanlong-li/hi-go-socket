package db

type builder struct {
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
}

//查询构造器
func (_b builder) Find() *queryBuilder {
	_XBuilder := new(queryBuilder)
	_XBuilder.builder = _b
	return _XBuilder
}

//插入构造器
func (_b builder) Insert() *insertBuilder {
	_XBuilder := new(insertBuilder)
	_XBuilder.builder = _b
	// 读取模型参数
	_XBuilder.modelFillArgs()
	return _XBuilder
}

//删除构造器
func (_b builder) Delete() *deleteBuilder {
	_XBuilder := new(deleteBuilder)
	_XBuilder.builder = _b
	return _XBuilder
}

func Model(model interface{}) *builder {
	orm := new(builder)
	orm.model = model
	orm.tableNames()
	return orm
}
