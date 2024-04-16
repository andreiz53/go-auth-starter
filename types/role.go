package types

const (
	UserRoleSuperAdmin = iota + 1
	UserRoleAdmin
	UserRoleSupervizor
	UserRoleContributor
)

type Role struct {
	ID   int    `gorm:"primaryKey;not null;"`
	Name string `gorm:"not null;"`
}
