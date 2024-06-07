package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	UserId   uuid.UUID
	Fullname string
	Username string
	Gmail    string
	Password string
}

type Todo struct {
	TodoId     uuid.UUID
	Task       string
	CreatedAt  time.Time
	IsComleted bool
}

type GetTodosResp struct {
	Todos []Todo
	Count int
}
