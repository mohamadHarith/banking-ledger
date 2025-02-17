package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/dto"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
)

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var i struct {
		InitialBalance uint32 `json:"initialBalance" validate:"required,min=1"`
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

	protoResp, err := h.transactionProcessor.CreateAccount(r.Context(), &pb.CreateAccountRequest{
		UserId:         &userId,
		InitialBalance: &i.InitialBalance,
	})
	if err != nil {
		writeResp(w, err.Error(), http.StatusInternalServerError, nil, nil)
		return
	}

	log.Println("here")

	account := dto.FromProtoToEntityAccount(protoResp.Account)

	log.Println("here2=> ", account.CreatedAt.Location())

	writeResp(w, "success", http.StatusOK, account, nil)
}
