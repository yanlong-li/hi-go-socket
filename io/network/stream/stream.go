package stream

import (
	"reflect"
)

type Interface interface {
	Marshal(interface{})
	Unmarshal(interface{}) []reflect.Value
	GetData() []byte
	ToData() []byte
	GetOpCode() uint32
	GetLen() uint16
}

type BaseStream struct {
	data   []byte //数据储存体
	Index  uint16 //当前指针
	len    uint16 //数据长度--来自消息告知
	OpCode uint32 //操作码
}

func (bs *BaseStream) GetData() []byte {
	return bs.data
}

func (bs *BaseStream) SetData(Data []byte) {
	bs.data = Data
}

func (bs *BaseStream) GetOpCode() uint32 {
	return bs.OpCode
}
func (bs *BaseStream) SetOpCode(OpCode uint32) {
	bs.OpCode = OpCode
}

func (bs *BaseStream) GetLen() uint16 {
	return bs.len
}
func (bs *BaseStream) SetLen(Len uint16) {
	bs.len = Len
}

func (bs *BaseStream) ToData() []byte {
	//创建固定长度的数组节省内存
	var data []byte
	data = append(data, Uint16ToBytes(bs.GetLen()+4)...)
	data = append(data, Uint32ToBytes(bs.OpCode)...)
	data = append(data, bs.data...)
	return data
}
