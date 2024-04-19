package main

import (
	"log"
	"os"

	"github.com/andreiz53/go-auth-starter/api"
	"github.com/andreiz53/go-auth-starter/db"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

var serverConfig = fiber.Config{
	ErrorHandler: api.ErrorHandler,
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	userStore, err := db.NewSQLUserStore(os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	userStore.Init()

	var (
		store = db.NewStore(userStore)
		app   = fiber.New(serverConfig)
		uh    = api.NewUserHandler(store.User)
		ah    = api.NewAuthHandler(store.User)
		auth  = app.Group("/api/auth")
		apiv1 = app.Group("/api/v1", api.JWTAuth(store.User))
		admin = apiv1.Group("/admin", api.AdminAuth)
	)
	auth.Post("/", ah.HandleAuthenticate)

	apiv1.Get("/user", uh.HandleGetAllUsers)
	apiv1.Get("/user/:id", uh.HandleGetUser)
	apiv1.Post("/user", uh.HandleCreateUser)
	apiv1.Delete("/user/:id", uh.HandleDeleteUser)
	apiv1.Put("/user/:id", uh.HandleUpdateUser)

	admin.Get("/", uh.HandleGetAllUsers)

	log.Fatal(app.Listen(os.Getenv("API_URL")))
}
