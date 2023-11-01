package main

import (
	"net/http"

	"github.com/jscmurph/blog_aggregator/internal/auth"
	"github.com/jscmurph/blog_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        apiKey, err := auth.GetAPIKey(r.Header)
        if err != nil {
            respondWithError(w, http.StatusUnauthorized, "Could not find API Key")
            return
        }

        user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
        if err != nil {
            respondWithError(w, http.StatusNotFound, "Could not get User")
            return
        }

        handler(w,r,user)
    }
}
