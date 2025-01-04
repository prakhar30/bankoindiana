// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: verify_emails.sql

package db

import (
	"context"
)

const createVerifyEmail = `-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (username, email, secret_code) VALUES ($1, $2, $3) RETURNING id, username, email, secret_code, is_used, created_at, expired_at
`

type CreateVerifyEmailParams struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	SecretCode string `json:"secret_code"`
}

func (q *Queries) CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error) {
	row := q.db.QueryRow(ctx, createVerifyEmail, arg.Username, arg.Email, arg.SecretCode)
	var i VerifyEmail
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.SecretCode,
		&i.IsUsed,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}
