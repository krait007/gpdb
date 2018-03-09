package services

import (
	"golang.org/x/net/context"
	pb "gp_upgrade/idl"
)

func (s *HubClient) StatusConversion(ctx context.Context, in *pb.StatusConversionRequest) (*pb.StatusConversionReply, error) {
	return nil, nil
}
