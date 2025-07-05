package models

// Represent the format to change password
type Password struct {
	New     string `json:new`
	Current string `json:current`
}
