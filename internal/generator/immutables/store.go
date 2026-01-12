package immutables

func StoreTemplate() (string, string) {
	filename := "store.go"

	content := `
package db

type Store interface {
	sqlc.Queries
}

type SQLStore struct {
	connPool *pgxpool.Pool
	*sqlc.Queries
}

func NewStore(connPool *pgxpool.Pool) *SQLStore {
	return &SQLStore{
		connPool: connPool,
		Queries:  sqlc.New(connPool),
	}
}

func (s *SQLStore) Ping(ctx context.Context) error {
	err := s.connPool.Ping(ctx)
	if err != nil {
		return err
	}
	return nil
}

	`

	return content, filename
}
