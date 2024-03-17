package pkg

import (
	"context"
	"github.com/cuongnd9/go-grpc/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RunServer() error {
	ctx := context.Background()

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: config.BuildDSN(),
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return RunGRPC(ctx, db, "50000")
}
