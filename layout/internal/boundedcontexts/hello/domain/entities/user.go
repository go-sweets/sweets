package entities

type User struct {
	ID       int64
	NickName string
}

func NewUser(id int64, nickName string) *User {
	return &User{
		ID:       id,
		NickName: nickName,
	}
}
