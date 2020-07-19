package main

import (
	"context"
	"log"

	"github.com/neocortical/mysvc"
	mysvcgrpc "github.com/neocortical/mysvc/grpc"
)

// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
	userService mysvc.Service
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService mysvc.Service) mysvcgrpc.UserServiceServer {
	return &userServiceController{
		userService: userService,
	}
}

// GetUsers calls the core service's GetUsers method and maps the result to a grpc service response.
func (ctlr *userServiceController) GetUsers(ctx context.Context, req *mysvcgrpc.GetUsersRequest) (resp *mysvcgrpc.GetUsersResponse, err error) {
	resultMap, err := ctlr.userService.GetUsers(req.GetIds())
	if err != nil {
		return
	}

	resp = &mysvcgrpc.GetUsersResponse{}
	for _, u := range resultMap {
		resp.Users = append(resp.Users, marshalUser(&u))
	}

	log.Printf("handled GetUsers(%v)\n", req.GetIds())
	return
}

// marshalUser marshals a business object User into a gRPC layer User.
func marshalUser(u *mysvc.User) *mysvcgrpc.User {
	return &mysvcgrpc.User{Id: u.ID, Name: u.Name}
}
