package server

func (s *server) setRoutes() {
	s.gin.GET("/4448f38b-cb6e-48be-aa25-fa4ee354e735", s.rotate)
}
