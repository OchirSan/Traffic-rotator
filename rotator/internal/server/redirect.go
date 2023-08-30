package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net/http"
)

var urls = []string{"https://landing-rotator.wpu.sh/serve?source_id=10"}

func (s *server) rotate(c *gin.Context) {
	ind := rand.Intn(len(urls))
	finalURL := urls[ind]
	//event := &database.Event{EventTime: time.Now(), ClickURL: finalURL}
	//s.collector.AsyncUpdate(event)
	c.Redirect(http.StatusFound, finalURL)
	log.Info().Msgf("click to %s", finalURL)
}
