package controllers

import "net/http"

// Create
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
}

// Search all
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
}

// Search one
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
}

// Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
}

// Delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
}
