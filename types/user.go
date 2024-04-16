package types

import (
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	minUsernameLength = 3
	minPasswordLength = 8
	minNameLength     = 2
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique;" json:"username"`
	Email    string `gorm:"not null;unique;" json:"email"`
	Password string `gorm:"not null;" json:"-"`
	Name     string `gorm:"not null;" json:"name"`
	RoleID   int    `json:"-"`
	Role     Role   `json:"-"`
}

type UpdateUserParams struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type UserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (p UserParams) Validate() map[string]string {
	var errors = map[string]string{}
	if _, err := mail.ParseAddress(p.Email); err != nil {
		errors["email"] = strings.Split(err.Error(), "mail: ")[1]
	}
	if len(p.Name) <= minNameLength {
		errors["name"] = fmt.Sprintf("name too short, should be at least %d characters", minNameLength)
	}
	if len(p.Username) <= minUsernameLength {
		errors["username"] = fmt.Sprintf("username too short, should be at least %d characters", minUsernameLength)
	}
	if len(p.Password) <= minPasswordLength {
		errors["password"] = fmt.Sprintf("password too short, should be at least %d characters", minPasswordLength)
	}
	return errors
}

func NewUserFromParams(u UserParams) (*User, error) {
	encPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: u.Username,
		Email:    u.Email,
		Password: string(encPassword),
		Name:     u.Name,
		RoleID:   UserRoleContributor,
	}, nil
}

func NewAdminFromParams(u UserParams) (*User, error) {
	encPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: u.Username,
		Email:    u.Email,
		Password: string(encPassword),
		Name:     u.Name,
		RoleID:   UserRoleAdmin,
	}, nil
}
