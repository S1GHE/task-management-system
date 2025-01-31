package users

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"task-management-system-api/internal/domains/users"
	"task-management-system-api/internal/infrastructure"
	"task-management-system-api/pkg/request"
	"time"
)

type userDB struct {
	id       uuid.UUID
	email    string
	password string
	created  time.Time
}

type tokenDB struct {
	id        uuid.UUID
	userId    uuid.UUID
	token     string
	expiresAt time.Time
	createdAt time.Time
}

type PostgresRepo struct {
	db *pgxpool.Pool
	qd squirrel.StatementBuilderType
	infrastructure.UsersRepo
}

func NewPostgresRepo(db *pgxpool.Pool, qb squirrel.StatementBuilderType) *PostgresRepo {
	return &PostgresRepo{
		db: db,
		qd: qb,
	}
}

func (r *PostgresRepo) FindByToken(token string) (uuid.UUID, error) {
	res := r.qd.Select("user_id").
		From("USERS.AUTH").
		Where(squirrel.Eq{"token": "$1"}).
		Where("expires_at > NOW()")
	query, args, err := res.ToSql()
	if err != nil {
		return uuid.Nil, err
	}

	var userId uuid.UUID
	if err := r.db.QueryRow(context.Background(), query, args...).Scan(&userId); err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}

func (r *PostgresRepo) CreateToken(userID uuid.UUID, token string) error {
	expiresAt := time.Now().UTC().Add(24 * time.Hour)

	res := r.qd.Insert("USERS.AUTH").
		Columns("user_id", "token", "expires_at").
		Values(userID, token, expiresAt)

	query, args, err := res.ToSql()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	_, err = r.db.Exec(context.Background(), query, args...)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (r *PostgresRepo) CreateUser(user request.CreateUser) error {
	res := r.qd.Insert("USERS.CLIENT").
		Columns("email", "password_hash").
		Values(user.Email, user.Password)

	query, args, err := res.ToSql()
	if err != nil {
		return fmt.Errorf("build insert query error: %w", err)
	}

	_, err = r.db.Exec(context.Background(), query, args...)
	if err != nil {
		return fmt.Errorf("execute query error: %w", err)
	}

	return nil
}

func (r *PostgresRepo) GetUserByEmail(email string) (*users.User, error) {
	res := r.qd.Select("id", "email", "password_hash").
		From("USERS.CLIENT").
		Where(squirrel.Eq{"email": email})

	query, args, err := res.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build insert query error: %w", err)
	}

	var userdb userDB
	err = r.db.QueryRow(context.Background(), query, args...).Scan(&userdb.id, &userdb.email, &userdb.password)
	if err != nil {
		return nil, fmt.Errorf("execute query error: %w", err)
	}

	user, err := users.NewUser(userdb.id, userdb.email, userdb.password, userdb.created)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	return user, nil
}
