package db

import (
	"fmt"
	"log"
	"strings"
)

// In
func (query *queryBuilder) getWhere() string {
	var where string
	if len(query.argsSql) > 0 {
		where = "WHERE " + query.argsSql
	} else {
		where = ""
	}
	return where
}

// In
func (query *queryBuilder) WhereIn(field string, list []interface{}) *queryBuilder {
	return query
}

// In
func (query *queryBuilder) WhereNotIn(field string, list []interface{}) *queryBuilder {
	return query
}

// 区间
func (query *queryBuilder) WhereBetween(field string, list []interface{}) *queryBuilder {
	return query
}

// 非区间
func (query *queryBuilder) WhereNotBetween(field string, list []interface{}) *queryBuilder {
	return query
}

// or
func (query *queryBuilder) OrWhere(args ...interface{}) *queryBuilder {
	query.whereBuild("OR", args...)
	return query
}

// and
func (query *queryBuilder) AndWhere(args ...interface{}) *queryBuilder {
	query.whereBuild("AND", args...)
	return query
}

// 默认 and
func (query *queryBuilder) Where(args ...interface{}) *queryBuilder {
	query.AndWhere(args...)
	return query
}

func (query *queryBuilder) whereBuild(op string, args ...interface{}) {

	_where := query.whereArgsAuto(args...)

	if len(query.argsSql) == 0 {
		query.argsSql = _where
	} else {
		query.argsSql = "(" + query.argsSql + ") " + op + " (" + _where + ")"
	}

}

func (query *queryBuilder) whereArgsAuto(args ...interface{}) string {
	var _where string
	if len(args) == 1 {
		_where = query.whereArgs1(args[0])
	} else if len(args) == 2 {
		_where = query.whereArgs2(args[0], args[1])
	} else if len(args) == 3 {
		_where = query.whereArgs3(args[1], args[0], args[2])
	} else if len(args) == 4 {
		_where = query.whereArgs4(args[0], args[1], args[2], args[2])
	} else if len(args) > 4 {
		log.Panic("最多可接受 4 个参数，参数太多了。。。")
	}
	return _where

}

func (query *queryBuilder) whereArgs1(value1 interface{}) string {
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
				_where = query.whereArgsAuto(value...)
			} else {
				__where := ""
				for _, _v := range value[1:] {

					if __v, ok := _v.([]interface{}); ok {
						if len(__v) <= 2 {
							__where += op + " (" + query.whereArgsAuto(__v...) + ") "
						} else {
							__where += op + " (" + query.whereArgs1(__v) + ") "
						}
					} else {
						__where += op + " " + query.whereValue2Str(__v) + " "
					}
				}
				_where = __where[len(op):]
			}
		}

	} else if __v, ok := value1.(map[interface{}]interface{}); ok {
		_where = query.whereValue2Str(__v)
	}
	return _where
}

// 针对字段 普通类型+map+切片
func (query *queryBuilder) whereArgs2(value1, value2 interface{}) string {
	_where := ""
	if _value, ok := value2.([]interface{}); ok {
		_where = query.whereArgs3(value1, "IN", _value)
	} else {
		_where = query.whereArgs3(value1, "=", value2)
	}
	return _where
}

// 针对字段
func (query *queryBuilder) whereArgs3(value1, value2, value3 interface{}) string {
	if _where, ok := value1.(string); !ok {
		log.Panic("字段必须是字符串类型", _where)
	}
	if _where, ok := value2.(string); !ok {
		log.Panic("比较类型必须是字符串类型", _where)
	}

	return fmt.Sprintf("`%s` %s %s", value1.(string), value2.(string), query.whereValue2Str(value3))
}

// 针对字段
func (query *queryBuilder) whereArgs4(value1, value2, value3, value4 interface{}) string {
	if _where, ok := value1.(string); !ok {
		log.Panic("字段必须是字符串类型", _where)
	}
	if _where, ok := value2.(string); !ok {
		log.Panic("比较类型必须是字符串类型", _where)
	}

	return fmt.Sprintf("`%s` %s %s,%s", value1.(string), value2.(string), query.whereValue2Str(value3), query.whereValue2Str(value4))
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

func (query *queryBuilder) whereValue2Str(inter interface{}) string {
	var _value = "?"
	switch value := inter.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string:
		query.args = append(query.args, value)
	case []interface{}:
		_tmp := ""
		for _, _v := range value {
			_tmp += query.whereValue2Str(_v) + ","
		}
		_tmp = _tmp[0 : len(_tmp)-1]
		_value = "(" + _tmp + ")"
	case map[interface{}]interface{}:
		_tmp := ""
		for _k, _v := range value {
			_tmp += query.whereArgs2(_k, _v) + " AND"
		}
		_tmp = _tmp[0 : len(_tmp)-4]
		_value = "(" + _tmp + ")"
	}
	return _value
}
