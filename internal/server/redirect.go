package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net/http"
)

var urls = []string{"https://google.com"}

func (s *server) rotate(c *gin.Context) {
	ind := rand.Intn(len(urls))
	finalURL := urls[ind]
	c.Redirect(http.StatusFound, finalURL)
	log.Info().Msgf("click to %s", finalURL)
}
