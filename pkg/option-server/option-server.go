package optionserver

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/ubombar/option-price-calculator/pkg/option-service"
)

type OptionServer struct {
	pb.UnimplementedOptionServiceServer
}

func (s *OptionServer) OptionPrice(ctx context.Context, req *pb.OptionPriceRequest) (*pb.OptionPriceResponse, error) {
	logrus.Infoln("Server received a request")

	var option_price float32 = 10.0

	rep := &pb.OptionPriceResponse{
		StockPrice:           req.StockPrice,
		ExercisePrice:        req.ExercisePrice,
		Volatility:           req.Volatility,
		ExpirationTime:       req.ExpirationTime,
		RiskFreeInterestRate: req.RiskFreeInterestRate,
		OptionPrice:          option_price,
	}

	return rep, nil
}
