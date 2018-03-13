package endpoint

import (
	context "context"
	service "github.com/burxtx/fault/foo/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// FooRequest collects the request parameters for the Foo method.
type FooRequest struct {
	Where      string        `json:"where"`
	CondiBeans []interface{} `json:"condi_beans"`
}

// FooResponse collects the response parameters for the Foo method.
type FooResponse struct {
	M   []model.Fault `json:"m"`
	Err error         `json:"err"`
}

// MakeFooEndpoint returns an endpoint that invokes Foo on the service.
func MakeFooEndpoint(s service.FooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		m, err := s.Foo(ctx, req.Where, req.CondiBeans)
		return FooResponse{
			Err: err,
			M:   m,
		}, nil
	}
}

// Failed implements Failer.
func (r FooResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Foo implements Service. Primarily useful in a client.
func (e Endpoints) Foo(ctx context.Context, where string, CondiBeans ...interface{}) (m []model.Fault, err error) {
	request := FooRequest{
		CondiBeans: CondiBeans,
		Where:      where,
	}
	response, err := e.FooEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FooResponse).M, response.(FooResponse).Err
}
