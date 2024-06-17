package main

import (
	"fmt"
	"net/http"

	"github.com/Rohan556/rss-generator/internal/auth"
	"github.com/Rohan556/rss-generator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAuthAPI(r.Header)

		if err != nil {
			handleError(w, 400, fmt.Sprintf("Error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			handleError(w, 401, fmt.Sprintf("%v", "Authorization failed"))
			return
		}

		handler(w, r, user)
	}
}
