// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: api_keys.sql

package database

import (
	"context"
)

const createAPIKey = `-- name: CreateAPIKey :one
INSERT INTO api_keys (key, created_at, expired_at) VALUES ($1, NOW(), NULL) RETURNING key, created_at, expired_at
`

func (q *Queries) CreateAPIKey(ctx context.Context, key string) (ApiKey, error) {
	row := q.db.QueryRowContext(ctx, createAPIKey, key)
	var i ApiKey
	err := row.Scan(&i.Key, &i.CreatedAt, &i.ExpiredAt)
	return i, err
}

const deleteExpiredAPIKeys = `-- name: DeleteExpiredAPIKeys :exec
DELETE FROM api_keys WHERE expired_at <= NOW()
`

func (q *Queries) DeleteExpiredAPIKeys(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteExpiredAPIKeys)
	return err
}

const expireAPIKey = `-- name: ExpireAPIKey :exec
UPDATE api_keys SET expired_at=NOW() WHERE key=$1
`

func (q *Queries) ExpireAPIKey(ctx context.Context, key string) error {
	_, err := q.db.ExecContext(ctx, expireAPIKey, key)
	return err
}

const getAPIKey = `-- name: GetAPIKey :one
SELECT key, created_at, expired_at FROM api_keys WHERE key=$1
`

func (q *Queries) GetAPIKey(ctx context.Context, key string) (ApiKey, error) {
	row := q.db.QueryRowContext(ctx, getAPIKey, key)
	var i ApiKey
	err := row.Scan(&i.Key, &i.CreatedAt, &i.ExpiredAt)
	return i, err
}
