package entities

import (
	"time"

	"github.com/google/uuid"
)

type Tweet struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func NewTweet() *Tweet {
	now := time.Now()
	tweet := Tweet{
		ID:        uuid.NewString(),
		CreatedAt: now.UTC().String(),
	}

	return &tweet
}
