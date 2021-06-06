package server

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/lolmourne/go-accounts/resource/acc"
	pb "github.com/lolmourne/go-accounts/rpc"
	"github.com/lolmourne/go-accounts/usecase/userauth"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	pb.UnimplementedAccountsServer
	userAuthUc userauth.UsecaseItf
	dbRsc      acc.DBItf
}

func (gs *GrpcServer) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	accessToken := req.AccessToken

	if len(accessToken) < 1 {
		return nil, errors.New("Access token is empty")
	}

	userID, err := gs.userAuthUc.ValidateSession(accessToken)
	if err != nil {
		return nil, errors.New("Access token is not valid")
	}

	user, err := gs.dbRsc.GetUserByUserID(userID)
	if err != nil {
		return nil, errors.New("User Not Found")
	}

	return &pb.GetUserInfoResponse{
		UserID:         user.UserID,
		ProfilePicture: user.ProfilePic,
		UserName:       user.Username,
		TimeStamp:      user.CreatedAt.Unix(),
	}, nil
}

func (gs *GrpcServer) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	return nil, nil
}

func NewGrpcServer(userAuthUc userauth.UsecaseItf, dbRsc acc.DBItf) *GrpcServer {
	return &GrpcServer{
		userAuthUc: userAuthUc,
		dbRsc:      dbRsc,
	}
}

func InitGrpcServer(userAuthUc userauth.UsecaseItf, dbRsc acc.DBItf) {
	log.Println("Starting GRPC Server at 7575")
	lis, err := net.Listen("tcp", ":7575")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountsServer(s, &GrpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
