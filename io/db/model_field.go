package db

import (
	"reflect"
)

// 处理字段
func (_b *builder) scanFields() {

	_b.fields = []string{}

	t1 := reflect.TypeOf(_b.model).Elem()
	for i := 0; i < t1.NumField(); i++ {
		_b.fields = append(_b.fields, snakeCase(t1.Field(i).Name))
	}

}

func (_b *builder) getFields() string {
	_b.scanFields()
	v7 := ""
	for _, v := range _b.fields {
		v7 += "`" + snakeCase(v) + "`,"
	}
	return v7[0 : len(v7)-1]
}

func (_b *builder) getNotListFields(args ...string) string {

	ig := make(map[string]int)

	for k, v := range args {
		ig[v] = k
	}

	_b.scanFields()
	v7 := ""
	for _, v := range _b.fields {
		if _, ok := ig[v]; ok {
			continue
		}
		v7 += "`" + snakeCase(v) + "`,"
	}
	return v7[0 : len(v7)-1]
}
