package service

import (
	context "context"

	"github.com/burxtx/fault/fault/pkg/model"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(FaultService) FaultService

type loggingMiddleware struct {
	logger log.Logger
	next   FaultService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a FaultService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next FaultService) FaultService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Find(ctx context.Context, where string, condiBean ...interface{}) (m []model.Fault, err error) {
	defer func() {
		l.logger.Log("method", "Find", "m", m, "err", err)
	}()
	return l.next.Find(ctx, where, condiBean)
}
func (l loggingMiddleware) Get(ctx context.Context, id string) (m model.Fault, err error) {
	defer func() {
		l.logger.Log("method", "Get", "m", m, "err", err)
	}()
	return l.next.Get(ctx, id)
}
func (l loggingMiddleware) Add(ctx context.Context, fault model.Fault) (i int64, err error) {
	defer func() {
		l.logger.Log("method", "Add", "fault", fault.Hostname, "i", i, "err", err)
	}()
	return l.next.Add(ctx, fault)
}
