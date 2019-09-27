package stream

import (
	"fmt"
	"log"
	"reflect"
)

// 从字节流中反射出对应的结构体并注入到指定方法中
func (ps *PacketStream) Unmarshal(f interface{}) []reflect.Value {
	t := reflect.TypeOf(f)
	//构造一个存放函数实参 Value 值的数纽
	in := make([]reflect.Value, t.NumIn())
	// 取出所有需要注入的依赖参数
	for i := 0; i < t.NumIn()-1; i++ {
		// 获取顺序的 参数
		params := t.In(i)
		// 创建一个reflect.value类型的params需要的指针类型的数据
		elem := reflect.New(params).Elem()
		for k := 0; k < elem.NumField(); k++ {
			field := elem.Field(k)
			switch field.Kind() {
			case reflect.String:
				field.SetString(ps.ReadString())
			case reflect.Uint8:
				field.SetUint(uint64(ps.ReadUInt8()))
			case reflect.Uint16:
				field.SetUint(uint64(ps.ReadUInt16()))
			case reflect.Uint32:
				field.SetUint(uint64(ps.ReadUInt32()))
			case reflect.Uint64:
				field.SetUint(ps.ReadUInt64())
			case reflect.Int8:
				field.SetInt(int64(ps.ReadInt8()))
			case reflect.Int16:
				field.SetInt(int64(ps.ReadInt16()))
			case reflect.Int32:
				field.SetInt(int64(ps.ReadInt32()))
			case reflect.Int64:
				field.SetInt(ps.ReadInt64())
			case reflect.Float32:
				field.SetFloat(float64(ps.ReadFloat32()))
			case reflect.Float64:
				field.SetFloat(ps.ReadFloat64())
			case reflect.Slice:

				fmt.Println(field.Type())
				num := ps.ReadUInt16()
				field.SetCap(int(num))
				field.SetLen(int(num))
				fmt.Println(field.Len())
				for i := 0; uint16(i) < num; i++ {
					e := field.Index(i)
					e.SetString(ps.ReadString())
				}
			default:
				log.Fatal("不支持读取的类型")
			}
		}
		in[i] = elem
	}
	return in
}
