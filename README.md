# go-grpc

Go 🤝 gRPC

# 🧑‍💻 Project Structure

### api/

This directory contains gRPC definitions.

### cmd/

The cmd/ directory contains the main applications of the project. In this case, there are client and server directories.

### config/

The config/ directory holds configuration files  such as environment or settings.

### model/

The model/ directory contains database entities.

### store/

The store/ directory contains sql logic for models. Eg: user store logics like queryProfile, editProfile. And it's related user model.

### module/

The module/ directory is collections of domain logic or features.

### pkg/

The pkg/ directory typically houses reusable packages or libraries that can be used across the project or potentially shared with other projects.

### proto/

This directory contains Protobuf files, which are used for defining gRPC services and messages.

### sql/

The sql/ directory contains SQL scripts, particularly release notes related to database schema changes or updates.

# 🐧 Commands

### 1. Local Machine

```shell script
brew install protobuf
protoc --version
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

#### .zshrc or.bashrc

```shell script
export PATH="$PATH:$(go env GOPATH)/bin"
```

### 2. Project dependencies

```shell script
go mod tidy
go mod vendor
```

### 3. MySQL

```shell script
docker run -it --name database -p 3306:3306 -e MYSQL_ROOT_PASSWORD=cuongpo -e MYSQL_DATABASE=cuongpo mysql:latest
```

### 4. Generate Go's protobuf code

```shell script
protoc --proto_path=proto \
    --go_out=api --go_opt=paths=source_relative \
    --go-grpc_out=api --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
    proto/*.proto
```

### 5. Run gRPC server

```shell script
cd cmd/server && go run main.go
```

### 6. Run gRPC client

```shell script
cd cmd/client && go run main.go
```

## License

MIT © [Cuong Nguyen](https://github.com/cuongnd9/) 2024
