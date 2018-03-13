package service

import (
	"context"

	"github.com/burxtx/fault/foo/pkg/model"
)

// FooService describes the service.
type FooService interface {
	// Add your methods here
	Foo(ctx context.Context, where string, CondiBeans ...interface{}) (m []model.Fault, err error)
}

type basicFooService struct{}

func (b *basicFooService) Foo(ctx context.Context, where string, CondiBeans ...interface{}) (m []model.Fault, err error) {
	// TODO implement the business logic of Foo
	return m, err
}

// NewBasicFooService returns a naive, stateless implementation of FooService.
func NewBasicFooService() FooService {
	return &basicFooService{}
}

// New returns a FooService with all of the expected middleware wired in.
func New(middleware []Middleware) FooService {
	var svc FooService = NewBasicFooService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
