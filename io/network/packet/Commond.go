package packet

func init() {
	register(0x04, GetUserList{})
	register(0x05, UserList{List: make([]string, 0)})
	register(0x06, Disconnect{})
}

type GetUserList struct {
}

type UserList struct {
	List []string
}
type Disconnect struct {
}
