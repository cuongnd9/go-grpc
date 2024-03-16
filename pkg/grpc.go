package pkg

import (
	"context"
	"database/sql"
	pb "github.com/cuongnd9/go-grpc/internal/pb"
	"github.com/cuongnd9/go-grpc/module/todo"
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

	todoService := todo.NewToDoServiceServer(db)

	pb.RegisterToDoServiceServer(server, todoService)

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
