package main

import (
	"fmt"
)

type Server struct{}

func (s *Server) Listen(port int) {
	fmt.Printf("listening to port %v\n", port)
}

type Logging struct{}

func (s *Logging) Aggregate() {
	fmt.Println("aggregate distributed logs")
}

type Monitoring struct{}

func (s *Monitoring) Subscribe(event string) {
	fmt.Printf("continously watch %s\n", event)
}

type Rate struct{}

func (s *Rate) RateLimiting() {
	fmt.Println("rate limiting")
}

type Circuit struct{}

func (s *Rate) CircuitBreaker() {
	fmt.Println("circuit breaker")
}

type SystemFacade struct {
	webServer *Server
	log       *Logging
	monitor   *Monitoring
	rate      *Rate
	circuit   *Circuit
}

func NewSystemFacade() *SystemFacade {
	return &SystemFacade{
		&Server{},
		&Logging{},
		&Monitoring{},
		&Rate{},
		&Circuit{},
	}
}

func (s *SystemFacade) start() {
	s.rate.RateLimiting()
	s.rate.CircuitBreaker()
	s.webServer.Listen(80)
	s.log.Aggregate()
	s.monitor.Subscribe("network traffic")
	s.monitor.Subscribe("cpu resource")
	s.monitor.Subscribe("uakh")
}

func main() {
	system := NewSystemFacade()
	system.start()
}
