package postgres

import (
	"context"
	"errors"
)

func CheckUsername(username string) (bool, error) {
	if pool == nil {
		return false, errors.New("database not initialized")
	}

	var exists bool
	err := pool.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)",
		username).Scan(&exists)

	if err != nil {
		return false, err
	}

	return !exists, nil
}
