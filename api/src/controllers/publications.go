package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var pub models.Publication
	if err = json.Unmarshal(requestBody, &pub); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	pub.AuthorID = userId

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryPublication(db)
	pub.ID, err = repository.Create(pub)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, pub)
}

func SarchPublications(w http.ResponseWriter, r *http.Request) {}

func SearchPublication(w http.ResponseWriter, r *http.Request) {}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {}

func DeletePublication(w http.ResponseWriter, r *http.Request) {}
