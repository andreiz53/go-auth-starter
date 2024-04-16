package api

import (
	"fmt"

	"github.com/andreiz53/go-auth-starter/db"
	"github.com/andreiz53/go-auth-starter/types"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userStore db.UserStorage
}

type AuthResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

func NewAuthHandler(store db.UserStorage) *AuthHandler {
	return &AuthHandler{
		userStore: store,
	}
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ah *AuthHandler) HandleAuthenticate(ctx fiber.Ctx) error {
	var params AuthParams
	if err := ctx.Bind().Body(&params); err != nil {
		return err
	}

	user, err := ah.userStore.GerUserByEmail(ctx.Context(), params.Email)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return err
	}
	token := GenerateTokenFromUser(user)
	if token == "" {
		return fmt.Errorf("unauthorized here")
	}
	res := AuthResponse{
		User:  user,
		Token: token,
	}
	return ctx.JSON(res)
}
