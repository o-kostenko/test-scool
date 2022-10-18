package restapi

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"test-school/services"
)

type Handlers struct {
	service services.Services
}

func NewHandler(s services.Services) *Handlers {
	return &Handlers{service: s}
}

func (h *Handlers) GetProfileByID(w http.ResponseWriter, r *http.Request) { //
	var (
		vars = mux.Vars(r)
		key  = vars["id"]
		ctx  = r.Context()
	)

	userID, err := strconv.Atoi(key)
	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, "error parse user id")

		return
	}

	profile, err := h.service.GetProfile(ctx, userID)
	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, err)
	}

	responseWithJSON(w, http.StatusOK, profile)
}

func (h *Handlers) GetProfileList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profiles, err := h.service.GetProfileList(ctx)
	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, err)
	}

	responseWithJSON(w, http.StatusOK, profiles)
}

func responseWithJSON(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
