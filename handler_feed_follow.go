package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jscmurph/blog_aggregator/internal/database"
)

func (cfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Feed_id uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters for handlerCreateFeedFollow")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.Feed_id,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Was unable to follow feed as requested")
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))

}

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowString := chi.URLParam(r, "feedFollowID")


    feedFollowID, err := uuid.Parse(feedFollowString)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
        return
    }

    err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams {
        UserID: user.ID,
        FeedID: feedFollowID,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Could not delete feed follow")
        return
    }

    respondWithStatus(w, http.StatusOK)

}

func (cfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

    feedFollows, err := cfg.DB.GetFeedFollowByUser(r.Context(), user.ID)

    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Could not get feeds for user")
        return
    }

    respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}
