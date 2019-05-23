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

		StorageURL    *url.URL
		PasswordsSalt []byte
	}
)

func usersSettingsGetAndValidate(v *viper.Viper) *UsersSettings {
	const (
		kusersstorageurl     = "users.storageurl"
		kusersspasswordssalt = "users.passwordsalthex"
	)

	// Setup defaults

	// Validations
	storageURL, err := url.Parse(v.GetString(kusersstorageurl))
	if err != nil {
		panic(kusersstorageurl + " has invalid url format")
	}

	salt, err := hex.DecodeString(v.GetString(kusersspasswordssalt))
	if err != nil {
		panic(kusersspasswordssalt + " has invalid hex format")
	}

	return &UsersSettings{
		StorageURLString: v.GetString(kusersstorageurl),
		PasswordsSaltHex: v.GetString(kusersspasswordssalt),

		StorageURL:    storageURL,
		PasswordsSalt: salt,
	}
}
