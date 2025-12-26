package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	CreateUser(ctx context.Context, name *string, id uint) (*string, uuid.UUID)
}
