package main

import (
	"context"
	"flag"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/go-kit/log"

	"github.com/go-kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"estore-backend/auth/account"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	port, result := os.LookupEnv("ACCOUNT_SERVER_PORT")
	if !result {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	httpAddrVal := fmt.Sprintf(":%s", port)
	var httpAddr = flag.String("http", httpAddrVal, "http listen address")

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	flag.Parse()
	ctx := context.Background()
	var srv account.Service
	{
		srv = account.NewService(logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := account.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
