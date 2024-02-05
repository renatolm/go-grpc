package service

import (
	"context"

	"github.com/renatolm/go-grpc/internal/database"
	"github.com/renatolm/go-grpc/internal/pb"
)

type ManufacturerService struct {
	pb.UnimplementedManufacturerServiceServer
	ManufacturerDB database.Manufacturer
}

func NewManufacturerService (manufacturerDB database.Manufacturer) *ManufacturerService {
	return &ManufacturerService{
		ManufacturerDB: manufacturerDB,
	}
}

func (c *ManufacturerService) CreateManufacturer(ctx context.Context, inp *pb.CreateManufacturerRequest) (*pb.Manufacturer, error) {
	manufacturer, err := c.ManufacturerDB.Create(inp.Name, inp.Origincountry)
	if err != nil {
		return nil, err
	}

	manufacturerResponse := &pb.Manufacturer{
		Id:	manufacturer.ID,
		Name: manufacturer.Name,
		Origincountry: *manufacturer.OriginCountry,
	}

	// return &pb.ManufacturerResponse{
	// 	Manufacturer: manufacturerResponse,
	// }, nil
	return manufacturerResponse, nil
}