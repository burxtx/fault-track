package endpoint

import (
	context "context"

	"github.com/burxtx/fault/fault/pkg/model"
	service "github.com/burxtx/fault/fault/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct {
	ID string `json:"id"`
}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	M   model.Fault `json:"m"`
	Err error       `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.FaultService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		m, err := s.Get(ctx, req.ID)
		return GetResponse{
			Err: err,
			M:   m,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Fault model.Fault `json:"fault"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	I   int64 `json:"i"`
	Err error `json:"err"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.FaultService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		i, err := s.Add(ctx, req.Fault)
		return AddResponse{
			Err: err,
			I:   i,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Err
}

// FindRequest collects the request parameters for the Find method.
type FindRequest struct {
	Where string        `json:"where"`
	Condi []interface{} `json:"condi"`
}

// FindResponse collects the response parameters for the Find method.
type FindResponse struct {
	M   []model.Fault `json:"m"`
	Err error         `json:"err"`
}

// MakeFindEndpoint returns an endpoint that invokes Find on the service.
func MakeFindEndpoint(s service.FaultService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindRequest)
		m, err := s.Find(ctx, req.Where, req.Condi)
		return FindResponse{
			Err: err,
			M:   m,
		}, nil
	}
}

// Failed implements Failer.
func (r FindResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context, id string) (m model.Fault, err error) {
	request := GetRequest{
		ID: id,
	}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).M, response.(GetResponse).Err
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, fault model.Fault) (i int64, err error) {
	request := AddRequest{Fault: fault}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).I, response.(AddResponse).Err
}

// Find implements Service. Primarily useful in a client.
func (e Endpoints) Find(ctx context.Context, where string, condi ...interface{}) (m []model.Fault, err error) {
	request := FindRequest{
		Where: where,
		Condi: condi,
	}
	response, err := e.FindEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FindResponse).M, response.(FindResponse).Err
}
