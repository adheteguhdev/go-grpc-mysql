package services

import (
	"context"

	pb "github.com/adheteguhdev/go-grpc-mysql/proto"
	"gorm.io/gorm"
)

type ProductService struct{
	pb.UnimplementedProductServiceServer
	DB *gorm.DB
}

func(p *ProductService) GetProducts(context.Context, *pb.Empty) (*pb.Products, error){
	
	products := &pb.Products{
		Pagination: &pb.Pagination{
			Total: 10,
			PerPage: 5,
			CurrentPage: 3,
			LastPage: 4,
		},
		Data: []*pb.Product{
			{
				Id: 1,
				Name: "Iphone 5",
				Price: 1000000,
				Stock: 5,
				Category: &pb.Category{
					Id: 1,
					Name: "Phone",
				},
			},
			{
				Id: 2,
				Name: "Samsung 23",
				Price: 120000,
				Stock: 3,
				Category: &pb.Category{
					Id: 1,
					Name: "Phone",
				},
			},
		},
	}

	return products, nil
}