package module

import (
	"context"
	"errors"
	"github.com/cuongnd9/go-grpc/api"
	"github.com/cuongnd9/go-grpc/model"
	"github.com/cuongnd9/go-grpc/store"
)

type ToDoService struct {
	store *store.GlobalStore
}

func NewToDoService(store *store.GlobalStore) *ToDoService {
	return &ToDoService{store: store}
}

func (s *ToDoService) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	input := model.Todo{
		Title:       req.ToDo.Title,
		Description: req.ToDo.Description,
	}
	result := s.store.TodoStore.Create(&input)
	if result == nil {
		return nil, errors.New("can not create todo")
	}

	return &api.CreateResponse{
		Id: *result,
	}, nil
}

func (s *ToDoService) Update(ctx context.Context, req *api.UpdateRequest) (*api.UpdateResponse, error) {
	result := s.store.TodoStore.Update(req.ToDo.Id, req.ToDo.Title, req.ToDo.Description)

	return &api.UpdateResponse{
		Updated: result,
	}, nil
}

func (s *ToDoService) Delete(ctx context.Context, req *api.DeleteRequest) (*api.DeleteResponse, error) {
	result := s.store.TodoStore.Delete(req.Id)

	return &api.DeleteResponse{
		Deleted: result,
	}, nil
}

func (s *ToDoService) Read(ctx context.Context, req *api.ReadRequest) (*api.ReadResponse, error) {
	result := s.store.TodoStore.Read(req.Id)
	if result == nil {
		return nil, errors.New("not found todo item")
	}

	return &api.ReadResponse{
		ToDo: toTodoApi(result),
	}, nil
}

func (s *ToDoService) ReadAll(ctx context.Context, req *api.ReadAllRequest) (*api.ReadAllResponse, error) {
	result := s.store.TodoStore.ReadAll()

	return &api.ReadAllResponse{
		ToDos: toTodosApi(result),
	}, nil
}

func toTodoApi(todoModel *model.Todo) *api.ToDo {
	return &api.ToDo{
		Id:          todoModel.ID,
		Title:       todoModel.Title,
		Description: todoModel.Description,
	}
}

func toTodosApi(todoModels []model.Todo) []*api.ToDo {
	list := make([]*api.ToDo, len(todoModels))
	for i, v := range todoModels {
		list[i] = &api.ToDo{
			Id:          v.ID,
			Title:       v.Title,
			Description: v.Description,
		}
	}
	return list
}
