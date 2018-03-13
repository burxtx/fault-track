// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint "github.com/burxtx/fault/todo/pkg/endpoint"
	http1 "github.com/burxtx/fault/todo/pkg/http"
	service "github.com/burxtx/fault/todo/pkg/service"
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"Add":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Add", logger))},
		"Delete":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Delete", logger))},
		"Get":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Get", logger))},
		"RemoveComplete": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "RemoveComplete", logger))},
		"SetComplete":    {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SetComplete", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Get"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Get")), endpoint.InstrumentingMiddleware(duration.With("method", "Get"))}
	mw["Add"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Add")), endpoint.InstrumentingMiddleware(duration.With("method", "Add"))}
	mw["SetComplete"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SetComplete")), endpoint.InstrumentingMiddleware(duration.With("method", "SetComplete"))}
	mw["RemoveComplete"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "RemoveComplete")), endpoint.InstrumentingMiddleware(duration.With("method", "RemoveComplete"))}
	mw["Delete"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Delete")), endpoint.InstrumentingMiddleware(duration.With("method", "Delete"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Get", "Add", "SetComplete", "RemoveComplete", "Delete"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
