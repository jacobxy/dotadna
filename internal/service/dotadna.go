package service

import (
	"context"

	pb "dotadna/api/dotadna"
)

type DotadnaService struct {
	pb.UnimplementedDotadnaServer
}

func NewDotadnaService() *DotadnaService {
	return &DotadnaService{}
}

func (s *DotadnaService) GetMatchList(ctx context.Context, req *pb.MatchListReq) (*pb.MatchListResp, error) {
	return &pb.MatchListResp{}, nil
}
func (s *DotadnaService) GetMatchDetial(ctx context.Context, req *pb.MatchDetialReq) (*pb.MatchDetialResp, error) {
	return &pb.MatchDetialResp{}, nil
}
