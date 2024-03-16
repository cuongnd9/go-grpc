package pkg

import (
	"context"
	"database/sql"
	api "github.com/cuongnd9/go-grpc/api"
	"github.com/cuongnd9/go-grpc/store"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func RunGRPC(ctx context.Context, db *sql.DB, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()

	todoService := store.NewToDoServiceServer(db)

	api.RegisterToDoServiceServer(server, todoService)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Printf("💅 server ready at 0.0.0.0:%s", port)
	return server.Serve(listen)
}
