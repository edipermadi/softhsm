package users

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetUser(ctx context.Context, username string) (*User, error)
	RecordLoginAttempt(ctx context.Context, userId int64, succeeded bool) error
}

func NewRepository(logger *logrus.Logger, db *sqlx.DB) Repository {
	return &repository{logger: logger, db: db}
}

type repository struct {
	logger *logrus.Logger
	db     *sqlx.DB
}

func (r *repository) GetUser(ctx context.Context, username string) (*User, error) {
	query := `
	SELECT
		u.id,
		u.type,
		u.username,
		up.hashed_password,
		sula.created_at AS last_succeeded_login_attempt,
		fula.created_at AS last_failed_login_attempt,
		up.expired_at   AS password_expired_at,
		u.created_at,
		u.updated_at
	FROM users u
		JOIN user_passwords up ON u.id = up.user_id
		LEFT JOIN user_login_attempts sula ON u.id = sula.user_id AND sula.succeeded
		LEFT JOIN user_login_attempts fula ON u.id = fula.user_id AND NOT fula.succeeded
	WHERE u.username = $1
	ORDER BY
	    up.id DESC,
		sula.id DESC,
		fula.id DESC
	LIMIT 1;`

	var user User
	err := r.db.GetContext(ctx, &user, query, username)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

func (r *repository) RecordLoginAttempt(ctx context.Context, userId int64, succeeded bool) error {
	query := `
		INSERT INTO user_login_attempts (user_id, succeeded)
		VALUES ( $1, $2 );`

	_, err := r.db.ExecContext(ctx, query, userId, succeeded)
	return err
}
