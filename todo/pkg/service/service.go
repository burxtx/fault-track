package service

import (
	"context"

	"github.com/burxtx/fault/todo/pkg/io"
	"github.com/go-xorm/xorm"
)

// TodoService describes the service.
type TodoService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Get(ctx context.Context) (t []io.Todo, error error)
	Add(ctx context.Context, todo io.Todo) (t io.Todo, error error)
	SetComplete(ctx context.Context, id string) (error error)
	RemoveComplete(ctx context.Context, id string) (error error)
	Delete(ctx context.Context, id string) (error error)
}

type basicTodoService struct{ db *xorm.Engine }

func (b *basicTodoService) Get(ctx context.Context) (t []io.Todo, error error) {
	// TODO implement the business logic of Get
	return t, error
}
func (b *basicTodoService) Add(ctx context.Context, todo io.Todo) (t io.Todo, error error) {
	// TODO implement the business logic of Add
	return t, error
}
func (b *basicTodoService) SetComplete(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of SetComplete
	return error
}
func (b *basicTodoService) RemoveComplete(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of RemoveComplete
	return error
}
func (b *basicTodoService) Delete(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of Delete
	return error
}

// NewBasicTodoService returns a naive, stateless implementation of TodoService.
func NewBasicTodoService(db *xorm.Engine) TodoService {
	return &basicTodoService{db}
}

// New returns a TodoService with all of the expected middleware wired in.
func New(db *xorm.Engine, middleware []Middleware) TodoService {
	var svc TodoService = NewBasicTodoService(db)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
