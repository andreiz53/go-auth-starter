package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/andreiz53/go-auth-starter/db"
	"github.com/andreiz53/go-auth-starter/types"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func TestCreateUser(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := NewUserHandler(tdb)
	app.Post("/", uh.HandleCreateUser)

	params := types.UserParams{
		Username: "test",
		Email:    "test@gmail.com",
		Password: "test1ngasdasd",
		Name:     "Testing dude",
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fail()
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fail()
	}
	slog.Info(fmt.Sprintf("%+v", res.Status))

}

type testdb struct {
	db.UserStorage
}

func setup(t *testing.T) *testdb {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Could not load .env file")
	}
	tdb, err := db.NewSQLUserStore(os.Getenv("DB_NAME_TESTING"))
	if err != nil {
		t.Fatal("Could not create user store")
	}
	db.CreateTables(tdb.DB)
	db.SeedRoles(tdb.DB)

	return &testdb{
		UserStorage: &db.SQLUserStore{
			DB: tdb.DB,
		},
	}
}
func (tdb *testdb) destroy(t *testing.T) {
	if err := tdb.UserStorage.Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
}
