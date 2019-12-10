package stream

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// 从字节流中反射出对应的结构体并注入到指定方法中
func Unmarshal(f interface{}, data []byte) []reflect.Value {
	t := reflect.TypeOf(f)
	//构造一个存放函数实参 Value 值的数纽
	in := make([]reflect.Value, t.NumIn())

	var p interface{}
	err := json.Unmarshal(data, &p)

	if err != nil {
		fmt.Println("无法解析的数据", string(data))
	}

	//keys:=p.keys

	value := reflect.ValueOf(p)
	keys := value.MapKeys()

	for i := 0; i < t.NumIn()-1; i++ {
		packet := t.In(i)
		// 创建一个reflect.value类型的params需要的指针类型的数据
		elem := reflect.New(packet).Elem()

		for k := 0; k < elem.NumField() && k < len(keys); k++ {
			UnmarshalConverter(elem.FieldByName(keys[k].String()), value.MapIndex(keys[k]).Elem())
		}
		in[i] = elem
	}
	return in
}

func UnmarshalConverter(field, field2 reflect.Value) reflect.Value {

	switch field.Kind() {
	case reflect.String:
		field.SetString(field2.String())
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		field.SetUint(uint64(field2.Float()))
		//fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		//field.SetInt(field2.Int())
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		field.SetFloat(field2.Float())
	case reflect.Slice:
		field2.Len()
		// 读取数量
		num := field2.Len()
		newV := reflect.MakeSlice(field.Type(), 1, int(num))
		for i := 0; i < int(num); i++ {
			newV = reflect.Append(newV, UnmarshalConverter(field, newV.Index(0)))
		}
		field.Set(newV.Slice(1, newV.Len()))
	default:
		log.Fatal("未知类型", field.Kind())
	}
	return field

}
