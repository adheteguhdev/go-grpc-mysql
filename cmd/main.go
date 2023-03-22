package main

import (
	"log"
	"net"

	"github.com/adheteguhdev/go-grpc-mysql/cmd/configs"
	"github.com/adheteguhdev/go-grpc-mysql/cmd/services"
	pb "github.com/adheteguhdev/go-grpc-mysql/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("Failed to listen %v", err.Error())
	} else {
		log.Printf("Server started at %v", lis.Addr())
	}

	db := configs.ConnectDatabase();

	grpcServer := grpc.NewServer()

	productService := services.ProductService{DB: db}

	pb.RegisterProductServiceServer(grpcServer, &productService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}

}
