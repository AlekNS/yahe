package config

import (
	"encoding/hex"
	"net/url"

	"github.com/spf13/viper"
)

type (
	// UsersSettings .
	UsersSettings struct {
		StorageURLString string
		PasswordsSaltHex string

		SQLUserByLogin string
		SQLUserByID    string
		SQLUserInsert  string
		SQLUserUpdate  string

		StorageURL    *url.URL
		PasswordsSalt []byte
	}
)

func usersSettingsGetAndValidate(v *viper.Viper) *UsersSettings {
	const (
		kuserstorageurl     = "users.storageurl"
		kuserspasswordssalt = "users.passwordsalthex"

		kuserssqluserbylogin = "users.sql.userByLogin"
		kuserssqluserbyid    = "users.sql.userById"
		kuserssqluserinsert  = "users.sql.userInsert"
		kuserssqluserupdate  = "users.sql.userUpdate"
	)

	// Setup defaults
	v.SetDefault(kuserstorageurl, "postgres://postgres:postgres@localhost:5432/postgres")

	v.SetDefault(kuserssqluserbylogin, `SELECT u.id, u.name, u.login, u.password_hash, u.is_active FROM "users" u where u.login = ? LIMIT 1`)
	v.SetDefault(kuserssqluserbyid, `SELECT u.id, u.name, u.login, u.password_hash, u.is_active FROM "users" u where u.id = ? LIMIT 1`)
	v.SetDefault(kuserssqluserinsert, `INSERT INTO "users" (id, created_at, updated_at, name, login, password_hash, is_active) VALUES (?, now(), now(), ?, ?, ?, ?)`)
	v.SetDefault(kuserssqluserupdate, `UPDATE "users" SET updated_at = now(), password_hash = ?, is_active = ? WHERE id = ?`)

	// Validations
	storageURL, err := url.Parse(v.GetString(kuserstorageurl))
	if err != nil {
		panic(kuserstorageurl + " has invalid url format")
	}

	salt, err := hex.DecodeString(v.GetString(kuserspasswordssalt))
	if err != nil {
		panic(kuserspasswordssalt + " has invalid hex format")
	}

	return &UsersSettings{
		StorageURLString: v.GetString(kuserstorageurl),
		PasswordsSaltHex: v.GetString(kuserspasswordssalt),

		SQLUserByLogin: v.GetString(kuserssqluserbylogin),
		SQLUserByID:    v.GetString(kuserssqluserbyid),
		SQLUserInsert:  v.GetString(kuserssqluserinsert),
		SQLUserUpdate:  v.GetString(kuserssqluserupdate),

		StorageURL:    storageURL,
		PasswordsSalt: salt,
	}
}
