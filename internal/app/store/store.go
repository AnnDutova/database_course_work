package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //анонимный импорт, чтобы все методы сразу не перекидывать
)

type Store struct {
	config   *Config
	db       *sql.DB
	userrepo *UserRepository
}

func New(c *Config) *Store {
	return &Store{
		config: c,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userrepo != nil {
		return s.userrepo
	}
	s.userrepo = &UserRepository{
		store: s,
	}
	return s.userrepo
}
