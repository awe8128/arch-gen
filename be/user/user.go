package user

import "context"

type User struct {
	Id   uint
	Name *string
	Age  int64
}

func NewUser(ctx context.Context, id uint, name *string, age int64) *User {

	ent := &User{
		Id:   id,
		Name: name,
		Age:  age,
	}

	return ent
}
