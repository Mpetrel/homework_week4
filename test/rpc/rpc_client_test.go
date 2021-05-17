package main

import (
	"context"
	"google.golang.org/grpc"
	v1 "homework_week4/api/user/v1"
	"testing"
)

// test rpc call

func TestAddUser(t *testing.T) {
	var expectedCode int32 = 0
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	c := v1.NewUserClient(conn)
	r, err := c.AddUser(context.Background(), &v1.AddUserRequest{
		Name: "Larry",
		Sex:  "male",
		Age:  20,
	})
	if err != nil {
		t.Error(err)
	}
	if expectedCode != r.GetCode() {
		t.Errorf("add user expected code %d, got %d\n, message: %s", expectedCode, r.GetCode(), r.GetMessage())
	}
}

func TestFetchUser(t *testing.T) {
	var minSize = 1
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	c := v1.NewUserClient(conn)
	r, err := c.FetchUser(context.Background(), &v1.EmptyMessage{})
	if err != nil {
		t.Error(err)
	}
	if minSize > len(r.Users) {
		t.Errorf("fetch users expected min size %d, got %d\n", minSize, len(r.GetUsers()))
	}
}