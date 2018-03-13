package service

import (
	context "context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(FooService) FooService

type loggingMiddleware struct {
	logger log.Logger
	next   FooService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a FooService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next FooService) FooService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Foo(ctx context.Context, where string, CondiBeans ...interface{}) (m []model.Fault, err error) {
	defer func() {
		l.logger.Log("method", "Foo", "where", where, "CondiBeans", CondiBeans, "m", m, "err", err)
	}()
	return l.next.Foo(ctx, where, CondiBeans)
}
