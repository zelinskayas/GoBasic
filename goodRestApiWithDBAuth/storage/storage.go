package storage

import (
	"database/sql"
	_ "github.com/microsoft/go-mssqldb" //для того чтобы отработала функция init()
	"log"
)

// instance of storage
type Storage struct {
	config *Config
	//Database FileDescriptor
	db *sql.DB
	//subfield for repo interfacing (model user)
	userRepository *UserRepository
	//subfield for repo interfacing (model article)
	articleRepository *ArticleRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("sqlserver", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfully!")
	return nil
}

// close connection
func (storage *Storage) Close() {
	storage.db.Close()
}

// public repo for article
func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		storage: s,
	}
	return s.articleRepository
}

// public repo for user
func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return s.userRepository
}
