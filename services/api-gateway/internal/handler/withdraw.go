package handler

import (
	"encoding/json"
	"io"
	"net/http"

	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
)

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	var i struct {
		AccountId   string `json:"accountId" validate:"required"`
		Amount      uint32 `json:"amount" validate:"required"`
		Description string `json:"description"`
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

	_, err = h.transactionProcessor.Withdraw(r.Context(), &pb.WithdrawRequest{
		UserId:      &userId,
		AccountId:   &i.AccountId,
		Amount:      &i.Amount,
		Description: &i.Description,
	})
	if err != nil {
		writeResp(w, err.Error(), http.StatusInternalServerError, nil, nil)
		return
	}

	writeResp(w, "success", http.StatusOK, nil, nil)
}
