package repoI

import (
	"app/models"
	"context"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUsers(ctx context.Context, limit, page string) ([]models.User, error)
	GetUserById(ctx context.Context, userId string) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUserById(ctx context.Context, userId string) error
}

type TodoRepoI interface {
	CreateTodo(ctx context.Context, todo models.User) error
	GetTodos(ctx context.Context, limit, page string) (*models.GetTodosResp, error)
	GetTodoById(ctx context.Context, todoId string) (*models.User, error)
	UpdateTodo(ctx context.Context, todo models.User) error
	DeleteTodoById(ctx context.Context, todoId string) error
}
