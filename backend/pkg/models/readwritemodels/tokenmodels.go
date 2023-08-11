package readwritemodels

type ContextKey int

type Header struct {
	Alg string
}

type Payload struct {
	UserId      string `json:"user_id"`
	DisplayName string `json:"display_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	IsLoggedIn  int    `json:"is_logged_in"`
	Role        int    `json:"role"`
	Exp         int64  `json:"exp"`
	Iat         int64  `json:"iat"`
}
