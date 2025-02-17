package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/dto"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_logger_proto"
)

func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	var i struct {
		AccountId string `json:"accountId" validate:"required"`
		Page      uint32 `json:"page" validate:"required"`
	}

	ctx := r.Context()

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

	resp, err := h.transactionLogger.GetTransactionLogs(ctx, &pb.GetTransactionLogsRequest{
		AccountId: &i.AccountId,
		UserId:    &userId,
		Page:      &i.Page,
	})
	if err != nil {
		writeResp(w, err.Error(), http.StatusInternalServerError, nil, nil)
		return
	}

	writeResp(w, "success", http.StatusOK, dto.FromProtoToEntityTransactionLogs(resp.TransactionLogs), nil, dto.Pagination{
		TotalRecords: *resp.TotalRecords,
		CurrentPage:  *resp.CurrentPage,
		TotalPages:   *resp.TotalPages,
		NextPage:     *resp.NextPage,
		PrevPage:     *resp.PrevPage,
	})
}
