package db

import (
	"log"
	"reflect"
)

func refs(model interface{}) []interface{} {
	v9 := reflect.ValueOf(model)
	v10 := reflect.Indirect(v9)
	var refs = make([]interface{}, v10.NumField())
	for i := 0; i < v10.NumField(); i++ {
		field := v10.Field(i)
		switch field.Kind() {
		case reflect.String:
			var ref string
			refs[i] = &ref
		case reflect.Uint8:
			var ref uint8
			refs[i] = &ref
		case reflect.Uint16:
			var ref uint16
			refs[i] = &ref
		case reflect.Uint32:
			var ref uint32
			refs[i] = &ref
		case reflect.Uint64:
			var ref uint64
			refs[i] = &ref
		case reflect.Int8:
			var ref int8
			refs[i] = &ref
		case reflect.Int16:
			var ref int16
			refs[i] = &ref
		case reflect.Int32:
			var ref int32
			refs[i] = &ref
		case reflect.Int64:
			var ref int64
			refs[i] = &ref
		case reflect.Float32:
			var ref float32
			refs[i] = &ref
		case reflect.Float64:
			var ref float64
			refs[i] = &ref
		default:
			log.Fatal("未知类型", field.Kind())
		}
	}
	return refs
}

func unmarshalConverter(field reflect.Value, ref interface{}) reflect.Value {

	field2 := reflect.Indirect(reflect.ValueOf(ref)).Interface()
	switch field.Kind() {
	case reflect.String:
		field.SetString(field2.(string))
	case reflect.Uint8:
		field.SetUint(uint64(field2.(uint8)))
	case reflect.Uint16:
		field.SetUint(uint64(field2.(uint16)))
	case reflect.Uint32:
		field.SetUint(uint64(field2.(uint32)))
	case reflect.Uint64:
		field.SetUint(field2.(uint64))
	case reflect.Int8:
		field.SetInt(int64(field2.(int8)))
	case reflect.Int16:
		field.SetInt(int64(field2.(int16)))
	case reflect.Int32:
		field.SetInt(int64(field2.(int32)))
	case reflect.Int64:
		field.SetInt(field2.(int64))
	case reflect.Float32:
		field.SetFloat(float64(field2.(float32)))
	case reflect.Float64:
		field.SetFloat(field2.(float64))
	default:
		log.Fatal("未知类型", field.Kind())
	}
	return field

}
