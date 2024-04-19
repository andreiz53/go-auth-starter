package api

import (
	"net/http"

	"github.com/andreiz53/go-auth-starter/db"
	"github.com/andreiz53/go-auth-starter/types"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userStore db.UserStorage
}

func NewUserHandler(store db.UserStorage) *UserHandler {
	return &UserHandler{
		userStore: store,
	}
}

func (h *UserHandler) HandleGetAllUsers(ctx fiber.Ctx) error {
	users, err := h.userStore.GetAllUsers(ctx.Context())
	if err != nil {
		return ErrInternalError()
	}
	return ctx.JSON(users)
}
func (h *UserHandler) HandleGetUser(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := h.userStore.GetUserByID(ctx.Context(), id)
	if err != nil {
		return ErrInvalidID()
	}
	return ctx.JSON(user)
}

func (h *UserHandler) HandleCreateUser(ctx fiber.Ctx) error {
	var u types.CreateUserParams
	if err := ctx.Bind().Body(&u); err != nil {
		return ErrBadRequest()
	}

	if errs := u.Validate(); len(errs) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(errs)
	}

	user, err := types.NewUserFromParams(u)
	if err != nil {
		return ErrBadRequest()
	}

	newUser, err := h.userStore.CreateUser(ctx.Context(), user)
	if err != nil {
		return ErrInternalError()
	}

	return ctx.JSON(newUser)
}

func (h *UserHandler) HandleDeleteUser(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	if err := h.userStore.DeleteUser(ctx.Context(), id); err != nil {
		return ErrInvalidID()
	}

	return ctx.JSON(map[string]string{"message": "user deleted"})
}

func (h *UserHandler) HandleUpdateUser(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	var u types.UpdateUserParams
	if err := ctx.Bind().Body(&u); err != nil {
		return ErrBadRequest()
	}
	err := h.userStore.UpdateUser(ctx.Context(), u, id)
	if err != nil {
		return ErrInternalError()
	}

	return ctx.JSON(map[string]string{"message": "user updated"})
}
