package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"reflect"
)

var db *sql.DB

func ConfigDb(driverName, dsn string) {
	_db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Panic(err)
	}
	db = _db
}

type ORM struct {
	// 模型
	model interface{}
	// 查询条件
	where        interface{}
	fields       []string
	selectFields string
}

func Find(model interface{}) *ORM {
	orm := new(ORM)
	orm.model = model
	orm.getField()
	return orm
}

func (orm *ORM) getField() {

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
		v4 = append(v4, v.String())
		v7 += "`" + v.String() + "`,"
	}
	orm.selectFields = v7[0 : len(v7)-1]

	orm.fields = v4

}

func (orm *ORM) row(refs []interface{}) {

	v9 := reflect.ValueOf(orm.model)
	v10 := reflect.Indirect(v9)
	for k, _ := range orm.fields {
		v11 := v10.Field(k)
		unmarshalConverter(v11, refs[k])
	}
}

func (orm *ORM) One() {
	// 准备查询字段
	//todo 准备查询条件
	row := db.QueryRow("select " + orm.selectFields + " from users")
	refs := refs(orm.model)
	_ = row.Scan(refs...)
	orm.row(refs)
}

func New() *ORM {
	return new(ORM)
}
