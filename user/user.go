package user

import (
	"errors"
	"time"
)

var (
	ErrGet      = errors.New("could not get user")
	ErrSave     = errors.New("could not get user")
	ErrNotFound = errors.New("could not find user")
)

type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
}
