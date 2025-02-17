package handler

import (
	"context"
	"net/http"
	"strings"

	pb "github.com/mohamadHarith/banking-ledger/shared/proto/authentication_service_proto"
)

type contextKey string

const userIdKey contextKey = "userId"

func (h *Handler) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			writeResp(w, "access token not found", http.StatusUnauthorized, nil, nil)
			return
		}

		token := parts[1]

		resp, err := h.authenticator.ValidateToken(context.Background(), &pb.ValidateRequest{
			Token: token,
		})
		if err != nil {
			writeResp(w, err.Error(), http.StatusUnauthorized, nil, nil)
			return
		}

		if !resp.Valid {
			writeResp(w, "access token expired", http.StatusUnauthorized, nil, nil)
			return
		}

		ctx := context.WithValue(r.Context(), userIdKey, resp.UserId)
		next(w, r.WithContext(ctx))
	}
}
