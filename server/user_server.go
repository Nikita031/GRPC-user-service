package server

import (
	"context"
	"errors"
	"grpc-user-service/proto"
	"strings"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
	users []proto.User
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{
		users: []proto.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "Nikita", City: "DL", Phone: 9233567890, Height: 5.0, Married: false},
			{Id: 3, Fname: "John", City: "NY", Phone: 8272882826, Height: 5.3, Married: true},

			// Add more users as needed
		},
	}
}

func (s *UserServiceServer) GetUserByID(ctx context.Context, req *proto.UserIDRequest) (*proto.UserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return &proto.UserResponse{User: &user}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserServiceServer) GetUsersByIDs(ctx context.Context, req *proto.UserIDsRequest) (*proto.UsersResponse, error) {
	var users []*proto.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				users = append(users, &user)
			}
		}
	}
	return &proto.UsersResponse{Users: users}, nil
}

func (s *UserServiceServer) SearchUsers(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	var users []*proto.User
	for _, user := range s.users {
		if (req.Fname == "" || strings.EqualFold(user.Fname, req.Fname)) &&
			(req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(req.Height == 0 || user.Height == req.Height) {
			users = append(users, &user)
		}
	}
	return &proto.SearchResponse{Users: users}, nil
}
