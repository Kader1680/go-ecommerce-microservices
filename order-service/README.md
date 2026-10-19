# Order Service

Simple Go microservice (standard library only, no frameworks) demonstrating
layered architecture: handler -> service -> repository.

## Structure

```
order-service/
├── cmd/server/main.go              # entrypoint, wiring, routing
├── internal/
│   ├── models/order.go             # domain model
│   ├── service/order_service.go    # business logic + repo interface
│   ├── repository/order_repository.go  # in-memory storage impl
│   └── handler/order_handler.go    # HTTP handlers
└── go.mod
```

## Run

```bash
go run cmd/server/main.go
```

Server starts on `:8080`.

## Try it

```bash
# Create an order
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"customer_id":"c1","product_id":"p1","quantity":2}'

# Get an order (replace the id with the one returned above)
curl http://localhost:8080/orders/ord_1234567890

# Cancel an order
curl -X POST http://localhost:8080/orders/ord_1234567890/cancel
```

## Notes

- `sync.RWMutex` protects the in-memory map since Go maps are not thread-safe
  and HTTP requests are handled concurrently.
- `OrderRepository` interface lives in the `service` package (Dependency
  Inversion) so storage can later be swapped for Postgres/Mongo without
  touching business logic, and so it's easy to unit test with a mock.
- Sentinel errors (`ErrInvalidQuantity`, `ErrOrderNotFound`) let handlers map
  errors to the right HTTP status code via `errors.Is`.

## Next steps (ask Claude if you want these built)

- Dockerfile
- Unit tests for the service layer (with a mock repository)
- Swap in-memory repo for Postgres
- Add a router/framework (chi, gin) once you're comfortable with stdlib
