package models

// Auth is used by users action.
// Userauth is Email or Phone user validation.
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
