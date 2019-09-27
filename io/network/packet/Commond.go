package packet

func init() {
	register(0x04, GetUserList{})
	register(0x05, UserList{List: make([]string, 0)})
}

type GetUserList struct {
}

type UserList struct {
	List []string
}
