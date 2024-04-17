package api

import (
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

func (params *AuthParams) Validate() bool {
	if params.Email == "" || params.Password == "" {
		return false
	}
	return true
}

func (ah *AuthHandler) HandleAuthenticate(ctx fiber.Ctx) error {
	var params AuthParams
	if err := ctx.Bind().Body(&params); err != nil || params.Email == "" {
		return ErrBadRequest()
	}
	if !params.Validate() {
		return ErrBadRequest()
	}

	user, err := ah.userStore.GerUserByEmail(ctx.Context(), params.Email)
	if err != nil {
		return ErrNotFound()
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return ErrBadRequest()
	}
	token := GenerateTokenFromUser(user)
	if token == "" {
		return ErrUnauthorized()
	}
	res := AuthResponse{
		User:  user,
		Token: token,
	}
	return ctx.JSON(res)
}
