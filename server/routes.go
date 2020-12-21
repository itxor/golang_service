package server

func (s *Server) routes() {
	s.router.HandleFunc("/test", s.handleIndex())

	s.router.HandleFunc("/schedule", s.handleCallSchedule()).Methods("POST")
}
