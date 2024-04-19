package api

import (
	"fmt"
	"os"
	"time"

	"github.com/andreiz53/go-auth-starter/db"
	"github.com/andreiz53/go-auth-starter/types"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	ID    string `json:"id"`
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
}

func JWTAuth(userStore db.UserStorage) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		tokenArr := ctx.GetReqHeaders()["X-Jwt"]
		if len(tokenArr) == 0 {
			return ErrUnauthorized()
		}
		token := string(tokenArr[0])
		claims, err := ValidateToken(token)
		if err != nil {
			return ErrUnauthorized()
		}

		expFloat := claims["exp"].(float64)
		exp := int64(expFloat)
		if time.Now().Unix() > exp {
			return ErrUnauthorized()
		}

		u := claims["id"].(string)
		user, err := userStore.GetUserByID(ctx.Context(), u)
		if err != nil {
			return ErrUnauthorized()
		}
		ctx.Context().SetUserValue("user", user)
		return ctx.Next()
	}
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnauthorized()
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, ErrUnauthorized()
	}
	if !token.Valid {
		return nil, ErrUnauthorized()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrUnauthorized()
	}
	return claims, nil
}

func GenerateTokenFromUser(user *types.User) string {
	claims := JWTClaims{
		ID:    fmt.Sprint(user.ID),
		Email: user.Email,
		Exp:   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return ""
	}
	return tokenStr
}
