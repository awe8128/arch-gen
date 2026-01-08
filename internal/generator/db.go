package generator

func NewInitDB() (string, string) {
	filename := "db.go"
	template := `
package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"main.go/utils/config"
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
	return template, filename
}
