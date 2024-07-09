package server

import (
	"context"
	"testing"

	"grpc-user-service/proto"
)

func TestGetUserByID(t *testing.T) {
	s := UserServiceServer{
		users: []proto.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "Nikita", City: "DL", Phone: 9233567890, Height: 5.0, Married: false},
			{Id: 3, Fname: "John", City: "NY", Phone: 8272882826, Height: 5.3, Married: true},
		},
	}

	req := &proto.UserIDRequest{Id: 1}
	res, err := s.GetUserByID(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if res.User.Fname != "Steve" {
		t.Fatalf("expected user Steve, got %v", res.User.Fname)
	}
}

func TestGetUsersByIDs(t *testing.T) {
	s := UserServiceServer{
		users: []proto.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "Nikita", City: "DL", Phone: 9233567890, Height: 5.0, Married: false},
			{Id: 3, Fname: "John", City: "NY", Phone: 8272882826, Height: 5.3, Married: true},
		},
	}

	req := &proto.UserIDsRequest{Ids: []int32{1, 2}}
	res, err := s.GetUsersByIDs(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(res.Users) != 2 {
		t.Fatalf("expected 2 users, got %v", len(res.Users))
	}
}

func TestSearchUsers(t *testing.T) {
	s := UserServiceServer{
		users: []proto.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "Nikita", City: "DL", Phone: 9233567890, Height: 5.0, Married: false},
			{Id: 3, Fname: "John", City: "NY", Phone: 8272882826, Height: 5.3, Married: true},
		},
	}

	req := &proto.SearchRequest{City: "LA"}
	res, err := s.SearchUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(res.Users) != 1 {
		t.Fatalf("expected 1 user, got %v", len(res.Users))
	}

	if res.Users[0].Fname != "Steve" {
		t.Fatalf("expected user Steve, got %v", res.Users[0].Fname)
	}
}
