package module

import (
	"context"
	"github.com/cuongnd9/go-grpc/api"
	"github.com/cuongnd9/go-grpc/store"
)

type ToDoService struct {
	store *store.GlobalStore
}

func NewToDoService(store *store.GlobalStore) *ToDoService {
	return &ToDoService{store: store}
}

func (s *ToDoService) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	newTodo, err := s.store.TodoStore.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.CreateResponse{
		Id: newTodo.Id,
	}, nil
}

func (s *ToDoService) Read(ctx context.Context, req *api.ReadRequest) (*api.ReadResponse, error) {
	readResp, err := s.store.TodoStore.Read(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.ReadResponse{
		ToDo: readResp.ToDo,
	}, nil
}

func (s *ToDoService) Update(ctx context.Context, req *api.UpdateRequest) (*api.UpdateResponse, error) {
	updateResp, err := s.store.TodoStore.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.UpdateResponse{
		Updated: updateResp.Updated,
	}, nil
}

func (s *ToDoService) Delete(ctx context.Context, req *api.DeleteRequest) (*api.DeleteResponse, error) {
	deleteResp, err := s.store.TodoStore.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.DeleteResponse{
		Deleted: deleteResp.Deleted,
	}, nil
}

func (s *ToDoService) ReadAll(ctx context.Context, req *api.ReadAllRequest) (*api.ReadAllResponse, error) {
	readAllResp, err := s.store.TodoStore.ReadAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.ReadAllResponse{
		ToDos: readAllResp.ToDos,
	}, nil
}
