package packet

func init() {
	register(0x00, HelloWorld{})
	register(0x01, Login{})
	register(0x02, Token{})
	register(0x03, LoginFail{})
}

// hello world
type HelloWorld struct {
	Message string
}

// login
type Login struct {
	Username string
	Password string
}

type Token struct {
	Token string
}

type LoginFail struct {
	Code    int32
	Message string
}
