package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jscmurph/blog_aggregator/internal/database"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
    type parameters struct {
        Name string `json:"name"`
        Url string `json:"url"`
    }

   decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters for handlerUsersCreate")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:       uuid.New() ,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      user.Name,
        Url: params.Url,
        UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create feed")
		fmt.Println(err)
		return
	}

    respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

