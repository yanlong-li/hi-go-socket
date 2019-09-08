package handle

// 注册包处理动作

var Packets = make(map[uint16]interface{})

// 包服务注册应由引用源程序处理并注册到当前包的packets之下
func Register(op uint16, packet interface{}) {
	Packets[op] = packet
}

func Handle(op uint16) interface{} {

	if v, ok := Packets[op]; ok {
		return v
	}

	return nil
}
