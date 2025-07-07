package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	if err := pub.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

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

func SearchPublications(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	pubID, err := strconv.ParseUint(parameters["pubID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryPublication(db)
	pub, err := repository.GetById(pubID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, pub)
}

func SearchPublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repositories.NewRepositoryPublication(db)
	pubs, err := repository.Search(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, pubs)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	parameters := mux.Vars(r)
	pubID, err := strconv.ParseUint(parameters["pubID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryPublication(db)
	pubSavedInDatabase, err := repository.GetById(pubID)
	if err != nil {
		responses.Error(w, http.StatusNotFound, err)
		return
	}

	if pubSavedInDatabase.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível atulizar uma publicação que não é sua"))
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

	pub.ID = pubID

	if err = pub.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(pubID, pub); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletePublication(w http.ResponseWriter, r *http.Request) {}
