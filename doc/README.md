# 使用教程

> 包含客户端和服务端的使用教程

## 服务端

[demo](./demo/server/server.go)

### 数据包

首先定义数据包，数据包是用来和客户端通信的数据模型 采用go语言中的结构体作为一个数据包, 支持类型有 bool、int/uint8~64、Float32/64、string、slice、struct(嵌套)

暂不支持map类型，后续会支持

不支持指针类型

```go
package main

type demo struct {
	Id   uint32
	Name string
}

```

### 注册数据包

数据包模型定义后需要注册，才能被正常使用 数据包注册需要两个参数，1 操作码，2 数据包

操作码：uint32 唯一，用于识别数据包，0-5999 系统保留，需要从 6000 开始自定义 数据包：上一步中定义的数据包，需要实例化后的空对象 `demo{}`

* 请注意操作码唯一性，否则后定义将覆盖前定义

```go
package main

import "github.com/yanlong-li/hi-go-socket/packet"

type demo struct {
	Id   uint32
	Name string
}

func init() {
	packet.Register(6000, demo{})
}

```

### 路由、动作

又称动作 action ：收到数据包后执行的动作,接收两个参数

第一个是要接收的包类型，如：demo 结构体,

第二个是固定的 [connect.Connector](../connect/connector.go) 类型

```go
package main

import (
	"fmt"
	"github.com/yanlong-li/hi-go-socket/connect"
)

type demo struct {
	Id   uint32
	Name string
}

func recvDemo(d demo, c connect.Connector) {
	fmt.Println(d.Id, d.Name)
}

```

### 注册路由

和注册数据包的意思一致，要注册才能被调用 支持匿名函数注册

```go
package main

import (
	"fmt"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

type demo struct {
	Id   uint32
	Name string
}

//定义函数
func recvDemo(d demo, c connect.Connector) {
	fmt.Println(d.Id, d.Name)
}

func main() {
	//路由注册，空结构体对象，函数体
	route.Register(demo{}, recvDemo)
	// 支持匿名函数注册
	route.Register(demo{}, func(d demo, c connect.Connector) {
		fmt.Println(d, c)
	})
}

```

### 运行服务

上面准备好后即可开始监听服务，目前支持 [Socket](../socket/server.go) 和 [Websocket](../websocket/server.go)

可以二选一，也可以同时开启，需要传递一个字符串形式的监听地址和端口，监听所有地址可以使用`:3000`,省略地址

用 goroutine 的方式运行其中任何一个服务，跳过阻塞，或者两个都可以用 goroutine 方式运行，用其它方式保持程序运行

批量监听多个端口:你可以多开几个 .Server() 传递不同的端口

```go
package main

import (
	"fmt"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"github.com/yanlong-li/hi-go-socket/socket"
	"github.com/yanlong-li/hi-go-socket/websocket"
)

type demo struct {
	Id   uint32
	Name string
}

//定义含糊
func recvDemo(d demo, c connect.Connector) {
	fmt.Println(d.Id, d.Name)
}

//初始化
func main() {
	//路由注册，空结构体对象，函数体
	route.Register(demo{}, recvDemo)
	// 支持匿名函数注册
	route.Register(demo{}, func(d demo, c connect.Connector) {
		fmt.Println(d, c)
	})

	//用 goroutine 的方式运行其中任何一个服务，跳过阻塞，或者两个都可以用 goroutine 方式运行，用其它方式保持程序运行
	go socket.Server("127.0.0.1:3000")
	go socket.Server(":3001")
	websocket.Server("127.0.0.1:3011")
}


```

### 系统保留操作码

系统保留 0- 5999 操作码 用于基础操作使用，虽然可能用不到几个，但是从6000开始 的 uint32 足够肆意挥霍了，如果不够用？这个系统架构支撑不了你的业务

    0(connect.Connector) 客户端连接后调用,只有一个 connect.Connector 参数
    1(uint64)   连接 ID，后期会改成弱化版的一个接口
    2(stream.BaseStream,connect.Connector)bool 接收前置,可以用于如权限判定等行为. BaseStream 中包含操作码和数据`[]byte`,return bool: true 继续,false 停止处理-丢弃该数据包
    3(stream.Interface)[]byte 发送前置 这里可以对 数据压缩和加密，并return []byte 返回改变后的数据

```go
package main

import (
	"fmt"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/packet"
	"github.com/yanlong-li/hi-go-socket/route"
	"github.com/yanlong-li/hi-go-socket/socket"
	"github.com/yanlong-li/hi-go-socket/stream"
)

type c struct {
}
type d struct {
}
type r struct {
}
type s struct {
}

//演示加密的模型
type BytesData struct {
	Data []byte
	//可以增加其他的诸如type标识加密类型，加密序号等
}

//初始化
func main() {
	//注册系统包，0,1,2,3
	packet.Register(packet.CONNECTION, c{})
	packet.Register(packet.DISCONNECTION, d{})
	packet.Register(packet.BEFORE_RECVING, r{})
	packet.Register(packet.BEFORE_SENDING, s{})
	// 普通路由，注册加密后的包操作码，自定义操作码9999
	packet.Register(9999, BytesData{})
	//新连接触发
	route.Register(c{}, func(connector connect.Connector) {
		fmt.Println("新客户端连接...")
	})
	//断开连接触发
	route.Register(d{}, func(id uint64) {
		fmt.Println("一个客户端断开...")
	})
	//收到数据包触发
	route.Register(r{}, func(bs stream.BaseStream, conn connect.Connector) bool {
		fmt.Println("权限验证通过...")
		return true
		//fmt.Println("权限验证失败...")
		//return false
	})
	//发送触发，加密
	route.Register(s{}, func(ps stream.Interface) []byte {
		//简单做了一层封装
		model := BytesData{}
		// 要调用 ToData获取标准的数据 []
		data := ps.ToData()
		//todo 这里可以对 model.Data 中的数据进行处理
		//将加密后的数据保存到数据包中
		model.Data = data
		// 注入填充好数据的加密模型
		ps.Marshal(model)
		// 获取注入后的标准数据并返回
		return ps.ToData()
	})
	//注册自定义路由 9999 ，用于解密操作
	route.Register(BytesData{}, func(bd BytesData, conn connect.Connector) {
		data := bd.Data
		//todo 解密 data
		//将解密后的数据传递给连接器，连接器会再次寻找路由
		conn.HandleData(data)
	})
	socket.Server("127.0.0.1:3000")
}

```

## 客户端

[demo](./demo/client/client.go)

### 运行服务

```go
package main

import "github.com/yanlong-li/hi-go-socket/socket"

func main() {
	//todo 定义数据包 同服务端
	//todo 注册数据包 同服务端
	//todo 定义路由 同服务端
	//todo 注册路由 同服务端
	//todo 系统路由 同服务端
	socket.Server("127.0.0.1:3000")
}
```

* 注意服务端和客户端的操作码定义，要保持一致，并且避免循环（心跳检测可以利用循环，但也要给个延迟）

更多请参考：

#### 客户端实现：

    https://github.com/Yanlong-LI/HelloWorldClient

#### 服务端实现：

    https://github.com/Yanlong-LI/HelloWorldServer