package packet

// 这里是系统内置的一些动作注册包，不得用于数据交互，否则解包时数据格式无法匹配
// 0 ~ 5999
const (
	// 链接成功
	Connection uint32 = iota
	//断开链接,包括已断开和准备断开
	Disconnection
	// 收到数据包处理之前
	BeforeRecv
	// 发送数据包之前
	BeforeSending

	MaxCode = 5999
)
