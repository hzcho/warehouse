package repository

import (
	"auth/internal/domain/model"
	"auth/internal/repository/converter"
	"auth/internal/repository/dao"

	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	usersTable = "users"
	rolesTable = "roles"
)

type User struct {
	pool *pgxpool.Pool
}

func NewUser(pool *pgxpool.Pool) *User {
	return &User{
		pool: pool,
	}
}

func (r *User) Get(ctx context.Context, login string) (model.User, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return model.User{}, err
	}
	defer conn.Release()

	var daoUser dao.User
	userQuery :=
		`SELECT id, login, pass_hash, role_id, refresh_token, token_expire 
        FROM users 
        WHERE login = $1`

	row := conn.QueryRow(ctx, userQuery, login)
	if err := row.Scan(&daoUser.ID, &daoUser.Login, &daoUser.PassHash, &daoUser.RoleId, &daoUser.RefreshToken, &daoUser.TokenExpiry); err != nil {
		return model.User{}, err
	}

	var daoRole dao.Role
	roleQuery :=
		`SELECT id, role_name 
        FROM roles 
        WHERE id = $1`

	roleRow := conn.QueryRow(ctx, roleQuery, daoUser.RoleId)
	if err := roleRow.Scan(&daoRole.ID, &daoRole.Role); err != nil {
		return model.User{}, err
	}

	return converter.ToDomainUser(daoUser, daoRole), nil
}

func (r *User) Create(ctx context.Context, user model.User) (uuid.UUID, error) {
	userID := uuid.New()

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	var roleID uuid.UUID
	err = tx.QueryRow(ctx, "SELECT id FROM roles WHERE role_name = $1", user.Role).Scan(&roleID)
	if err != nil {
		return uuid.Nil, err
	}

	var insertedID string
	query :=
		`INSERT INTO users (id, login, pass_hash, role_id) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id`

	err = tx.QueryRow(ctx, query, userID.String(), user.Login, user.Password, roleID).Scan(&insertedID)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

func (r *User) Update(ctx context.Context, user model.User) error {
	var roleID uuid.UUID
	err := r.pool.QueryRow(ctx, "SELECT id FROM roles WHERE role_name = $1", user.Role).Scan(&roleID)
	if err != nil {
		return err
	}

	query :=
		`UPDATE users 
        SET login = $1, 
            pass_hash = $2, 
            role_id = $3, 
            refresh_token = $4, 
            token_expire = $5 
        WHERE id = $6`

	_, err = r.pool.Exec(ctx, query,
		user.Login,
		user.Password,
		roleID,
		user.RefreshToken,
		user.TokenExpiry,
		user.ID,
	)

	return err
}

func (r *User) Delete(ctx context.Context, id uuid.UUID) error {
	query :=
		`DELETE FROM users 
        WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, id)
	return err
}
