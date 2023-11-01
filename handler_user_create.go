package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jscmurph/blog_aggregator/internal/database"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters for handlerUsersCreate")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Could not create user")
        fmt.Println(err)
        return
    }
	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}
