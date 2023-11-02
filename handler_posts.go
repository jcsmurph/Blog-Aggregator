package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jscmurph/blog_aggregator/internal/database"
)


func (cfg *apiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {

    limitStr := r.URL.Query().Get("limit")
    limit := 10

    if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
        limit = specifiedLimit
    }

    posts, err := cfg.DB.GetPostByUser(r.Context(), database.GetPostByUserParams{
       UserID: user.ID,
       Limit: int32(limit),
    })

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get posts")
		fmt.Println(err)
		return
	}


	respondWithJSON(w, http.StatusCreated, databasePostsToPosts(posts))
}
