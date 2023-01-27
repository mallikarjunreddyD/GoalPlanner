package rpc

import (
	"log"
	"net"

	"github.com/mallikarjun/GoalPlanner/Back-end/internal/adapters/framework/left/grpc/pb"
	"github.com/mallikarjun/GoalPlanner/Back-end/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdpater(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	newGoalService := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterNewGoalServiceServer(grpcServer, newGoalService)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("Failed to serve gRPC server over port 9000: %v", err)
	}
}
