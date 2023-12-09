package sessions

import (
	"errors"
	"github.com/MeM0rd/q-auth-svc/pkg/client/postgres"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"
)

func CreateSession(userId *int) (Session, error) {
	session := NewSession()
	var err error

	sessionToken := uuid.NewString()
	expiredAt := time.Now().Add(5 * time.Hour)

	q := `INSERT INTO sessions(user_id, token, expired_at) VALUES ($1, $2, $3) RETURNING id, token, user_id, expired_at`

	err = postgres.DB.QueryRow(q, userId, sessionToken, expiredAt).Scan(&session.Id, &session.Token, &session.UserId, &session.ExpiredAt)
	if err != nil {
		return Session{}, err
	}

	session.ExpiredAt = expiredAt

	return session, nil
}

func CheckSession(r *http.Request) error {
	session := NewSession()

	cookie, err := r.Cookie(os.Getenv("COOKIE_NAME"))
	if err != nil {
		return err
	}

	q := `SELECT id, token, expired_at, user_id FROM sessions WHERE token = $1`

	err = postgres.DB.QueryRow(q, cookie.Value).Scan(&session.Id, &session.Token, &session.ExpiredAt, &session.UserId)
	if err != nil {
		return err
	}

	if session.IsExpired() {
		q = `DELETE FROM sessions WHERE token = $1`

		postgres.DB.Exec(q, session.Token)
		return errors.New("token is expired")
	}

	return nil
}

func DeleteSession(token string) error {

	q := `DELETE FROM sessions WHERE token = $1`

	err := postgres.DB.QueryRow(q, token).Err()
	if err != nil {
		return err
	}

	return nil
}
