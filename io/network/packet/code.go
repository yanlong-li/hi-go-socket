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
	//保留范围 0-ReservedCode  自定义 Code 不能小于 ReservedCode
	ReservedCode = 5999
)

// 基础数据定义
const (
	// 标识码占位长度 4 字节 uint32
	OpCodeLen = 4
	// 数据包长度标识符占位 2 字节 uint16
	BufLenLen = 2
)
