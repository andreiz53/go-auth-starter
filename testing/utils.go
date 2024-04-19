package test

import (
	"context"
	"os"
	"testing"

	"github.com/andreiz53/go-auth-starter/db"
	"github.com/joho/godotenv"
)

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
	db.SeedTestUser(tdb.DB)

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
