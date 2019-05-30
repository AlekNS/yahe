package services

import (
	"database/sql"
	"strings"

	"github.com/alekns/yahe/internal/auth/app"
	"github.com/alekns/yahe/internal/auth/config"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type (
	userSqlx struct {
		ID           string
		Name         string
		Login        string
		PasswordHash string
		IsActive     bool
	}

	userRepositoryServiceSqlx struct {
		db       *sqlx.DB
		settings *config.UsersSettings
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

	row := ur.db.QueryRow(
		ur.db.Rebind(ur.settings.SQLUserByLogin), login)
	err := row.Scan(&userDto.ID, &userDto.Name, &userDto.Login, &userDto.PasswordHash, &userDto.IsActive)
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

	row := ur.db.QueryRow(ur.db.Rebind(
		ur.settings.SQLUserByID), userID)
	err := row.Scan(&userDto.ID, &userDto.Name, &userDto.Login, &userDto.PasswordHash, &userDto.IsActive)
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
		_, err := ur.db.Exec(ur.db.Rebind(
			ur.settings.SQLUserInsert), uuidValue, user.Name, user.Login, user.Password, user.IsActive)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				return nil, app.ErrorUserAlreadyExists
			}
			return nil, err
		}

		user.ID = uuidValue
		return user, nil
	}

	_, err := ur.db.Exec(ur.db.Rebind(
		ur.settings.SQLUserUpdate), user.Password, user.IsActive, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// NewUserRepositoryServiceSqlx .
func NewUserRepositoryServiceSqlx(settings *config.UsersSettings, db *sqlx.DB) app.UserRepositoryService {
	return &userRepositoryServiceSqlx{
		db:       db,
		settings: settings,
	}
}
