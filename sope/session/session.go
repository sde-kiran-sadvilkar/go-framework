package session

import (
	"database/sql"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/postgresstore"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifeTime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
	DBPool         *sql.DB
}

func (s *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	minutes, err := strconv.Atoi(s.CookieLifeTime)

	if err != nil {
		minutes = 60
	}

	if strings.ToLower(s.CookiePersist) == "true" {
		persist = true
	} else {
		persist = false
	}

	if strings.ToLower(s.CookieSecure) == "true" {
		secure = true
	} else {
		secure = false
	}

	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = s.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = s.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	switch strings.ToLower(s.SessionType) {
	case "redis":

	case "mysql", "mariadb":
		session.Store = mysqlstore.New(s.DBPool)
	case "postgres", "postgresql":
		session.Store = postgresstore.New(s.DBPool)
	default:
		//cookie
	}

	return session

}
