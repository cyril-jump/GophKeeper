package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
)

func (r *Provider) Create(ctx context.Context, user domain.User) error {
	crUserStmt, err := r.db.PrepareContext(ctx, "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id;")
	if err != nil {
		return &domain.StatementPSQLError{Err: err}
	}
	defer crUserStmt.Close()

	if err := crUserStmt.QueryRowContext(ctx, user.Login, user.Password).Scan(&user.ID); err != nil {
		errCode := err.(*pq.Error).Code
		if pgerrcode.IsIntegrityConstraintViolation(string(errCode)) {
			return &domain.AlreadyExistsError{Err: domain.ErrUserAlreadyExists}
		}
		return &domain.ExecutionPSQLError{Err: err}
	}

	return nil
}

func (r *Provider) GetByCredentials(ctx context.Context, login, password string) (domain.User, error) {
	user := domain.User{}

	getUserStmt, err := r.db.PrepareContext(ctx, "SELECT id,password FROM users WHERE login=$1;")
	if err != nil {
		return user, &domain.StatementPSQLError{Err: err}
	}
	defer getUserStmt.Close()

	if err := getUserStmt.QueryRowContext(ctx, login).Scan(&user.ID, &user.Password); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return user, &domain.NotFoundError{Err: domain.ErrUserNotFound}
		default:
			return user, &domain.ExecutionPSQLError{Err: err}
		}
	}

	if user.Password == password {
		user.Login = login
		return user, nil
	} else {
		return user, domain.ErrUserBadPassword
	}
}
