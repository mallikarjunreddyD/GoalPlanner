package main

import (
	"log"

	"github.com/mallikarjun/GoalPlanner/Back-end/internal/adapters/app/api"
	"github.com/mallikarjun/GoalPlanner/Back-end/internal/adapters/core/goal"
	gRPC "github.com/mallikarjun/GoalPlanner/Back-end/internal/adapters/framework/left/grpc"
	"github.com/mallikarjun/GoalPlanner/Back-end/internal/adapters/framework/right/db"
	"github.com/mallikarjun/GoalPlanner/Back-end/internal/ports"
)

func main() {
	var err error

	//ports
	var dbaseAdapter ports.DbPort
	var core ports.GoalPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	//dbasePath := os.Getenv("DB_PATH")
	dbaseAdapter, err = db.NewAdpater("db")
	if err != nil {
		log.Fatalf("Failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = goal.NewAdpater()

	appAdapter = api.NewAdpater(dbaseAdapter, core)

	grpcAdapter = gRPC.NewAdpater(appAdapter)

	grpcAdapter.Run()

}
