package models

import "fmt"

// Credentials signup structure
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Signup with credentials TODO May move direct to model
func (c *Credentials) Signup(name, email, password string) (bool, error) {
	return false, fmt.Errorf("signups disabled")
}
