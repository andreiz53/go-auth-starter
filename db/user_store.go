package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/andreiz53/go-auth-starter/types"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dropper interface {
	Drop(ctx context.Context) error
}

// TODO: ID must be string everywhere
// TODO: get rid of id in UpdateUser
type UserStorage interface {
	Dropper

	CreateUser(ctx context.Context, user *types.User) (*types.User, error)
	GetUserByID(ctx context.Context, id int) (*types.User, error)
	GerUserByEmail(ctx context.Context, email string) (*types.User, error)
	GetAllUsers(ctx context.Context) ([]*types.User, error)
	UpdateUser(ctx context.Context, values types.UpdateUserParams, idStr string) error
	DeleteUser(ctx context.Context, id int) error
}

type SQLUserStore struct {
	DB *gorm.DB
}

func NewSQLUserStore(dbname string) (*SQLUserStore, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/?parseTime=true", DB_USER, DB_PASS, DB_HOST)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		return &SQLUserStore{}, err
	}

	db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbname))
	db.Exec(fmt.Sprintf("USE %s;", dbname))

	return &SQLUserStore{DB: db}, nil
}

func (s *SQLUserStore) Init() {
	CreateTables(s.DB)
	SeedRoles(s.DB)
	SeedSuperAdmin(s.DB)
}

func (s *SQLUserStore) GetUserByID(ctx context.Context, id int) (*types.User, error) {
	var user types.User
	tx := s.DB.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
func (s *SQLUserStore) GerUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	tx := s.DB.First(&user, "email = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (s *SQLUserStore) GetAllUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	result := s.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *SQLUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	tx := s.DB.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (s *SQLUserStore) UpdateUser(ctx context.Context, values types.UpdateUserParams, idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("could not update user")
	}
	tx := s.DB.Model(&types.User{}).Where("id = ?", id).Updates(values)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("could not update user")
	}

	return nil
}

func (s *SQLUserStore) DeleteUser(ctx context.Context, id int) error {
	tx := s.DB.Delete(&types.User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("could not find user")
	}
	return nil
}

func (s *SQLUserStore) Drop(ctx context.Context) error {
	err := godotenv.Load("../.env")
	if err != nil {
		slog.Info(fmt.Sprintf("Could not load .env file %+v", err))
		return err
	}
	sql, err := s.DB.DB()
	if err != nil {
		return err
	}
	_, err = sql.Exec("DROP DATABASE " + os.Getenv("DB_NAME_TESTING") + ";")
	if err != nil {
		return err
	}
	return nil
}
