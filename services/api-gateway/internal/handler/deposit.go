package handler

import (
	"encoding/json"
	"io"
	"net/http"

	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
)

func (h *Handler) Deposit(w http.ResponseWriter, r *http.Request) {
	var i struct {
		UserId      string `json:"userId" validate:"required"`
		AccountId   string `json:"accountId" validate:"required"`
		Amount      uint32 `json:"amount" validate:"required"`
		Description string `json:"description"`
	}

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

	_, err = h.transactionProcessor.Deposit(r.Context(), &pb.DepositRequest{
		UserId:      &i.UserId,
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
