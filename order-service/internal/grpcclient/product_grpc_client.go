package grpcclient

type ProductGRPCClient struct {
    client pb.ProductServiceClient  
}

func NewProductGRPCClient(conn *grpc.ClientConn) *ProductGRPCClient {
    return &ProductGRPCClient{client: pb.NewProductServiceClient(conn)}
}

func (c *ProductGRPCClient) ReserveStock(ctx context.Context, productID string, quantity int) error {
    resp, err := c.client.ReserveStock(ctx, &pb.ReserveStockRequest{
        ProductId: productID,
        Quantity:  int32(quantity),
    })
    if err != nil {
        return err
    }
    if !resp.Success {
        return errors.New(resp.Message)
    }
    return nil
}