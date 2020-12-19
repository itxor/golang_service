package server

func (s *Server) routes() {
	s.router.HandleFunc("/test", s.handleIndex())
}
