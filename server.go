package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	ht "github.com/ogen-go/ogen/http"
	api "github.com/timtoronto634/sample-ogen/petstore"
)

func newServer() (http.Handler, error) {
	s, err := api.NewServer(TrivialHandler{})
	if err != nil {
		return nil, err
	}
	return s, nil
}

// TrivialHandler is no-op Handler which returns http.ErrNotImplemented.
type TrivialHandler struct{}

// MeGet implements GET /me operation.
//
// GET /me
func (TrivialHandler) MeGet(ctx context.Context) (r *api.User, _ error) {
	me, err := doSomething()
	if err != nil {
		return nil, err
	}

	return &api.User{ID: me.ID}, nil
}

// UsersPost implements POST /users operation.
//
// POST /users
func (TrivialHandler) UsersPost(ctx context.Context, req *api.User) error {
	return ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (TrivialHandler) NewError(ctx context.Context, err error) (r *api.ErrorStatusCode) {
	r = new(api.ErrorStatusCode)
	if strings.Contains(err.Error(), "[UserError]") {
		r = new(api.ErrorStatusCode)
		r.StatusCode = 400
		r.Response = api.Error{
			Code:    113355,
			Message: err.Error(),
		}
		return r
	}
	if strings.Contains(err.Error(), "[ServerError]") {
		fmt.Println("sending error to sentry...")
		// TODO: send error to sentry

		r = new(api.ErrorStatusCode)
		r.StatusCode = 500
		r.Response = api.Error{
			Code:    555000,
			Message: "please try again later",
		}
		return r
	}
	r.StatusCode = 500
	r.Response = api.Error{
		Code:    555000,
		Message: "unknown error",
	}
	return r
}

func doSomething() (*api.User, error) {
	return &api.User{ID: 123}, nil
}
