package main

import (
	"net/http"

	"github.com/jscmurph/blog_aggregator/internal/database"
)

func (cfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}
