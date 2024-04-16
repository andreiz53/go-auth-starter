package api

import (
	"github.com/andreiz53/go-auth-starter/types"
	"github.com/gofiber/fiber/v3"
)

func AdminAuth(ctx fiber.Ctx) error {
	user, ok := ctx.Context().UserValue("user").(*types.User)
	if !ok {
		return ErrUnauthorized()
	}
	if user.RoleID > types.UserRoleAdmin {
		return ErrUnauthorized()
	}
	return ctx.Next()
}
