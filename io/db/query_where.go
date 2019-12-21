package db

import (
	"fmt"
	"log"
	"strconv"
)

// In
func (orm *QueryBuilder) getWhere() string {
	var where string
	if len(orm.where) > 0 {
		where = "WHERE" + orm.where
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
	var _where string
	if len(args) == 1 {
		_where = whereArgs1(args[0])
	} else if len(args) == 2 {
		_where = whereArgs2(args[0], args[1])
	} else if len(args) == 3 {
		_where = whereArgs3(args[0], args[1], args[2])
	} else if len(args) > 3 {
		log.Panic("最多可接受 3 个参数，参数太多了。。。")
	}

	if len(orm.where) > 0 {
		orm.where = _where
	} else {
		orm.where = "(" + orm.where + ") " + op + " (" + _where + ")"
	}

}

func whereArgs1(value1 interface{}) string {
	_where := ""
	if value, ok := value1.(string); ok {
		_where = value
	} else if value, ok := value1.([]interface{}); ok {
		if len(value) < 3 {
			log.Panic("使用切片时，参数不能小于3")

		}
		__where := ""
		for _, _v := range value[1:] {

			if __v, ok := _v.([]interface{}); ok {
				if len(__v) == 2 {
					__where += " AND " + whereArgs2(__v[0], __v[1])
				} else if len(__v) == 3 {
					__where += " AND " + whereArgs3(__v[1], __v[0], __v[2])
				}
			} else if __v, ok := _v.(string); ok {
				__where += " AND " + __v
			} else {
				log.Panic("不支持的类型")
			}
		}
		_where = __where[4:]

	}
	return _where
}
func whereArgs2(value1, value2 interface{}) string {
	_where := ""
	if _value, ok := value2.([]interface{}); ok {
		_where = whereArgs3(value1, "IN", _value)
	} else {
		_where = whereArgs3(value1, "=", value2)
	}
	return _where
}
func whereArgs3(value1, value2, value3 interface{}) string {
	if _where, ok := value1.(string); !ok {
		log.Panic("字段必须是字符串类型", _where)
	}
	if _where, ok := value2.(string); !ok {
		log.Panic("比较类型必须是字符串类型", _where)
	}
	return fmt.Sprintf("`%s` %s %s", value1.(string), value2.(string), whereValue2Str(value3))
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

func whereValue2Str(inter interface{}) string {
	var _value = ""
	switch value := inter.(type) {
	case int:
		_value = strconv.Itoa(value)
	case int8:
		_value = strconv.Itoa(int(value))
	case int16:
		_value = strconv.Itoa(int(value))
	case int32:
		_value = strconv.Itoa(int(value))
	case int64:
		_value = strconv.Itoa(int(value))
	case uint:
		_value = strconv.Itoa(int(value))
	case uint8:
		_value = strconv.Itoa(int(value))
	case uint32:
		_value = strconv.Itoa(int(value))
	case uint16:
		_value = strconv.Itoa(int(value))
	case uint64:
		_value = strconv.Itoa(int(value))
	case float32:
		_value = strconv.FormatFloat(float64(value), 'f', 6, 64)
	case float64:
		_value = strconv.FormatFloat(value, 'f', 6, 64)
	case string:
		_value = value
	case []interface{}:
		_tmp := ""
		for _, _v := range value {
			_tmp += whereValue2Str(_v) + ","
		}
		_tmp = _tmp[0 : len(_tmp)-1]
		_value = "(" + _tmp + ")"
	case map[interface{}]interface{}:
		_tmp := ""
		for _k, _v := range _value {
			_tmp += whereArgs3(whereValue2Str(_k), "=", whereValue2Str(_v)) + " AND"
		}
		_tmp = _tmp[0 : len(_tmp)-4]
		_value = "(" + _tmp + ")"
	}
	return _value
}
