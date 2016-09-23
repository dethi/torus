package torus

import "time"

var AnonUser = User{}

type User struct {
	Id           uint64
	Email        string
	PasswordHash string

	CreationTime time.Time

	IsAdmin bool
}
