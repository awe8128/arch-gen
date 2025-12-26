package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	CreateUser(ctx context.Context, id uint, name *string) (*string, uuid.UUID)
}
