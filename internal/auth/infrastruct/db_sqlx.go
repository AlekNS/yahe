package infrastruct

import (
	"fmt"
	"net/url"

	// pq
	_ "github.com/lib/pq"

	"github.com/alekns/yahe/pkg/utils"
	"github.com/jmoiron/sqlx"
)

// NewDBSqlx .
func NewDBSqlx(url *url.URL) (*sqlx.DB, error) {
	dbUser, dbPassword := "", ""
	if url.User != nil && len(url.User.Username()) > 0 {
		dbUser = " user=" + url.User.Username()
		dbPassword = ""
		if password, exists := url.User.Password(); exists {
			dbPassword = "password=" + password
		}
	}

	var dialect = url.Scheme[:len(url.Scheme)]
	if dialect != "postgres" {
		return nil, fmt.Errorf("NewDBSqlx, Invalid sql dialect, only postgres is supported")
	}

	var dsn = fmt.Sprintf("host=%s port=%s %s %s dbname=%s sslmode=disable",
		utils.Coalesce(url.Hostname(), "localhost"),
		utils.Coalesce(url.Port(), "5432"),
		dbUser,
		dbPassword,
		utils.Coalesce(url.Path, "/postgres")[1:])

	conn, err := sqlx.Connect(dialect, dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
