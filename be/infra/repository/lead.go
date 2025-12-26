package repository

import (
	"log/slog"

	"github.com/awe8128/arch-gen/be/infra/db"
)

type LeadRepository interface {
}

type LeadStore struct {
	db    *db.SQLStore
	attrs []slog.Attr
}

func NewLeadRepository(db *db.SQLStore) LeadRepository {
	return &LeadStore{
		db: db,
		attrs: []slog.Attr{
			slog.String("layer", "repository"),
			slog.String("domain", "task"),
		},
	}
}
