package main

import (
	"database/sql"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "github.com/mattn/go-sqlite3"

	"github.com/renatolm/go-grpc/internal/database"
	"github.com/renatolm/go-grpc/internal/service"
	"github.com/renatolm/go-grpc/internal/pb"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	manufacturerDB := database.NewManufacturer(db)
	manufacturerService := service.NewManufacturerService(*manufacturerDB)	

	grpcServer := grpc.NewServer()
	pb.RegisterManufacturerServiceServer(grpcServer, manufacturerService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}