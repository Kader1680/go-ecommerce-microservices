package grpcserver

import (
	"context"

	"github.com/yourname/product-service/internal/service"
	pb "github.com/yourname/product-service/pb"
)

type ProductGRPCServer struct {
	pb.UnimplementedProductServiceServer
	productService *service.ProductService
}

func NewProductGRPCServer(s *service.ProductService) *ProductGRPCServer {
	return &ProductGRPCServer{productService: s}
}

func (g *ProductGRPCServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	p, err := g.productService.GetProduct(req.ProductId)
	if err != nil {
		return nil, err
	}
	return &pb.ProductResponse{
		Id:    p.ID,
		Name:  p.Name,
		Price: p.Price,
		Stock: int32(p.Stock),
	}, nil
}

func (g *ProductGRPCServer) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := g.productService.ListProducts()
	if err != nil {
		return nil, err
	}

	resp := &pb.ListProductsResponse{}
	for _, p := range products {
		resp.Products = append(resp.Products, &pb.ProductResponse{
			Id:    p.ID,
			Name:  p.Name,
			Price: p.Price,
			Stock: int32(p.Stock),
		})
	}
	return resp, nil
}

func (g *ProductGRPCServer) ReserveStock(ctx context.Context, req *pb.ReserveStockRequest) (*pb.ReserveStockResponse, error) {
	err := g.productService.ReserveStock(req.ProductId, int(req.Quantity))
	if err != nil {
		return &pb.ReserveStockResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.ReserveStockResponse{Success: true}, nil
}