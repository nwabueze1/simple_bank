package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const minSecretKeySize = 32

// JWTMaker implements the Maker interface
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker with the given secret key
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("secret key cannot be %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey: secretKey}, nil
}

// CreateToken creates a new token for a specific username duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {

	now := time.Now()
	claims, err := NewPayload(username, duration)
	claims.IssuedAt = 

}

// verifyToken checks if token is valid
func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	var claims
}
