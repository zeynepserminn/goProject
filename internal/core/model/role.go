package model

type UserRole int

const (
	RoleAdmin UserRole = iota
	RoleUser
)

func (role UserRole) String() string {
	return [...]string{"admin", "user"}[role]
}
