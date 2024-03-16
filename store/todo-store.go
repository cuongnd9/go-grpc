package store

import (
	"context"
	"database/sql"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cuongnd9/go-grpc/api"
)

type ToDoStore struct {
	db *sql.DB
}

func NewToDoStore(db *sql.DB) *ToDoStore {
	return &ToDoStore{db: db}
}

func (s *ToDoStore) connect(ctx context.Context) (*sql.Conn, error) {
	db, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "database connection failed")
	}
	return db, nil
}

func (s *ToDoStore) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {

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

	return &api.CreateResponse{
		Id: id,
	}, nil
}

func (s *ToDoStore) Read(ctx context.Context, req *api.ReadRequest) (*api.ReadResponse, error) {
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

	var todo api.ToDo
	var reminder time.Time
	if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
		return nil, status.Errorf(codes.Unknown, "scanning data from Todo failed: %v", err)
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, "founding multiple ToDo with ID")
	}

	return &api.ReadResponse{
		ToDo: &todo,
	}, nil
}

func (s *ToDoStore) Update(ctx context.Context, req *api.UpdateRequest) (*api.UpdateResponse, error) {
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

	return &api.UpdateResponse{
		Updated: rows,
	}, nil
}

func (s *ToDoStore) Delete(ctx context.Context, req *api.DeleteRequest) (*api.DeleteResponse, error) {
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

	return &api.DeleteResponse{
		Deleted: rows,
	}, nil
}

func (s *ToDoStore) ReadAll(ctx context.Context, req *api.ReadAllRequest) (*api.ReadAllResponse, error) {
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
	var list []*api.ToDo
	for rows.Next() {
		todo := new(api.ToDo)
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
			return nil, status.Errorf(codes.Unknown, "scanning data from Todo failed: %v", err)
		}
		list = append(list, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "selecting data failed")
	}

	return &api.ReadAllResponse{
		ToDos: list,
	}, nil
}
