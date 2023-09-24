package account

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Use(authMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeEmailReq,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteUser,
		decodeDeleteUserReq,
		encodeResponse,
	))

	r.Methods("PUT").Path("/user").Handler(httptransport.NewServer(
		endpoints.UpdateUser,
		decodeUpdateUserReq,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

type claims_key string

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			SECRETKEY, found_key := os.LookupEnv("SECRETKEY")
			if !found_key {
				fmt.Println("Missing secret key")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Server missing secret key"))
			} else {
				token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(SECRETKEY), nil
				})
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					var key claims_key = "props"
					ctx := context.WithValue(r.Context(), key, claims)
					// Access context values in handlers like this
					// props, _ := r.Context().Value("props").(jwt.MapClaims)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					fmt.Println(err)
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
				}
			}
		}
	})
}
