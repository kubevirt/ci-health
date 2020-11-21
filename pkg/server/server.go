package server

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct{}

func (s *Server) Start() {
	fmt.Println("Starting metrics server...")

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
