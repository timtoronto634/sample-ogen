// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

// handleMeGetRequest handles GET /me operation.
//
// GET /me
func (s *Server) handleMeGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/me"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MeGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err error
	)

	var response *User
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "MeGet",
			OperationSummary: "",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *User
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.MeGet(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.MeGet(ctx)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			recordError("Internal", err)
		}
		return
	}

	if err := encodeMeGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUsersPostRequest handles POST /users operation.
//
// POST /users
func (s *Server) handleUsersPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/users"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UsersPost",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "UsersPost",
			ID:   "",
		}
	)
	request, close, err := s.decodeUsersPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *UsersPostCreated
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "UsersPost",
			OperationSummary: "",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *User
			Params   = struct{}
			Response = *UsersPostCreated
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.UsersPost(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.UsersPost(ctx, request)
	}
	if err != nil {
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			recordError("Internal", err)
		}
		return
	}

	if err := encodeUsersPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}