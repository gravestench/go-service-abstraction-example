package models

import (
	"database/sql"
)

type UserRole = int

const (
	UserRoleLimitedUser UserRole = iota + 1 // the database starts at index 1 !!
	UserRoleAdmin
)

type User struct {
	ID           int          `json:"id" db:"id"`
	Username     string       `json:"username" db:"username"`
	Password     string       `json:"password" db:"password"`
	EmailAddress string       `json:"email" db:"email_address"`
	FirstName    string       `json:"firstName" db:"first_name"`
	LastName     string       `json:"lastName" db:"last_name"`
	Role         UserRole     `json:"role" db:"role"`
	Active       bool         `json:"active" db:"active"`
	LastLogin    sql.NullTime `json:"lastLogin" db:"last_login"`
}
