package userauth

import (
	"log"
	"time"

	pb "github.com/lolmourne/go-accounts/rpc"
	"google.golang.org/grpc"
)

type AuthClient struct {
	host    string
	timeout time.Duration
}

type AuthClientGRPC struct {
	grpcCli pb.AccountsClient
}

type User struct {
	UserID     int64     `json:"user_id"`
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
	ProfilePic string    `json:"profile_pic"`
}

type ClientItf interface {
	GetUserInfo(accessToken string) *User
	GetUserByID(userID int64) *User
}

func NewClient(host string, timeout time.Duration) ClientItf {
	return &AuthClient{
		host:    host,
		timeout: timeout,
	}
}

func NewGrpcClient(address string) ClientItf {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Cannot init grpc")
		return nil
	}
	defer conn.Close()
	grpcCli := pb.NewAccountsClient(conn)
	return &AuthClientGRPC{
		grpcCli: grpcCli,
	}
}
