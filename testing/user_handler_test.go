package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreiz53/go-auth-starter/api"
	"github.com/andreiz53/go-auth-starter/types"
	"github.com/gofiber/fiber/v3"
)

func TestCreateUserSuccessfully(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := api.NewUserHandler(tdb)
	app.Post("/", uh.HandleCreateUser)

	params := types.CreateUserParams{
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
	if err != nil || res.StatusCode != http.StatusOK {
		t.Fail()
	}

}

func TestCreateUserWithNoEmail(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := api.NewUserHandler(tdb)
	app.Post("/", uh.HandleCreateUser)

	params := types.CreateUserParams{
		Username: "test",
		Email:    "",
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
	if err != nil || res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestCreateUserWithBadEmail(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := api.NewUserHandler(tdb)
	app.Post("/", uh.HandleCreateUser)

	params := types.CreateUserParams{
		Username: "test",
		Email:    "thisiswrong",
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
	if err != nil || res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestCreateUserWithNoPassword(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := api.NewUserHandler(tdb)
	app.Post("/", uh.HandleCreateUser)

	params := types.CreateUserParams{
		Username: "test",
		Email:    "test@test.ts",
		Password: "",
		Name:     "Testing dude",
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fail()
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil || res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestCreateUserWithBadPassword(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := api.NewUserHandler(tdb)
	app.Post("/", uh.HandleCreateUser)

	params := types.CreateUserParams{
		Username: "test",
		Email:    "test@test.ts",
		Password: "as",
		Name:     "Testing dude",
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fail()
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil || res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestCreateUserWithNoUsername(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := api.NewUserHandler(tdb)
	app.Post("/", uh.HandleCreateUser)

	params := types.CreateUserParams{
		Username: "",
		Email:    "test@test.ts",
		Password: "password",
		Name:     "Testing dude",
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fail()
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil || res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestCreateUserWithBadUsername(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := api.NewUserHandler(tdb)
	app.Post("/", uh.HandleCreateUser)

	params := types.CreateUserParams{
		Username: "a",
		Email:    "test@test.ts",
		Password: "asasdasdasd",
		Name:     "Testing dude",
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fail()
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil || res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestCreateUserWithNoName(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := api.NewUserHandler(tdb)
	app.Post("/", uh.HandleCreateUser)

	params := types.CreateUserParams{
		Username: "asdasdasdasd",
		Email:    "test@test.ts",
		Password: "password",
		Name:     "",
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fail()
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil || res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestCreateUserWithBadName(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := api.NewUserHandler(tdb)
	app.Post("/", uh.HandleCreateUser)

	params := types.CreateUserParams{
		Username: "aasdasdasdsad",
		Email:    "test@test.ts",
		Password: "asasdasdasdasd",
		Name:     "t",
	}

	body, err := json.Marshal(params)
	if err != nil {
		t.Fail()
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil || res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGetUser(t *testing.T) {

	tdb := setup(t)
	defer tdb.destroy(t)

	app := fiber.New()
	uh := api.NewUserHandler(tdb)
	app.Get("/:id", uh.HandleGetUser)

	req := httptest.NewRequest("GET", "/1", nil)
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil || res.StatusCode != http.StatusOK {
		t.Fail()
	}
}
