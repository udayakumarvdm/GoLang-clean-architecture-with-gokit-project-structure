package main

import (
	cfg "GoLang_clean_architecture_with_gokit_project_structure/config"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	ep "GoLang_clean_architecture_with_gokit_project_structure/endpoint"
	repo "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/repository"
	uc "GoLang_clean_architecture_with_gokit_project_structure/our_project_service/usecase"
	s "GoLang_clean_architecture_with_gokit_project_structure/service"

	transport "GoLang_clean_architecture_with_gokit_project_structure/transport_http"

	logging "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()
	if config.GetBool(`app.debug`) {
		fmt.Println("Request Service RUN on DEBUG mode")
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {

	dbHost := config.GetString(`database.host`)
	dbPort := config.GetString(`database.port`)
	dbUser := config.GetString(`database.user`)
	dbPass := config.GetString(`database.pass`)
	dbName := config.GetString(`database.name`)

	httpServerAddr := config.GetString(`app.host`)

	db, err := gorm.Open("postgres",
		"host="+dbHost+" port="+dbPort+" user="+dbUser+" dbname="+dbName+" password="+dbPass+" sslmode=disable")
	if err != nil {
		fmt.Println("Database connection established failed ==>", err)
	} else {
		fmt.Println("Database connection established successfully")
	}

	logfile, err := os.OpenFile(config.GetString(`logfile`), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer logfile.Close()

	// Logging domain.
	var logger logging.Logger
	{
		w := logging.NewSyncWriter(logfile)
		logger = logging.NewLogfmtLogger(w)
		// AllowAll() AllowDebug() AllowInfo() AllowWarn() AllowError() AllowNone()
		// More info @ https://github.com/go-kit/kit/blob/master/log/level/level.go
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = logging.With(logger, "ts", logging.DefaultTimestampUTC)
		logger = logging.With(logger, "caller", logging.DefaultCaller)
	}

	flag.Parse()
	defer db.Close()

	ctx := context.Background()

	//Call repo
	repository := repo.NewRepoStruct(db, logger)

	//call usecase
	usecase := uc.NewUsecaseStruct(repository, logger)

	//Call service
	service := s.NewServiceStructFuction(usecase, logger)

	endpoints := ep.EndpointStruct{
		RequestFunctionEndpoint: ep.MakeEndpoint(service),
	}

	//Call transport layer httpServerAddr
	httpHandler := transport.MakeHttpHandler(ctx, endpoints, logger)

	errChan := make(chan error) //create error channel

	//HTTP transport
	go func() {
		fmt.Println("Go Http server at port: 8086")
		errChan <- http.ListenAndServe(httpServerAddr, httpHandler)
	}()
	//Given below go routine is used for controlling terminal presses Ex: CTRL + C
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errChan)
}
