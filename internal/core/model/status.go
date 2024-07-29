package model

type UserStatus int

const (
	Passive UserStatus = iota
	Active
	Deleted
)

func (status UserStatus) String() string {
	return [...]string{"passive", "active", "deleted"}[status]
}
