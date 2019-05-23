package services

import (
	"database/sql"
	"strings"

	"github.com/alekns/yahe/internal/auth/app"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type (
	userSqlx struct {
		ID           string `db:"id"`
		Name         string `db:"name"`
		Login        string `db:"login"`
		PasswordHash string `db:"password_hash"`
		IsActive     bool   `db:"is_active"`
	}

	userRepositoryServiceSqlx struct {
		db *sqlx.DB
	}
)

func userSqlxToAuthUser(userDto *userSqlx) *app.User {
	return &app.User{
		ID:       userDto.ID,
		Name:     userDto.Name,
		Login:    userDto.Login,
		Password: userDto.PasswordHash,
		IsActive: userDto.IsActive,
	}
}

// GetByLogin .
func (ur *userRepositoryServiceSqlx) GetByLogin(tenantID app.TenantID,
	login string) (*app.User, error) {

	var userDto = &userSqlx{}

	var err = ur.db.Get(userDto, ur.db.Rebind(`
SELECT u.id, u.name, u.login, u.password_hash, u.is_active
FROM "users" u where u.login = ? LIMIT 1
`), login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return userSqlxToAuthUser(userDto), nil
}

// GetByIDs .
func (ur *userRepositoryServiceSqlx) GetByIDs(tenantID app.TenantID,
	userID app.UserID) (*app.User, error) {

	var userDto = &userSqlx{}

	var err = ur.db.Get(userDto, ur.db.Rebind(`
SELECT u.id, u.name, u.login, u.password_hash, u.is_active
FROM "users" u where u.id = ? LIMIT 1
`), userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return userSqlxToAuthUser(userDto), nil
}

// Save .
func (ur *userRepositoryServiceSqlx) Save(user *app.User) (*app.User, error) {
	if len(user.ID) == 0 {
		uuidValue := uuid.NewV4().String()
		_, err := ur.db.Exec(ur.db.Rebind(`
INSERT INTO "users" (id, created_at, updated_at, name, login, password_hash)
VALUES (?, now(), now(), ?, ?, ?)
`), uuidValue, user.Name, user.Login, user.Password)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				return nil, app.ErrorUserAlreadyExists
			}
			return nil, err
		}

		user.ID = uuidValue
		return user, nil
	}

	_, err := ur.db.Exec(ur.db.Rebind(`
UPDATE "users" SET updated_at = now(), password_hash = ?
WHERE id = ?
`), user.Password, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// NewUserRepositoryServiceSqlx .
func NewUserRepositoryServiceSqlx(db *sqlx.DB) app.UserRepositoryService {
	return &userRepositoryServiceSqlx{
		db: db,
	}
}
