package users

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	id        uuid.UUID
	userId    uuid.UUID
	token     string
	expiresAt time.Time
	createdAt time.Time
}

func NewToken(id, userId uuid.UUID, token string, expiresAt, createdAt time.Time) *Token {
	return &Token{
		id:        id,
		userId:    userId,
		token:     token,
		expiresAt: expiresAt,
		createdAt: createdAt,
	}
}

func (t *Token) ID() uuid.UUID {
	return t.id
}

func (t *Token) UserId() uuid.UUID {
	return t.userId
}

func (t *Token) Token() string {
	return t.token
}

func (t *Token) ExpiresAt() time.Time {
	return t.expiresAt
}

func (t *Token) CreatedAt() time.Time {
	return t.createdAt
}
