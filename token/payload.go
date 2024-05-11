package token

import (
	"time"

	"github.com/google/uuid"
)

// Payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// creates a new token payload
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	creationTime := time.Now()

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  creationTime,
		ExpiredAt: creationTime.Add(duration),
	}

	return payload, nil
}
