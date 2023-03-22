package configs

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB{
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/go-grpc"))

	if err != nil{
		log.Fatalf("Database connection failed %v", err.Error())
	}else{
		log.Printf("Database connected!!!")
	}

	return db
}