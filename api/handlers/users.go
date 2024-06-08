package handlers

import (
	"app/models"
	"context"
	"fmt"

	"github.com/gofrs/uuid"
)

var ctx =context.Background()

func (h *Handlers)CreateUser(){

	var User models.User

	User.UserId=uuid.New()

	fmt.Print("enter username: ")
	fmt.Scanln(&User.Username)
		
	fmt.Print("enter fullname: ")
	fmt.Scanln(&User.fullname)
		
	fmt.Print("enter gmail: ")
	fmt.Scanln(&User.Gmail)
		
	fmt.Print("enter password: ")
	fmt.Scanln(&User.Password)

	h.storage.GetUserRepo().CreateUser(ctx, User)
}