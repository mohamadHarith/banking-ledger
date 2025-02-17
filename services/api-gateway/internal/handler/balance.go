package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	var i struct {
		AccountId string `json:"accountId" validate:"required"`
	}

	userId := r.Context().Value(userIdKey).(string)

	req, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeResp(w, err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	err = json.Unmarshal(req, &i)
	if err != nil {
		writeResp(w, err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	if err := h.validator(&i); err != nil {
		writeResp(w, err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	balance, err := h.repository.GetUserBalance(r.Context(), userId, i.AccountId)
	if err != nil {
		writeResp(w, err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	type o struct {
		Balance uint32 `json:"balance"`
	}

	writeResp(w, "success", http.StatusOK, o{Balance: balance}, nil)
}
