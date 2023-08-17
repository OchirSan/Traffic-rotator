package server

func (s *server) setRoutes() {
	s.gin.GET("/serve", s.rotate)
}
