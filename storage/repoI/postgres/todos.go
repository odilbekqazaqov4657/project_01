package postgres

import (
	"app/models"
	repoi "app/storage/repoI"
	"context"

	"github.com/jackc/pgx"
)

type TodoRepo struct {
	conn *pgx.Conn
}

func NewTodoRepo(conn *pgx.Conn) repoi.TodoRepoI {
	todoRepo := &TodoRepo{}

	return todoRepo
}

func (t *TodoRepo) CreateTodo(ctx context.Context, todo models.User) error {
	return nil
}

func (t *TodoRepo) GetTodos(ctx context.Context, limit, page string) (*models.GetTodosResp, error) {
	return nil, nil
}

func (t *TodoRepo) GetTodoById(ctx context.Context, todoId string) (*models.User, error) {
	return nil, nil
}

func (t *TodoRepo) UpdateTodo(ctx context.Context, todo models.User) error {
	return nil
}

func (t *TodoRepo) DeleteTodoById(ctx context.Context, todoId string) error {
	return nil
}
