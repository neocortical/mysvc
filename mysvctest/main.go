package main

import (
	"fmt"
	"log"
	"os"

	"github.com/neocortical/mysvc"
	mysvccore "github.com/neocortical/mysvc/core"
	mysvcgrpc "github.com/neocortical/mysvc/grpc/client"
	"github.com/xiam/to"
)

func main() {
	var localService, grpcService mysvc.Service

	localService = mysvccore.NewService()
	grpcService, err := mysvcgrpc.NewGRPCService(os.Getenv("GRPC_ADDR"))
	if err != nil {
		log.Printf("error instantiating gRPC service: %v\n", err)
		os.Exit(1)
	}

	var ids []int64
	for _, arg := range os.Args[1:] {
		id := to.Int64(arg)
		if id <= 0 {
			log.Printf("invalid input: %s\n", arg)
			os.Exit(1)
		}
		ids = append(ids, id)
	}

	if len(ids) == 0 {
		fmt.Println("you must supply at least one ID")
		os.Exit(1)
	} else if len(ids) == 1 {
		localResult, localErr := localService.GetUser(ids[0])
		if localErr != nil {
			fmt.Printf("localService.GetUser() returned an error: %v\n", localErr)
		} else {
			fmt.Printf("localService.GetUser() returned: %+v\n", localResult)
		}
		grpcResult, remoteErr := grpcService.GetUser(ids[0])
		if remoteErr != nil {
			fmt.Printf("grpcService.GetUser() returned an error: %v\n", remoteErr)
		} else {
			fmt.Printf("grpcService.GetUser() returned: %+v\n", grpcResult)
		}
	} else {
		localResult, localErr := localService.GetUsers(ids)
		if localErr != nil {
			fmt.Printf("localService.GetUsers() returned an error: %v\n", localErr)
		} else {
			fmt.Printf("localService.GetUsers() returned: %+v\n", localResult)
		}
		grpcResult, remoteErr := grpcService.GetUsers(ids)
		if remoteErr != nil {
			fmt.Printf("grpcService.GetUsers() returned an error: %v\n", remoteErr)
		} else {
			fmt.Printf("grpcService.GetUsers() returned: %+v\n", grpcResult)
		}
	}
	os.Exit(0)
}
