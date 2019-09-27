package stream

import (
	"log"
	"reflect"
)

//todo 将包结构体反射写入字节流中
func (ps *PacketStream) Marshal(packet interface{}) {

	elem := reflect.ValueOf(packet)
	for k := 0; k < elem.NumField(); k++ {
		field := elem.Field(k)
		switch field.Kind() {
		case reflect.String:
			ps.WriteString(field.String())
		case reflect.Uint8:
		case reflect.Uint16:
			ps.WriteUint16(uint16(field.Uint()))
		case reflect.Uint32:
			ps.WriteUint32(uint32(field.Uint()))
		case reflect.Uint64:
			ps.WriteUint64(uint64(field.Uint()))
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
		case reflect.Float32:
		case reflect.Float64:
		default:
			log.Fatal("不支持的类型")
		}
	}
}
