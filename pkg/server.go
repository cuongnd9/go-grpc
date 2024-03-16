package pkg

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func RunServer() error {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:cuongnguyenpo@/cuongnguyenpo?parseTime=true")
	if err != nil {
		return fmt.Errorf("opening database failed")
	}
	defer db.Close()

	return RunGRPC(ctx, db, "50000")
}
