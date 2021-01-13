package stream

import (
	"github.com/yanlong-li/hi-go-logger"
	"reflect"
)

// 从字节流中反射出对应的结构体并注入到指定方法中
func (ps *SocketPacketStream) Unmarshal(f interface{}) []reflect.Value {
	t := reflect.TypeOf(f)
	//构造一个存放函数实参 Value 值的数纽
	var in []reflect.Value
	// 取出所有需要注入的依赖参数
	for i := 0; i < t.NumIn()-1; i++ {
		// 获取顺序的 参数
		params := t.In(i)
		// 创建一个reflect.value类型的params需要的指针类型的数据
		elem := reflect.New(params).Elem()
		switch elem.Kind() {
		case reflect.Map:
			mapSize := ps.ReadInt64()
			var i int64 = 0
			keyType := elem.Type().Key()
			valType := elem.Type().Elem()

			e := reflect.MakeMap(elem.Type())

			for i = 0; i < mapSize; i++ {

				k := reflect.New(keyType).Elem()
				ps.UnmarshalConverter(k)
				v := reflect.New(valType).Elem()
				ps.UnmarshalConverter(v)
				e.SetMapIndex(k, v)
			}
			elem.Set(e)
		case reflect.Slice:
			fallthrough
		case reflect.Struct:
			for k := 0; k < elem.NumField(); k++ {
				field := elem.Field(k)
				value := ps.UnmarshalConverter(field)
				field.Set(value)
			}
		}
		in = append(in, elem)
	}

	return in

}

func (ps *SocketPacketStream) UnmarshalConverter(field reflect.Value) reflect.Value {

	switch field.Kind() {
	case reflect.String:
		field.SetString(ps.ReadString())
	case reflect.Bool:
		field.SetBool(ps.ReadBool())
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
		// 读取数量
		num := ps.ReadUInt16()
		newV := reflect.MakeSlice(field.Type(), 1, int(num)+1)
		newField := newV.Index(0)
		for i := 0; i < int(num); i++ {
			newV = reflect.Append(newV, ps.UnmarshalConverter(newField))
		}
		field.Set(newV.Slice(1, newV.Len()))
	case reflect.Struct:
		for k := 0; k < field.NumField(); k++ {
			field2 := field.Field(k)
			ps.UnmarshalConverter(field2)
		}
	case reflect.Map:
		mapSize := ps.ReadInt64()
		var i int64 = 0
		keyType := field.Type().Key()
		valType := field.Type().Elem()
		e := reflect.MakeMap(field.Type())
		for i = 0; i < mapSize; i++ {
			k := reflect.New(keyType).Elem()
			ps.UnmarshalConverter(k)
			v := reflect.New(valType).Elem()
			ps.UnmarshalConverter(v)
			e.SetMapIndex(k, v)
		}
		field.Set(e)
	default:
		logger.Fatal("未知类型", 0, field.Kind())
	}
	return field

}
