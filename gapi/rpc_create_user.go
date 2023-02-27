package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/maan19/bank-app-go/db/sqlc"
	"github.com/maan19/bank-app-go/pb"
	"github.com/maan19/bank-app-go/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Error(codes.AlreadyExists, "user already exists")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return resp, nil
}
