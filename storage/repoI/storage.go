package storage

import (
	"app/storage/repoI"
	"app/storage/repoI/postgres"

	"github.com/jackc/pgx"
)

type StorageI interface {
	GetUserRepo() repoi.UserRepoI
	GetTodoRepo() repoI.TodoRepoI
}

type Storage struct {
	UserRepo repoi.UserRepoI
	TodoRepo repoi.TodoRepoI
}

func NewStorage(conn *pgx.Conn) StorageI {
	storage := &storage{
		UserRepo: postgres.NewUserRepo(conn),
		TodoRepo: postgres.NewTodoRepo(conn),
	}
	return storage
}

func (s *Storage) GetUserRepo() repoi.UserRepoI {
	return s.UserRepo
}

func (s *Storage) GetTodoRepo() repoi.TodoRepoI {
	return s.TodoRepo
}
