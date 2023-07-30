package optionserver

import (
	"context"

	pb "github.com/ubombar/option-price-calculator/pkg/option-service"
)

type OptionServer struct {
	pb.UnimplementedOptionServiceServer
}

func (s *OptionServer) OptionPrice(ctx context.Context, req *pb.OptionPriceRequest) (*pb.OptionPriceResponse, error) {
	return nil, nil
}
