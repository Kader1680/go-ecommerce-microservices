package grpcserver

type ProductGRPCServer struct {
    pb.UnimplementedProductServiceServer   
    productService *service.ProductService  هو
}

func NewProductGRPCServer(s *service.ProductService) *ProductGRPCServer {
    return &ProductGRPCServer{productService: s}
}

func (g *ProductGRPCServer) ReserveStock(ctx context.Context, req *pb.ReserveStockRequest) (*pb.ReserveStockResponse, error) {
    err := g.productService.ReserveStock(req.ProductId, int(req.Quantity))
    if err != nil {
        return &pb.ReserveStockResponse{Success: false, Message: err.Error()}, nil
    }
    return &pb.ReserveStockResponse{Success: true}, nil
}