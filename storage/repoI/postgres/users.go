package postgres

import (
	"app/models"
	repoi "app/storage/repoI"
	"context"
	"log"

	"github.com/jackc/pgx"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn) repoi.UserRepoI {

	userRepo := &UserRepo{
		conn: conn,
	}
	return userRepo

}

func (u *UserRepo) CreateUser(ctx context.Context, user models.User) error {

	query := `
	INSERT INTO 
		users(
			user_id,
			fullname,
			username,
			gmail,
			password
		) VALUES (
			$1,$2,$3,$4,$5
		)`

	_, err := u.conn.Exec(
		ctx, query,
		user.UserId,
		user.Fullname,
		user.Username,
		user.Gmail,
		user.Password,
	)

	if err != nil {
		log.Println("error on executing !!!", err)
		return err
	}
	return nil
}

func (u *UserRepo) GetUsers(ctx context.Context, limit, page string) ([]models.User, error) {
	return nil, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	return nil, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, user models.User) error {
	return nil
}

func (u *UserRepo) DeleteUserById(ctx context.Context, userId string) error {
	return nil
}
