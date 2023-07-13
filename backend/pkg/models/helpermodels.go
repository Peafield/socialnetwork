package models

// helper is an interface of helper functions
type Helper interface {
	IsAlphaNumeric(s string) (bool, error)
	IsValidPath(path string) (bool, error)
}
