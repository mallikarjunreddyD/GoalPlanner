package rpc

import (
	"context"

	"github.com/mallikarjun/GoalPlanner/Back-end/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) SetGoal(ctx context.Context, req *pb.NewGoalRequest) (*pb.NewGoalResponse, error) {
	response := &pb.NewGoalResponse{}
	if req.GetTitle() == "" || req.GetPurpose1() == "" || req.GetPurpose2() == "" {
		return response, status.Error(codes.InvalidArgument, "missing required")
	}

	err := grpca.api.SetGoal(req.GetTitle(), req.GetPurpose1(), req.GetPurpose2())
	if err != nil {
		return response, status.Error(codes.Internal, "Unexpected error")
	}
	response = &pb.NewGoalResponse{Value: 2}
	return response, nil
}
