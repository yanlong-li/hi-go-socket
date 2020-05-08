package db

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"log"
	"strings"
)

const (
	WhereAnd        = "AND"
	WhereOr         = "OR"
	WhereIN         = "IN"
	WhereNotIN      = "NOT IN"
	WhereBetween    = "BETWEEN"
	WhereNotBetween = "NOT BETWEEN"
)

// In
func (_builder *builder) getWhere() string {
	var where string
	if len(_builder.argsSql) > 0 {
		where = "WHERE " + _builder.argsSql
	} else {
		where = ""
	}
	return where
}

// In
func (_builder *builder) whereIn(field string, list []interface{}) *builder {
	_builder.andWhere(WhereIN, field, list)
	return _builder
}

// Not In
func (_builder *builder) whereNotIn(field string, list []interface{}) *builder {
	_builder.andWhere(WhereNotIN, field, list)
	return _builder
}

// 区间
func (_builder *builder) whereBetween(field string, value1, value2 interface{}) *builder {
	_builder.andWhere(field, WhereBetween, value1, value2)
	return _builder
}

// 非区间
func (_builder *builder) whereNotBetween(field string, value1, value2 interface{}) *builder {
	_builder.andWhere(field, WhereNotBetween, value1, value2)
	return _builder
}

// or
func (_builder *builder) orWhere(args ...interface{}) *builder {
	_builder.whereBuild(WhereOr, args...)
	return _builder
}

// and
func (_builder *builder) andWhere(args ...interface{}) *builder {
	_builder.whereBuild(WhereAnd, args...)
	return _builder
}

// 默认 and
func (_builder *builder) where(args ...interface{}) *builder {
	_builder.andWhere(args...)
	return _builder
}

// 构建 where 条件
func (_builder *builder) whereBuild(link string, args ...interface{}) {

	_where := _builder.whereArgsAuto(args...)

	if len(_builder.argsSql) == 0 {
		_builder.argsSql = _where
	} else {
		_builder.argsSql = "(" + _builder.argsSql + ") " + link + " (" + _where + ")"
	}

}

// 自动多参数拆分构建
func (_builder *builder) whereArgsAuto(args ...interface{}) string {
	var _where string
	if len(args) == 1 {
		_where = _builder.whereArgs1(args[0])
	} else if len(args) == 2 {
		_where = _builder.whereArgs2(args[0].(string), args[1])
	} else if len(args) == 3 {
		_where = _builder.whereArgs3(args[1].(string), args[0].(string), args[2])
	} else if len(args) == 4 {
		_where = _builder.whereArgs4(args[0].(string), args[1].(string), args[2], args[3])
	} else if len(args) > 4 {
		log.Panic("最多可接受 4 个参数，参数太多了。。。")
	}
	return _where

}

func (_builder *builder) whereArgs1(value1 interface{}) string {
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
				_where = _builder.whereArgsAuto(value...)
			} else {
				__where := ""
				for _, _v := range value[1:] {

					if __v, ok := _v.([]interface{}); ok {
						if len(__v) <= 2 {
							__where += op + " (" + _builder.whereArgsAuto(__v...) + ") "
						} else {
							__where += op + " (" + _builder.whereArgs1(__v) + ") "
						}
					} else {
						__where += op + " " + _builder.whereValue2Str(__v) + " "
					}
				}
				_where = __where[len(op):]
			}
		}

	} else if __v, ok := value1.(map[interface{}]interface{}); ok {
		_where = _builder.whereValue2Str(__v)
	} else {
		logger.Fatal("不支持的数据类型", 0, value1)
	}
	return _where
}

// 针对字段 普通类型+map+切片
func (_builder *builder) whereArgs2(field string, values interface{}) string {
	_where := ""
	if _value, ok := values.([]interface{}); ok {
		_where = _builder.whereArgs3(field, WhereIN, _value)
	} else {
		_where = _builder.whereArgs3(field, "=", values)
	}
	return _where
}

// 针对字段
func (_builder *builder) whereArgs3(field, symbol string, values interface{}) string {
	return fmt.Sprintf("`%s` %s %s", field, symbol, _builder.whereValue2Str(values))
}

// 针对字段
func (_builder *builder) whereArgs4(field, symbol string, value1, value2 interface{}) string {
	return fmt.Sprintf("`%s` %s %s%s%s", field, symbol, _builder.whereValue2Str(value1), getSymbolLink(symbol), _builder.whereValue2Str(value2))
}
func getSymbolLink(symbol string) (link string) {
	switch symbol {
	case WhereNotBetween:
		link = " " + WhereAnd + " "
	case WhereBetween:
		link = " " + WhereAnd + " "
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

func (_builder *builder) whereValue2Str(inter interface{}, _value2StrArgs ...value2StrArgs) string {
	var _value = "?"

	// 附带参数
	value2StrArgs := value2StrArgs{Link: ",", Left: "(", Right: ")"}
	// 只允许附带一个
	if len(_value2StrArgs) > 1 {
		logger.Trace("whereValue2Str _value2StrArgs max length 1", 0)
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
		_builder.args = append(_builder.args, value)
	case []interface{}:
		_tmp := ""
		for _, _v := range value {
			_tmp += _builder.whereValue2Str(_v, _value2StrArgs...) + value2StrArgs.Link
		}
		_tmp = _tmp[0 : len(_tmp)-1]
		_value = value2StrArgs.Left + _tmp + value2StrArgs.Right
	case map[interface{}]interface{}:
		_tmp := ""
		for _k, _v := range value {
			_tmp += _builder.whereArgs2(_k.(string), _v) + " " + WhereAnd
		}
		_tmp = _tmp[0 : len(_tmp)-len(" "+WhereAnd)]
		_value = value2StrArgs.Left + _tmp + value2StrArgs.Right
	default:
		logger.Fatal("不支持的数据类型", 0, value)
	}
	return _value
}
