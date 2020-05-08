package db

import (
	"reflect"
)

// 处理字段
func (_builder *builder) scanFields() {

	_builder.fields = []string{}

	t1 := reflect.TypeOf(_builder.model).Elem()
	for i := 0; i < t1.NumField(); i++ {
		_builder.fields = append(_builder.fields, snakeCase(t1.Field(i).Name))
	}

}

func (_builder *builder) getFields() string {
	_builder.scanFields()
	v7 := ""
	for _, v := range _builder.fields {
		v7 += "`" + snakeCase(v) + "`,"
	}
	return v7[0 : len(v7)-1]
}

func (_builder *builder) getNotListFields(value2StrArgs value2StrArgs, args ...string) string {
	// 附带参数

	ig := make(map[string]int)

	for k, v := range args {
		ig[v] = k
	}

	_builder.scanFields()
	v7 := ""
	for _, v := range _builder.fields {
		if _, ok := ig[v]; ok {
			continue
		}
		v7 += value2StrArgs.Left + snakeCase(v) + value2StrArgs.Right + value2StrArgs.Link
	}
	return v7[0 : len(v7)-1]
}

func (_builder *builder) getInListFields(value2StrArgs value2StrArgs, arg string, args ...string) string {
	// 附带参数

	ig := make(map[string]int)
	ig[arg] = 0
	for k, v := range args {
		ig[v] = k
	}

	_builder.scanFields()
	v7 := ""
	for _, v := range _builder.fields {
		if _, ok := ig[v]; ok {
			v7 += value2StrArgs.Left + snakeCase(v) + value2StrArgs.Right + value2StrArgs.Link
		}
	}
	return v7[0 : len(v7)-1]
}
