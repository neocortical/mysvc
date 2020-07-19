package main

import (
	"log"
	"net"
	"os"

	mysvccore "github.com/neocortical/mysvc/core"
	mysvcgrpc "github.com/neocortical/mysvc/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// configure our core service
	userService := mysvccore.NewService()

	// configure our gRPC service controller
	userServiceController := NewUserServiceController(userService)

	// start a gRPC server
	server := grpc.NewServer()
	mysvcgrpc.RegisterUserServiceServer(server, userServiceController)
	reflection.Register(server)

	con, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		panic(err)
	}

	log.Printf("Starting gRPC user service on %s...\n", con.Addr().String())
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
