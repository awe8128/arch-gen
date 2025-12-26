package repository

import (
	"context"
	"log/slog"

	"github.com/awe8128/arch-gen/be/infra/db"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name *string, id uint) (*string, uuid.UUID)
}

type UserStore struct {
	db    *db.SQLStore
	attrs []slog.Attr
}

func NewUserRepository(db *db.SQLStore) UserRepository {
	return &UserStore{
		db: db,
		attrs: []slog.Attr{
			slog.String("layer", "repository"),
			slog.String("domain", "task"),
		},
	}
}
