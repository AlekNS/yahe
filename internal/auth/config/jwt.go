package config

import (
	"encoding/hex"
	"net/url"

	"github.com/spf13/viper"
)

type (
	// JwtSettings .
	JwtSettings struct {
		StorageURLString  string
		KeysPrefix        string
		Algorithm         string
		DefaultExpiration int
		SecretHex         string

		StorageURL *url.URL
		Secret     []byte
	}
)

func jwtSettingsGetAndValidate(v *viper.Viper) *JwtSettings {
	const (
		kjwtstorageurl = "jwt.storageurl"
		kjwtalgorithm  = "jwt.algorithm"
		kjwtkeyprefix  = "jwt.keyprefix"
		kjwtdefaultexp = "jwt.defaultexpiration"
		kjwtsecrethex  = "jwt.secrethex"
	)

	// Setup defaults
	v.SetDefault(kjwtstorageurl, "redis://localhost:6379/1")
	v.SetDefault(kjwtalgorithm, "HS256")
	v.SetDefault(kjwtdefaultexp, 604800)
	v.SetDefault(kjwtkeyprefix, "pa.")

	// Validations
	storageURL, err := url.Parse(v.GetString(kjwtstorageurl))
	if err != nil {
		panic(kjwtstorageurl + " has invalid url format")
	}

	if v.GetInt(kjwtdefaultexp) < 60 {
		panic(kjwtdefaultexp + " has invalid value")
	}

	secret, err := hex.DecodeString(v.GetString(kjwtsecrethex))
	if err != nil {
		panic(kjwtsecrethex + " has invalid hex format")
	}

	return &JwtSettings{
		StorageURLString:  v.GetString(kjwtstorageurl),
		StorageURL:        storageURL,
		Algorithm:         v.GetString(kjwtalgorithm),
		KeysPrefix:        v.GetString(kjwtkeyprefix),
		DefaultExpiration: v.GetInt(kjwtdefaultexp),
		SecretHex:         v.GetString(kjwtsecrethex),
		Secret:            secret,
	}
}
