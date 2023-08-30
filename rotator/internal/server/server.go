package server

import (
	"context"
	//"github.com/OchirSan/Traffic-rotator/internal/database"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type server struct {
	gin *gin.Engine
	//	collector eventsCollector
}

//type eventsCollector interface {
//	AsyncUpdate(event *database.Event)
//}

func New() *server {
	srv := &server{
		gin: gin.New(),
		//	collector: collector,
	}
	srv.gin.Use(gin.Recovery())
	srv.gin.Use(ginzerolog.Logger("gin"))
	srv.setRoutes()
	return srv
}

func (s *server) Run() error {
	log.Info().Msg("server started")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: s.gin,
	}
	s.listenForSignals(srv)
	return srv.ListenAndServe()
}

func (s *server) listenForSignals(srv *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func(quit chan os.Signal) {
		select {
		case <-quit:
			log.Debug().Msgf("Caught %s signal", os.Interrupt.String())
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			err := srv.Shutdown(ctx)
			if err != nil {
				log.Err(err).Msg("server shutdown problem")
			}
			log.Debug().Msg("server closed")
		}
	}(stop)
}
