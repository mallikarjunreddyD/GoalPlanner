package ports

import (
	"context"

	"github.com/mallikarjun/GoalPlanner/Back-end/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	SetGoal(ctx context.Context, req *pb.NewGoalRequest) (*pb.NewGoalResponse, error)
}
