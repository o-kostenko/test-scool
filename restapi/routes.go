package restapi

import (
	"net/http"

	"github.com/gorilla/mux"

	"test-school/services"
)

const APIKey = "api-key"

func NewAPI(s services.Services) *mux.Router {
	h := NewHandler(s)
	router := mux.NewRouter().StrictSlash(true)
	router.Use(h.Auth)

	router.HandleFunc("/profile/{id}", h.GetProfileByID).Methods("GET")
	router.HandleFunc("/profile", h.GetProfileList).Methods("GET")

	return router
}

// Auth middleware function return false if not exist api-key in database
func (h *Handlers) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			key = r.Header.Get(APIKey)
			ctx = r.Context()
		)

		ok, err := h.service.GetAuthKey(ctx, key)
		if err != nil {
			responseWithJSON(w, http.StatusInternalServerError, "parse api-key")

			return
		}

		if !ok {
			responseWithJSON(w, http.StatusForbidden, "wrong api-key")

			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
