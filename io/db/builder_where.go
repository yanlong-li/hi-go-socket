package db

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"log"
	"strings"
)

// In
func (_b *builder) getWhere() string {
	var where string
	if len(_b.argsSql) > 0 {
		where = "WHERE " + _b.argsSql
	} else {
		where = ""
	}
	return where
}

// In
func (_b *builder) whereIn(field string, list []interface{}) *builder {
	_b.andWhere("IN", field, list)
	return _b
}

// Not In
func (_b *builder) whereNotIn(field string, list []interface{}) *builder {
	_b.andWhere("NOT IN", field, list)
	return _b
}

// 区间
func (_b *builder) whereBetween(field string, value1, value2 interface{}) *builder {
	_b.andWhere(field, "BETWEEN ", value1, value2)
	return _b
}

// 非区间
func (_b *builder) whereNotBetween(field string, value1, value2 interface{}) *builder {
	_b.andWhere(field, "NOT BETWEEN ", value1, value2)
	return _b
}

// or
func (_b *builder) orWhere(args ...interface{}) *builder {
	_b.whereBuild("OR", args...)
	return _b
}

// and
func (_b *builder) andWhere(args ...interface{}) *builder {
	_b.whereBuild("AND", args...)
	return _b
}

// 默认 and
func (_b *builder) where(args ...interface{}) *builder {
	_b.andWhere(args...)
	return _b
}

// 构建 where 条件
func (_b *builder) whereBuild(link string, args ...interface{}) {

	_where := _b.whereArgsAuto(args...)

	if len(_b.argsSql) == 0 {
		_b.argsSql = _where
	} else {
		_b.argsSql = "(" + _b.argsSql + ") " + link + " (" + _where + ")"
	}

}

// 自动多参数拆分构建
func (_b *builder) whereArgsAuto(args ...interface{}) string {
	var _where string
	if len(args) == 1 {
		_where = _b.whereArgs1(args[0])
	} else if len(args) == 2 {
		_where = _b.whereArgs2(args[0].(string), args[1])
	} else if len(args) == 3 {
		_where = _b.whereArgs3(args[1].(string), args[0].(string), args[2])
	} else if len(args) == 4 {
		_where = _b.whereArgs4(args[0].(string), args[1].(string), args[2], args[3])
	} else if len(args) > 4 {
		log.Panic("最多可接受 4 个参数，参数太多了。。。")
	}
	return _where

}

func (_b *builder) whereArgs1(value1 interface{}) string {
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
				_where = _b.whereArgsAuto(value...)
			} else {
				__where := ""
				for _, _v := range value[1:] {

					if __v, ok := _v.([]interface{}); ok {
						if len(__v) <= 2 {
							__where += op + " (" + _b.whereArgsAuto(__v...) + ") "
						} else {
							__where += op + " (" + _b.whereArgs1(__v) + ") "
						}
					} else {
						__where += op + " " + _b.whereValue2Str(__v) + " "
					}
				}
				_where = __where[len(op):]
			}
		}

	} else if __v, ok := value1.(map[interface{}]interface{}); ok {
		_where = _b.whereValue2Str(__v)
	}
	return _where
}

// 针对字段 普通类型+map+切片
func (_b *builder) whereArgs2(field string, values interface{}) string {
	_where := ""
	if _value, ok := values.([]interface{}); ok {
		_where = _b.whereArgs3(field, "IN", _value)
	} else {
		_where = _b.whereArgs3(field, "=", values)
	}
	return _where
}

// 针对字段
func (_b *builder) whereArgs3(field, symbol string, values interface{}) string {
	return fmt.Sprintf("`%s` %s %s", field, symbol, _b.whereValue2Str(values))
}

// 针对字段
func (_b *builder) whereArgs4(field, symbol string, value1, value2 interface{}) string {
	return fmt.Sprintf("`%s` %s %s%s%s", field, symbol, _b.whereValue2Str(value1), getSymbolLink(symbol), _b.whereValue2Str(value2))
}
func getSymbolLink(symbol string) (link string) {
	switch symbol {
	case "BETWEEN":
		link = " AND "
	default:
		link = ", "
	}
	return
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

type value2StrArgs struct {
	Link  string
	Left  string
	Right string
}

func (_b *builder) whereValue2Str(inter interface{}, _value2StrArgs ...value2StrArgs) string {
	var _value = "?"

	// 附带参数
	value2StrArgs := value2StrArgs{Link: ",", Left: "(", Right: ")"}
	// 只允许附带一个
	if len(_value2StrArgs) > 1 {
		panic("whereValue2Str _value2StrArgs max length 1")
	} else if len(_value2StrArgs) == 1 {
		// 使用非默认值
		__value2StrArgs := _value2StrArgs[0]
		// 至少也得是个空格
		//参数连接符
		if __value2StrArgs.Link != "" {
			value2StrArgs.Link = __value2StrArgs.Link
		}
		//参数左连接符
		if __value2StrArgs.Left != "" {
			value2StrArgs.Left = __value2StrArgs.Left
		}
		//参数右连接符
		if __value2StrArgs.Right != "" {
			value2StrArgs.Right = __value2StrArgs.Right
		}
	}

	switch value := inter.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string:
		_b.args = append(_b.args, value)
	case []interface{}:
		_tmp := ""
		for _, _v := range value {
			_tmp += _b.whereValue2Str(_v, _value2StrArgs...) + value2StrArgs.Link
		}
		_tmp = _tmp[0 : len(_tmp)-1]
		_value = value2StrArgs.Left + _tmp + value2StrArgs.Right
	case map[interface{}]interface{}:
		_tmp := ""
		for _k, _v := range value {
			_tmp += _b.whereArgs2(_k.(string), _v) + " AND"
		}
		_tmp = _tmp[0 : len(_tmp)-4]
		_value = value2StrArgs.Left + _tmp + value2StrArgs.Right
	default:
		logger.Fatal("不支持的数据类型", 0, value)
	}
	return _value
}
