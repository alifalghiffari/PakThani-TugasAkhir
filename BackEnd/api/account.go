package api

import (
	"net/http"
	"strconv"
)

type Account struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type AccountErrorResponse struct {
	Error string `json:"error"`
}

type AccountSuccessResponse struct {
	Success string `json:"success"`
}

// GetAccount returns a single account.
func GetAccount(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	account, err := db.GetAccount(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, account)
}
