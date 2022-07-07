package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shikharvashistha/fampay/pkg/handlers/lifecycle"
	"github.com/shikharvashistha/fampay/pkg/store"
	"github.com/shikharvashistha/fampay/pkg/store/relational/models.go"
	"github.com/shikharvashistha/fampay/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	port        = "0.0.0.0:8087" //os.Getenv("PORT")
	usernameADB = "root"         //os.Getenv("ADB_USERNAME")
	passwordADB = "123456"       //os.Getenv("ADB_PASSWORD")
	dbNameADB   = "adb"          //os.Getenv("ADB_DBNAME")
	hostADB     = "0.0.0.0"      //os.Getenv("ADB_HOST")
	portADB     = "3306"         //os.Getenv("ADB_PORT")

	//dsnADB is the data source name for the ADB database
	dsnADB = usernameADB + ":" + passwordADB + "@tcp(" + hostADB + ":" + portADB + ")/" + dbNameADB + "?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {

	//Initialize a new logger
	logger := utils.NewLogger("main")

	logger.Info("Attempting to open connection to ADB...")

	//Open a connection to the ADB database
	db, err := connect(dsnADB)
	if err != nil {
		logger.Info("Exiting")
		os.Exit(1)
	}

	logger.Info("Connection to ADB opened successfully")

	logger.Info("Attempting to register schemas...")
	//Register the schemas
	err = models.RegisterSchema(db)
	if err != nil {
		logger.WithError(utils.ADB, err).Fatal("Failed to register schemas")
	}
	logger.Info("Schemas registered successfully")
	//Initialize a new gin router
	r := gin.Default()
	v1 := r.Group("/v1", gin.Logger())

	store := store.NewStore(db)
	deploymentSVC := lifecycle.NewDeploymentSvc(store, logger)

	lifecycle.RegisterHTTPHandlers(v1, deploymentSVC)
	s := &http.Server{
		Addr:         port,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
		Handler:      r,
	}

	//Start the server
	serveError := s.ListenAndServe()
	if serveError != nil {
		fmt.Println(serveError)
	}
}
func connect(dsn string) (*gorm.DB, error) {
	//Open a connection to the database
	logger := utils.NewLogger("connect")
	var (
		db      *gorm.DB
		err     error
		retries int = 10
	)
	//Retry connecting to the database
	for retries > 0 {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		} else {
			retries--
			logger.WithError(utils.ADB, err).Error("Failed to open connection to ADB")
			time.Sleep(10 * time.Second)
		}
	}
	if err != nil {
		logger.WithError(utils.ADB, err).Error("Failed to open connection to ADB")
		return nil, err
	}
	//Set the connection to auto migrate
	sqlDB, err := db.DB()
	if err != nil {
		logger.WithError(utils.ADB, err).Error("Failed to open connection to ADB")
		return nil, err
	}
	//Set the connection Settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, err
}
