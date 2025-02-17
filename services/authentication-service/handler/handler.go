package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mohamadHarith/banking-ledger/services/authentication-service/repository"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/authentication_service_proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	pb.UnimplementedAuthServiceServer
	repository *repository.Repository
}

const SecretKey = "hkjb9724$"

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repository: repo,
	}
}

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	u := entity.User{}
	u.Id = uuid.NewString()
	u.FullName = req.FullName
	u.Username = req.Username

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u.Password = string(hash)

	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now

	err = h.repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	user, err := h.repository.GetUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Second * 15).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	resp := &pb.LoginResponse{
		Token: signedToken,
	}

	return resp, nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	token, err := jwt.ParseWithClaims(req.Token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	t, err := token.Claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	if time.Now().After(t.Time) {
		return nil, errors.New("access token expired")
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		return nil, err
	}

	resp := &pb.ValidateResponse{
		UserId: sub,
		Valid:  true,
	}

	return resp, nil
}
