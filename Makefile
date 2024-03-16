.PHONY: mysql gen-proto run-server run-client

mysql:
	docker run -it --name database -p 3306:3306 -e MYSQL_ROOT_PASSWORD=cuongnguyenpo -e MYSQL_DATABASE=cuongnguyenpo mysql:latest

gen-proto:
	protoc --proto_path=proto \
	    --go_out=api --go_opt=paths=source_relative \
	    --go-grpc_out=api --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
	    proto/*.proto

run-server:
	go run cmd/server/main.go

run-client:
	go run cmd/client/main.go
