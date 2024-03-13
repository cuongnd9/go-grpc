# go-grpc

Go 🤝 gRPC

## Preparation

```shell script
brew install protobuf
protoc --version
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
```

## How to use?

### Start MySQL

```shell script
docker run -it --name database -p 3306:3306 -e MYSQL_ROOT_PASSWORD=cuongnguyenpo -e MYSQL_DATABASE=cuongnguyenpo mysql:latest
```

*Create table if not exist*

```sql
CREATE TABLE `todo` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `title` varchar(200) DEFAULT NULL,
    `description` varchar(1024) DEFAULT NULL,
    `reminder` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE (`id`)
);
```

### Generate Go's protobuf code

```shell script
protoc --proto_path=proto \
    --go_out=pkg/pb --go_opt=paths=source_relative \
    --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
    proto/*.proto
```

### gRPC server

```shell script
cd cmd/server && go run main.go
```

### gRPC client

```shell script
cd cmd/client && go run main.go
```

## license

MIT © [Cuong Nguyen](https://github.com/cuongnd9/) 2024
