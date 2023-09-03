package main

import (
	"github.com/OchirSan/Traffic-rotator/internal/server"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/rand"
	"time"
)

var (
	dbHost   = "host"
	dbPort   = 100
	user     = "user"
	password = "password"
	dbName   = "dbName"
)

func main() {
	// Set globals
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	rand.Seed(time.Now().UnixNano())

	log.Logger = log.With().Caller().Logger()
	gin.SetMode(gin.ReleaseMode)
	//dbConnectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s connect_timeout=5", dbHost, user, password, dbName, dbPort, false)
	//
	//sqldb, err := sqlx.Connect("postgres", dbConnectionString)
	//if err != nil {
	//	log.Fatal().Err(err).Msg("Database connection error")
	//}
	//
	//sqldb.SetMaxOpenConns(5)
	//sqldb.SetMaxIdleConns(5)

	//eventsCollector := collector.New()

	//db := database.New(sqldb, eventsCollector)

	serv := server.New()
	err := serv.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot run server")
	}
}
