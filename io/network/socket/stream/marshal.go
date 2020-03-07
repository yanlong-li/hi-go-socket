package stream

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"log"
	"reflect"
)

//将包结构体反射写入字节流中
func (ps *SocketPacketStream) Marshal(PacketModel interface{}) {
	ps.SetData([]byte{})

	ps.SetOpCode(packet.OpCode(PacketModel))

	elem := reflect.ValueOf(PacketModel)
	for k := 0; k < elem.NumField(); k++ {
		field := elem.Field(k)
		ps.MarshalConverter(field)
	}
	ps.SetLen(uint16(len(ps.GetData())))
}
func (ps *SocketPacketStream) MarshalConverter(field reflect.Value) {
	switch field.Kind() {
	case reflect.String:
		ps.WriteString(field.String())
	case reflect.Int:
		ps.WriteInt64(field.Int())
	case reflect.Uint8:
		ps.WriteUint8(uint8(field.Uint()))
	case reflect.Uint16:
		ps.WriteUint16(uint16(field.Uint()))
	case reflect.Uint32:
		ps.WriteUint32(uint32(field.Uint()))
	case reflect.Uint64:
		ps.WriteUint64(field.Uint())
	case reflect.Int8:
		ps.WriteInt8(int8(field.Int()))
	case reflect.Int16:
		ps.WriteInt16(int16(field.Int()))
	case reflect.Int32:
		ps.WriteInt32(int32(field.Int()))
	case reflect.Int64:
		ps.WriteInt64(int64(field.Int()))
	case reflect.Float32:
		ps.WriteFloat32(float32(field.Float()))
	case reflect.Float64:
		ps.WriteFloat64(field.Float())
	case reflect.Slice:
		ps.WriteUint16(uint16(field.Len()))
		for i := 0; i < field.Len(); i++ {
			elm := field.Index(i)
			ps.MarshalConverter(elm)
		}
	case reflect.Struct:
		for k := 0; k < field.NumField(); k++ {
			field := field.Field(k)
			ps.MarshalConverter(field)
		}
	default:
		log.Fatal("不支持写入的类型", field.Kind())
	}
}
