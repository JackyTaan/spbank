package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	//CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	//Verifytoken checks if the token is valud or not
	VerifyToken(token string) (*Payload, error)
}
