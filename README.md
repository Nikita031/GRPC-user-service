# gRPC-user-service

This is a Golang-based gRPC service that provides functionalities for managing user details and includes a search capability. The service supports fetching user details by user ID, retrieving a list of user details by user IDs, and searching user details based on specific criteria.

## Features

- Fetch user details by user ID
- Retrieve a list of user details by user IDs
- Search user details based on specific criteria (e.g., fname, city, phone, height)
- Case-insensitive search for user first name

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.20 or later)
- [Protobuf Compiler](https://grpc.io/docs/protoc-installation/)
- [Docker](https://www.docker.com/products/docker-desktop/)

  After Installation of all these softwares please make sure to add path in environment variables configuration.

  ### Steps

1. Clone the repository:

    ```bash
    git clone https://github.com/Nikita031/grpc-user-service.git
    cd grpc-user-service
    ```

2. Install the required Go modules:

    ```bash
    go mod download
    ```
3. Generate the gRPC code from the proto file:

    ```bash
    protoc --go_out=. --go-grpc_out=. proto/user.proto
    ```
- If any error occurs like this 'protoc-gen-go' is not recognized as an internal or external command, operable program or batch file. Then firstly run these commands, after that run above mentioned command to generate the gRPC code:
  
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    $env:PATH += ";$(go env GOPATH)\bin"
    ```


## Docker

### Building the Docker Image

1. Build the Docker image:

    ```bash
    docker build -t grpc-user-service .
    ```
    
### Running the Docker Container

1. Run the Docker container:

    ```bash
    docker run -p 50051:50051 grpc-user-service
    ```

## Usage

1. Use a gRPC client to interact with the server. You can use tools like [Postman](https://www.postman.com/downloads/) or [BloomRPC](https://github.com/uw-labs/bloomrpc) for testing.

## API Methods

- `GetUser`: Fetch user details by user ID
- `GetUsers`: Retrieve a list of user details by user IDs
- `SearchUsers`: Search user details based on specific criteria


## Testing

- Run unit tests:

    ```bash
    go test ./server
    ```
