package service

import (
	"context"
	"fmt"

	"github.com/burxtx/fault/fault/pkg/model"
	"github.com/go-xorm/xorm"
)

// FaultService describes the service.
type FaultService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Find(ctx context.Context, where string, condiBeans ...interface{}) ([]model.Fault, error)
	Get(ctx context.Context, id string) (model.Fault, error)
	Add(ctx context.Context, fault model.Fault) (int64, error)
}

type basicFaultService struct{ db *xorm.Engine }

func (b *basicFaultService) Find(ctx context.Context, where string, condiBean ...interface{}) ([]model.Fault, error) {
	// TODO implement the business logic of Find
	all := make([]model.Fault, 0)
	err := b.db.Find(&all)
	if err != nil {
		return all, err
	}
	return all, nil
}
func (b *basicFaultService) Get(ctx context.Context, id string) (model.Fault, error) {
	// TODO implement the business logic of Get
	f := new(model.Fault)
	has, err := b.db.Where("id=?", id).Get(f)
	if err != nil {
		return model.Fault{}, err
	}
	if has != true {
		return model.Fault{}, nil
	}
	return *f, nil
}
func (b *basicFaultService) Add(ctx context.Context, fault model.Fault) (int64, error) {
	// TODO implement the business logic of Add
	fmt.Printf("----%v",fault)
	affected, err := b.db.Insert(fault)
	fmt.Print(err)
	if err == nil {
		return affected, err
	}
	return affected, nil
}

// NewBasicFaultService returns a naive, stateless implementation of FaultService.
func NewBasicFaultService(db *xorm.Engine) FaultService {
	return &basicFaultService{db}
}

// New returns a FaultService with all of the expected middleware wired in.
func New(db *xorm.Engine, middleware []Middleware) FaultService {
	var svc FaultService = NewBasicFaultService(db)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
