package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/dto"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/authentication_service_proto"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var i struct {
		FullName string `json:"fullName" validate:"required"`
		UserName string `json:"userName" validate:"required"`
		Password string `json:"password" validate:"required"`
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

	_, err = h.authenticator.CreateUser(r.Context(), &pb.CreateUserRequest{
		Username: i.UserName,
		FullName: i.FullName,
		Password: i.Password,
	})
	if err != nil {
		writeResp(w, err.Error(), http.StatusInternalServerError, nil, nil)
		return
	}

	writeResp(w, "success", http.StatusOK, nil, nil)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var i struct {
		UserName string `json:"userName" validate:"required"`
		Password string `json:"password" validate:"required"`
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

	resp, err := h.authenticator.Login(r.Context(), &pb.LoginRequest{
		Username: i.UserName,
		Password: i.Password,
	})
	if err != nil {
		writeResp(w, err.Error(), http.StatusInternalServerError, nil, nil)
		return
	}

	var o struct {
		AccessToken string `json:"accessToken"`
	}

	o.AccessToken = resp.Token

	writeResp(w, "success", http.StatusOK, dto.ResponseMessage{
		Item: o,
	}, nil)
}
