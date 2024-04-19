package db

type Storage interface {
	UserStorage
}

type Store struct {
	User UserStorage
}

func NewStore(us UserStorage) *Store {
	return &Store{
		User: us,
	}
}
