package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/go-kit/log"

	"github.com/go-kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"estore-backend/account"

	"github.com/joho/godotenv"
)

func getVar(varName string) (string, bool) {
	host, result := os.LookupEnv(varName)
	return host, result
}

func getDBConn() (string, bool) {
	host, result := getVar("DB_HOST")
	if !result {
		return "", false
	}

	port, result := getVar("DB_PORT")
	if !result {
		return "", false
	}
	iport, err := strconv.Atoi(port)
	if err != nil {
		return "", false
	}

	user, result := getVar("DB_USER")
	if !result {
		return "", false
	}

	password, result := getVar("DB_PASSWORD")
	if !result {
		return "", false
	}

	dbname, result := getVar("DB_NAME")
	if !result {
		return "", false
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, iport, user, password, dbname)

	return psqlInfo, true
}

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
	port, result := getVar("DB_PORT")
	if !result {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	httpAddrVal := fmt.Sprintf(":%s", port)
	var httpAddr = flag.String("http", httpAddrVal, "http listen address")

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		// 	"password=%s dbname=%s sslmode=disable",
		// 	host, port, user, password, dbname)
		psqlInfo, result := getDBConn()
		if !result {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		println(psqlInfo)

		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			db.Close()
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	flag.Parse()
	ctx := context.Background()
	var srv account.Service
	{
		repository := account.NewRepo(db, logger)

		srv = account.NewService(repository, logger)
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
