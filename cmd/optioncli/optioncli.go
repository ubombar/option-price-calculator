package main

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	pb "github.com/ubombar/option-price-calculator/pkg/option-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:8080", opts...)

	if err != nil {
		logrus.Error(err)
		return
	}

	defer conn.Close()

	client := pb.NewOptionServiceClient(conn)

	response, err := client.OptionPrice(context.TODO(), &pb.OptionPriceRequest{
		StockPrice:           100.0,
		ExercisePrice:        120.0,
		Volatility:           3.0,
		ExpirationTime:       0.5,  // 6 months
		RiskFreeInterestRate: 0.22, // about today's interest rate
	})

	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Printf("response.OptionPrice: %v\n", response.OptionPrice)
}
