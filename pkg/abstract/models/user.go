package models

type UserRole int

const (
	UserRoleLimitedUser UserRole = iota
	UserRoleAdministrator
)

type User struct {
	Role UserRole
	Name string
}
