package db

import (
	"os"

	"github.com/andreiz53/go-auth-starter/types"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	db.FirstOrCreate(&types.Role{ID: types.UserRoleSuperAdmin, Name: "Super Admin"})
	db.FirstOrCreate(&types.Role{ID: types.UserRoleAdmin, Name: "Admin"})
	db.FirstOrCreate(&types.Role{ID: types.UserRoleSupervizor, Name: "Supervizor"})
	db.FirstOrCreate(&types.Role{ID: types.UserRoleContributor, Name: "Contributor"})
}

func SeedSuperAdmin(db *gorm.DB) {
	password, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("SUPER_ADMIN_PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		panic("failed to encrypt password when seeding super admin")
	}
	superAdmin := types.User{
		Username: os.Getenv("SUPER_ADMIN_USERNAME"),
		Email:    os.Getenv("SUPER_ADMIN_EMAIL"),
		Password: string(password),
		Name:     os.Getenv("SUPER_ADMIN_NAME"),
		RoleID:   types.UserRoleSuperAdmin,
	}
	db.FirstOrCreate(&superAdmin)
}

func SeedTestUser(db *gorm.DB) {
	password, err := bcrypt.GenerateFromPassword([]byte("passw0rdbr0"), bcrypt.DefaultCost)
	if err != nil {
		panic("failed to encrypt password when seeding super admin")
	}
	user := types.User{
		Username: "testing",
		Email:    "testing@testing.testing",
		Password: string(password),
		Name:     "testing",
		RoleID:   types.UserRoleAdmin,
	}
	db.FirstOrCreate(&user)
}
