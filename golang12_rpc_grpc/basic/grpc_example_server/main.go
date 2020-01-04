package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go-learn-project/golang12_rpc_grpc/basic/greetercenter"
	"go-learn-project/golang12_rpc_grpc/basic/grpc_example_server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "10.3.20.236:8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	eatService := service.EatService{}
	drinkService := service.DrinkService{}
	playService:=service.PlayService{}
	greetercenter.RegisterEatServer(s, &eatService)
	greetercenter.RegisterDrinkServer(s,&drinkService)
	greetercenter.RegisterPlayServer(s,&playService)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
