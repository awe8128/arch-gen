package user

import "context"

type User struct {
	Name *string
	Age  int64
	Id   uint
}

func NewUser(ctx context.Context, age int64, id uint, name *string) *User {
	return nil
}
