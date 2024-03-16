package todo

import (
	"context"
	"database/sql"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/cuongnd9/go-grpc/pkg/pb"
)

type ToDoServiceServer struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) *ToDoServiceServer {
	return &ToDoServiceServer{db: db}
}

func (s *ToDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	db, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "database connection failed")
	}
	return db, nil
}

func (s *ToDoServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	// database connection
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	// closing connection
	defer db.Close()

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format")
	}

	// inserting into ToDo
	query := "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?,?,?)"
	result, err := db.ExecContext(ctx, query,
		req.ToDo.Title, req.ToDo.Description)
	if err != nil {
		return nil, status.Error(codes.Unknown, "inserting into Todo failed")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "getting last insert id failed")
	}

	return &pb.CreateResponse{
		Id: id,
	}, nil
}

func (s *ToDoServiceServer) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	// database connection
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	// closing connection
	defer db.Close()

	// query ToDo by ID
	query := "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?"
	rows, err := db.QueryContext(ctx, query, req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "getting element from ToDo failed")
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "retrieving data from Todo failed")
		}
		return nil, status.Error(codes.NotFound, "not found ToDo with ID")
	}

	var todo pb.ToDo
	var reminder time.Time
	if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
		return nil, status.Errorf(codes.Unknown, "scanning data from Todo failed: %v", err)
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, "founding multiple ToDo with ID")
	}

	return &pb.ReadResponse{
		ToDo: &todo,
	}, nil
}

func (s *ToDoServiceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// database connection
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	// closing connection
	defer db.Close()

	query := "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`=? WHERE `ID`=?"
	result, err := db.ExecContext(ctx, query,
		req.ToDo.Title, req.ToDo.Description, req.ToDo.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updating ToDo failed")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "retrieving data from rows affected failed")
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, "ToDo with ID not found")
	}

	return &pb.UpdateResponse{
		Updated: rows,
	}, nil
}

func (s *ToDoServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// database connection
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	// closing connection
	defer db.Close()

	query := "DELETE FROM ToDo WHERE `ID`=?"
	result, err := db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "deleting ToDo with ID failed")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "retrieving data from rows affected failed")
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, "ToDo with ID not found")
	}

	return &pb.DeleteResponse{
		Deleted: rows,
	}, nil
}

func (s *ToDoServiceServer) ReadAll(ctx context.Context, req *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	// database connection
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	// closing connection
	defer db.Close()

	query := "SELECT `ID`, `Title`, `Description` FROM ToDo"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, status.Error(codes.Unknown, "selecting ToDos failed")
	}
	defer rows.Close()

	var reminder time.Time
	var list []*pb.ToDo
	for rows.Next() {
		todo := new(pb.ToDo)
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
			return nil, status.Errorf(codes.Unknown, "scanning data from Todo failed: %v", err)
		}
		list = append(list, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "selecting data failed")
	}

	return &pb.ReadAllResponse{
		ToDos: list,
	}, nil
}
