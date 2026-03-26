# gRPC Order & Payment Project

This project demonstrates a gRPC-based microservice architecture with two services (`OrderService` and `PaymentService`) communicating with each other and clustering.

## Project Structure

- `proto/`: Protocol Buffer definitions and generated Go code.
- `main.go`: Combined implementation for both `Order` and `Payment` services, differentiated by environment variables.
- `client/`: A simple gRPC client to trigger the order flow.
- `docker-compose.yml`: Defines the service cluster, including 3 instances of `PaymentService`.

## Development Mandates

- **Proto:** Use `proto3` for definitions.
- **Service Communication:** `OrderService` must call `PaymentService` for payment processing.
- **Port:** Both services use `50051` as the internal gRPC port.
- **Load Balancing:** For demonstration, Docker Compose DNS is used for basic load balancing between `payment-service` replicas.

## How to Run & Test

1. **Start the services:**
   ```bash
   make up
   ```
   This will build and start 1 instance of `OrderService` and 3 instances of `PaymentService`.

2. **Run the test client:**
   ```bash
   make test-client
   ```
   This script will send an `OrderRequest` to the `OrderService`, which in turn will call one of the `PaymentService` instances.

3. **Verify clustering:**
   Run the test client multiple times and check the logs:
   ```bash
   make logs
   ```
   You should see different `PaymentService` nodes processing the requests.

4. **Clean up:**
   ```bash
   make down
   ```
