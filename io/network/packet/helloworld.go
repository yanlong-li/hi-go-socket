package packet

import "fmt"

type HelloWorld struct {
	Name string
	Age  uint8
	Sex  bool
}

func handleHelloWorld(world HelloWorld) {
	fmt.Println(world)
}
