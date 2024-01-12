package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/msproject2/internal/models"
	interfaces "github.com/msproject2/pkg"
	pb "github.com/msproject2/proto"
	"google.golang.org/grpc"
)

type UserServ struct {
	usecase interfaces.UseUsecase
	pb.UnimplementedUserServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseUsecase) {
	userGrpc := &UserServ{usecase: usecase}
	pb.RegisterUserServiceServer(grpcServer, userGrpc)
}

func (s *UserServ) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.SuccessResponse, error) {
	data := s.transformUserRPC(req)
	if data.Email == "" || data.Name == "" {
		return &pb.SuccessResponse{}, errors.New("Please Provide all fields")
	}
	_, err := s.usecase.Create(data)
	if err != nil {
		return &pb.SuccessResponse{}, fmt.Errorf("Can't Create User :%w", err)
	}
	return&pb.SuccessResponse{Response: "User created successfully"}, nil
}
func (s *UserServ) Get(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	id := req.GetId()
	if id == "" {
		return &pb.UserProfileResponse{}, fmt.Errorf("The userid cannot be blank")
	}
	user, err := s.usecase.Get(id)
	if err != nil {
		return &pb.UserProfileResponse{}, fmt.Errorf("Can't fetch UserId :%w", err)
	}
	return s.transformUserModel(user), nil
}
func (s *UserServ) transformUserRPC(req *pb.CreateUserRequest) models.User {
	return models.User{Name: req.GetName(), Email: req.GetEmail()}
}

func (s *UserServ) transformUserModel(user models.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{Id: string(rune(user.ID)), Name: user.Name, Email: user.Email}
}
