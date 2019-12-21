package db

import (
	"fmt"
	"log"
	"strings"
)

// In
func (orm *QueryBuilder) getWhere() string {
	var where string
	if len(orm.where) > 0 {
		where = "WHERE " + orm.where
	} else {
		where = ""
	}
	return where
}

// In
func (orm *QueryBuilder) WhereIn(field string, list []interface{}) *QueryBuilder {
	return orm
}

// In
func (orm *QueryBuilder) WhereNotIn(field string, list []interface{}) *QueryBuilder {
	return orm
}

// 区间
func (orm *QueryBuilder) WhereBetween(field string, list []interface{}) *QueryBuilder {
	return orm
}

// 非区间
func (orm *QueryBuilder) WhereNotBetween(field string, list []interface{}) *QueryBuilder {
	return orm
}

// or
func (orm *QueryBuilder) OrWhere(args ...interface{}) *QueryBuilder {
	orm.whereBuild("OR", args...)
	return orm
}

// and
func (orm *QueryBuilder) AndWhere(args ...interface{}) *QueryBuilder {
	orm.whereBuild("AND", args...)
	return orm
}

// 默认 and
func (orm *QueryBuilder) Where(args ...interface{}) *QueryBuilder {
	orm.AndWhere(args...)
	return orm
}

func (orm *QueryBuilder) whereBuild(op string, args ...interface{}) {

	_where := orm.whereArgsAuto(args...)

	if len(orm.where) == 0 {
		orm.where = _where
	} else {
		orm.where = "(" + orm.where + ") " + op + " (" + _where + ")"
	}

}

func (orm *QueryBuilder) whereArgsAuto(args ...interface{}) string {
	var _where string
	if len(args) == 1 {
		_where = orm.whereArgs1(args[0])
	} else if len(args) == 2 {
		_where = orm.whereArgs2(args[0], args[1])
	} else if len(args) == 3 {
		_where = orm.whereArgs3(args[1], args[0], args[2])
	} else if len(args) == 4 {
		_where = orm.whereArgs4(args[0], args[1], args[2], args[2])
	} else if len(args) > 4 {
		log.Panic("最多可接受 4 个参数，参数太多了。。。")
	}
	return _where

}

func (orm *QueryBuilder) whereArgs1(value1 interface{}) string {
	_where := ""
	if value, ok := value1.(string); ok {
		_where = value
	} else if value, ok := value1.([]interface{}); ok {
		if len(value) < 3 {
			log.Panic("使用切片时，参数不能小于3")
		}

		if op, ok := value[0].(string); !ok {
			log.Panic("比较类型必须是字符串类型", value1)
		} else {
			var ops = []string{"and", "or"}
			var ob = false
			for _, _op := range ops {
				if _op == strings.ToLower(op) {
					ob = true
					break
				}
			}
			if !ob {
				_where = orm.whereArgsAuto(value...)
			} else {
				__where := ""
				for _, _v := range value[1:] {

					if __v, ok := _v.([]interface{}); ok {
						if len(__v) <= 2 {
							__where += op + " (" + orm.whereArgsAuto(__v...) + ") "
						} else {
							__where += op + " (" + orm.whereArgs1(__v) + ") "
						}
					} else {
						__where += op + " " + orm.whereValue2Str(__v) + " "
					}
				}
				_where = __where[len(op):]
			}
		}

	} else if __v, ok := value1.(map[interface{}]interface{}); ok {
		_where = orm.whereValue2Str(__v)
	}
	return _where
}

// 针对字段 普通类型+map+切片
func (orm *QueryBuilder) whereArgs2(value1, value2 interface{}) string {
	_where := ""
	if _value, ok := value2.([]interface{}); ok {
		_where = orm.whereArgs3(value1, "IN", _value)
	} else {
		_where = orm.whereArgs3(value1, "=", value2)
	}
	return _where
}

// 针对字段
func (orm *QueryBuilder) whereArgs3(value1, value2, value3 interface{}) string {
	if _where, ok := value1.(string); !ok {
		log.Panic("字段必须是字符串类型", _where)
	}
	if _where, ok := value2.(string); !ok {
		log.Panic("比较类型必须是字符串类型", _where)
	}

	return fmt.Sprintf("`%s` %s %s", value1.(string), value2.(string), orm.whereValue2Str(value3))
}

// 针对字段
func (orm *QueryBuilder) whereArgs4(value1, value2, value3, value4 interface{}) string {
	if _where, ok := value1.(string); !ok {
		log.Panic("字段必须是字符串类型", _where)
	}
	if _where, ok := value2.(string); !ok {
		log.Panic("比较类型必须是字符串类型", _where)
	}

	return fmt.Sprintf("`%s` %s %s,%s", value1.(string), value2.(string), orm.whereValue2Str(value3), orm.whereValue2Str(value4))
}

// 检测where 值是否受支持
func whereValueSupportType(value interface{}) bool {
	_ts := false
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, []interface{}, map[interface{}]interface{}:
		_ts = true
	}
	return _ts
}

func (orm *QueryBuilder) whereValue2Str(inter interface{}) string {
	var _value = "?"
	switch value := inter.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string:
		orm.whereArgs = append(orm.whereArgs, value)
	case []interface{}:
		_tmp := ""
		for _, _v := range value {
			_tmp += orm.whereValue2Str(_v) + ","
		}
		_tmp = _tmp[0 : len(_tmp)-1]
		_value = "(" + _tmp + ")"
	case map[interface{}]interface{}:
		_tmp := ""
		for _k, _v := range value {
			_tmp += orm.whereArgs2(_k, _v) + " AND"
		}
		_tmp = _tmp[0 : len(_tmp)-4]
		_value = "(" + _tmp + ")"
	}
	return _value
}
