package db

import (
	"reflect"
)

// 处理字段
func (query *queryBuilder) scanFields() {

	query.fields = []string{}

	t1 := reflect.TypeOf(query.model).Elem()
	for i := 0; i < t1.NumField(); i++ {
		query.fields = append(query.fields, snakeCase(t1.Field(i).Name))
	}

}

func (query *queryBuilder) getFields() string {
	query.scanFields()
	v7 := ""
	for _, v := range query.fields {
		v7 += "`" + snakeCase(v) + "`,"
	}
	return v7[0 : len(v7)-1]
}

func (query *queryBuilder) getNotListFields(args ...string) string {

	ig := make(map[string]int)

	for k, v := range args {
		ig[v] = k
	}

	query.scanFields()
	v7 := ""
	for _, v := range query.fields {
		if _, ok := ig[v]; ok {
			continue
		}
		v7 += "`" + snakeCase(v) + "`,"
	}
	return v7[0 : len(v7)-1]
}
