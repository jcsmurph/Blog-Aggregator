package main

import (
	"fmt"
	"net/http"

)

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

    feeds, err := cfg.DB.GetAllRssFeeds(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create user")
		fmt.Println(err)
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedsToFeeds(feeds))
}
