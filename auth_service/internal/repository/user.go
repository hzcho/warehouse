package repository

import (
	"auth/internal/domain/model"
	"auth/internal/domain/net/request"
	"auth/internal/repository/converter"
	"auth/internal/repository/dao"
	"fmt"
	"strings"

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

func (r *User) GetUsers(ctx context.Context, filter request.GetUsersFilter) ([]model.User, error) {
	query := fmt.Sprintf(`
		SELECT u.id, u.login, u.pass_hash, r.role_name, u.phone_number, u.email, u.refresh_token, u.token_expire
		FROM %s AS u
		LEFT JOIN %s AS r ON u.role_id = r.id
	`, usersTable, rolesTable)

	whereClauses := []string{}
	args := []interface{}{}
	argID := 1

	if filter.Role != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("r.role_name = $%d", argID))
		args = append(args, filter.Role)
		argID++
	}

	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		var role string
		if err := rows.Scan(
			&user.ID, &user.Login, &user.Password, &role, &user.PhoneNumber, &user.Email, &user.RefreshToken, &user.TokenExpiry,
		); err != nil {
			return nil, err
		}
		user.Role = role
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}

func (r *User) Get(ctx context.Context, login string) (model.User, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return model.User{}, err
	}
	defer conn.Release()

	var daoUser dao.User
	userQuery :=
		`SELECT id, login, pass_hash, role_id, phone_number, email, refresh_token, token_expire 
        FROM users 
        WHERE login = $1`

	row := conn.QueryRow(ctx, userQuery, login)
	if err := row.Scan(&daoUser.ID, &daoUser.Login, &daoUser.PassHash, &daoUser.RoleId, &daoUser.PhoneNumber, &daoUser.Email, &daoUser.RefreshToken, &daoUser.TokenExpiry); err != nil {
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
		`INSERT INTO users (id, login, pass_hash, role_id, phone_number, email) 
        VALUES ($1, $2, $3, $4, $5, $6) 
        RETURNING id`

	err = tx.QueryRow(ctx, query, userID.String(), user.Login, user.Password, roleID, user.PhoneNumber, user.Email).Scan(&insertedID)
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
			phone_number = $4,
			email = $5,
            refresh_token = $6, 
            token_expire = $7 
        WHERE id = $8`

	_, err = r.pool.Exec(ctx, query,
		user.Login,
		user.Password,
		roleID,
		user.PhoneNumber,
		user.Email,
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
