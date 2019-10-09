package stream

import "encoding/json"

//todo 将包结构体反射写入字节流中
func Marshal(packet interface{}) ([]byte, error) {

	return json.Marshal(packet)

}
