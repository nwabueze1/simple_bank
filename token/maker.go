package token

import "time"

//Maker is an interface for managing token
type Maker interface {
	//CreateToken creates a new token for a specific username duration
	CreateToken(username string, duration time.Duration)(string, error)
	//verifyToken checks if token is valid
	VerifyToken(token string) (*Payload, error)
}