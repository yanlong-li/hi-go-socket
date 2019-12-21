package db

import (
	"encoding/json"
	"reflect"
)

// 处理字段
func (orm *QueryBuilder) initField() {

	// 将模型解析成字节组
	v2, _ := json.Marshal(orm.model)
	// 定义恢复变量
	var v3 interface{}
	// 恢复到恢复变量
	_ = json.Unmarshal(v2, &v3)
	// 解析map数据
	v5 := reflect.ValueOf(v3)
	// 获取map中的keys
	v6 := v5.MapKeys()

	var v4 []string
	var v7 string
	for _, v := range v6 {
		v4 = append(v4, snakeCase(v.String()))
		v7 += "`" + snakeCase(v.String()) + "`,"
	}
	orm.selectFields = v7[0 : len(v7)-1]

	orm.fields = v4

}
