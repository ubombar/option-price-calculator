package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	optionserver "github.com/ubombar/option-price-calculator/pkg/option-server"
	pb "github.com/ubombar/option-price-calculator/pkg/option-service"
	"google.golang.org/grpc"
)

func main() {
	logrus.Infoln("Server Started")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOptionServiceServer(grpcServer, &optionserver.OptionServer{})
	grpcServer.Serve(lis)

	logrus.Infoln("Server Stopped")
}
