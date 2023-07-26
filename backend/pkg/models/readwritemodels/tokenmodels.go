package readwritemodels

import "time"

type ContextKey string

type Header struct {
	Alg string
}

type Payload struct {
	UserId    string    `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      int       `json:"role"`
	Exp       time.Time `json:"exp"`
	Iat       time.Time `json:"iat"`
}
