package userauth

import (
	"context"
	"log"
	"time"

	pb "github.com/lolmourne/go-accounts/rpc"
)

func (cg *AuthClientGRPC) GetUserInfo(accessToken string) *User {
	req := &pb.GetUserInfoRequest{
		AccessToken: accessToken,
	}

	resp, err := cg.grpcCli.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &User{
		UserID:     resp.UserID,
		Username:   resp.UserName,
		CreatedAt:  time.Unix(resp.TimeStamp, 0),
		ProfilePic: resp.ProfilePicture,
	}
}

func (cg *AuthClientGRPC) GetUserByID(userID int64) *User {
	req := &pb.GetUserByIDRequest{
		UserID: userID,
	}

	resp, err := cg.grpcCli.GetUserByID(context.Background(), req)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &User{
		UserID:     resp.UserID,
		Username:   resp.UserName,
		CreatedAt:  time.Unix(resp.TimeStamp, 0),
		ProfilePic: resp.ProfilePicture,
	}
}
