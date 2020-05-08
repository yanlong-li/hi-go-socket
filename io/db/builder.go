package db

import "reflect"

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
	// 主键 默认为 id
	pk string
	// 主键值
	pkValue interface{}
}

//查询构造器
func (_builder builder) Find() *queryBuilder {
	_XBuilder := new(queryBuilder)
	_XBuilder.builder = _builder
	// 增加默认排序 以主键顺序排序
	_XBuilder.OrderBy(_XBuilder.builder.pk, SortAsc)
	return _XBuilder
}

//插入构造器
func (_builder builder) Insert() *insertBuilder {
	_XBuilder := new(insertBuilder)
	_XBuilder.builder = _builder
	// 读取模型参数
	_XBuilder.builder.modelFillArgs(true)
	return _XBuilder
}

//删除构造器
func (_builder builder) Delete() *deleteBuilder {
	_XBuilder := new(deleteBuilder)
	_XBuilder.builder = _builder
	return _XBuilder
}

//更新构造器
func (_builder builder) Update() *updateBuilder {
	_XBuilder := new(updateBuilder)
	_XBuilder.builder = _builder
	// 读取模型参数
	_XBuilder.builder.modelFillArgs(false)
	return _XBuilder
}

func Model(model interface{}) *builder {
	orm := new(builder)
	orm.model = model
	orm.tableNames()
	orm.pk = "id"
	return orm
}

func (_builder *builder) modelFillArgs(fillSql bool) {

	p := reflect.ValueOf(_builder.model).Elem()
	_builder.pkValue = p.Field(0).Interface()
	for i := 1; i < p.NumField(); i++ {
		f := p.Field(i)
		field2 := f.Interface()
		_builder.args = append(_builder.args, field2)
		if fillSql {
			_builder.argsSql += "?,"
		}
	}
	if fillSql {
		_builder.argsSql = _builder.argsSql[0 : len(_builder.argsSql)-1]
	}
}

func (_builder *builder) Pk(pk string) *builder {
	_builder.pk = pk
	return _builder
}
