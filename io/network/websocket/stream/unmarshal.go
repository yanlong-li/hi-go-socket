package stream

import (
	"encoding/json"
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"log"
	"reflect"
)

// 从字节流中反射出对应的结构体并注入到指定方法中
func (wsps *WebSocketPacketStream) Unmarshal(f interface{}) []reflect.Value {
	t := reflect.TypeOf(f)
	//构造一个存放函数实参 Value 值的数纽
	var in []reflect.Value

	var p interface{}
	err := json.Unmarshal(wsps.GetData(), &p)

	if err != nil {
		logger.Debug("无法解析的数据", 0, string(wsps.GetData()))
	}

	//keys:=p.keys

	value := reflect.ValueOf(p)
	keys := value.MapKeys()

	for i := 0; i < t.NumIn()-1; i++ {
		packet := t.In(i)
		// 创建一个reflect.value类型的params需要的指针类型的数据
		elem := reflect.New(packet).Elem()

		for k := 0; k < len(keys); k++ {
			field := elem.FieldByName(keys[k].String())
			if !field.IsValid() {
				log.Println("未知字段：", keys[k].String())
				break
			}
			wsps.UnmarshalConverter(field, value.MapIndex(keys[k]).Elem())
		}
		in = append(in, elem)
	}
	return in
}

func (wsps *WebSocketPacketStream) UnmarshalConverter(field, field2 reflect.Value) reflect.Value {

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
		for i := 0; i < num; i++ {
			newV = reflect.Append(newV, wsps.UnmarshalConverter(field, newV.Index(0)))
		}
		field.Set(newV.Slice(1, newV.Len()))
	default:
		log.Println("忽略一个未知类型:", field.Kind(), field.String())
	}
	return field

}
