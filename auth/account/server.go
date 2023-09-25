package account

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type ctxRequestKey struct{}

func putRequestInCtx(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, ctxRequestKey{}, r)
}

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	serverOptions := []httptransport.ServerOption{
		httptransport.ServerBefore(putRequestInCtx),
	}

	r.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		endpoints.RegisterUser,
		decRegisterReq,
		encodeResponse,
		serverOptions...,
	))

	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.LoginUser,
		decLoginReq,
		encodeResponse,
		serverOptions...,
	))

	r.Methods("POST").Path("/refresh").Handler(httptransport.NewServer(
		endpoints.RefreshUser,
		decRefreshReq,
		encodeResponse,
		serverOptions...,
	))

	r.Methods("DELETE").Path("/unregister/{id}").Handler(httptransport.NewServer(
		endpoints.UnregisterUser,
		decUnregisterReq,
		encodeResponse,
		serverOptions...,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Authorization", auth)
		next.ServeHTTP(w, r)
	})
}
