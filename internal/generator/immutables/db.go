package immutables

import (
	"path/filepath"

	"github.com/awe8128/arch-gen/utils/fs"
)

func GenerateDBInit(root string) error {
	path := filepath.Join(root, "infra", "db")
	filename := "db.go"
	content := `
package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)
func Init(ctx context.Context, config *config.Config) (*SQLStore, error) {

dbConfig, err := pgxpool.ParseConfig(config.GetDSN())
if err != nil {
	return nil, err
}
conn, err := pgxpool.NewWithConfig(ctx, dbConfig)
if err != nil {
	return nil, err
}
err = conn.Ping(ctx)
if err != nil {
	return nil, err
}
store := NewStore(conn)

return store, nil
}`
	if err := fs.GenerateFile(content, path, filename); err != nil {
		return err
	}
	return nil
}
