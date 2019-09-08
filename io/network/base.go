package network

const (
	Tcp = "tcp"
	Udp = "udp" // 暂时不用
)

var packets map[uint16]interface{} = make(map[uint16]interface{})

// 包服务注册应由引用源程序处理并注册到当前包的packets之下

func Register(op uint16, packet interface{}) {
	packets[op] = packet
}
