package main

import (
	"fmt"
	"github.com/apoorvkrishna22/prometheus_local_testing/metrics"
	route2 "github.com/apoorvkrishna22/prometheus_local_testing/route"
	myuitls "github.com/apoorvkrishna22/prometheus_local_testing/uitls"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
	"time"
)

const (
	GetTicket    = "/api/getTicket"
	AddTicket    = "/api/addTicket"
	DeleteTicket = "/api/deleteTicket"
	UpdateTicket = "/api/updateTicket"
)

var routesEndpoints = []string{GetTicket, AddTicket, DeleteTicket, UpdateTicket}

type Server struct {
	routes chan string
	wg     sync.WaitGroup
}

func NewServer(routeBufferSize int) *Server {
	return &Server{
		routes: make(chan string, routeBufferSize),
	}
}

func (s *Server) AddRoute() {

	for {
		time.Sleep(time.Duration(myuitls.RandomInt(1, 5000)) * time.Millisecond)
		route := routesEndpoints[myuitls.RandomInt(1, 100)%4]
		fmt.Println("route requested ", route)
		s.routes <- route

	}

}

func (s *Server) ProcessRoutes() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for route := range s.routes {
			go s.handleRoute(route)
		}
	}()
}

func (s *Server) handleRoute(route string) {

	start := time.Now()

	var response string
	switch route {
	case GetTicket:
		response = route2.GetTicket()
	case AddTicket:
		response = route2.AddTicket()
	case DeleteTicket:
		response = route2.DelTicket()
	case UpdateTicket:
		response = route2.UpdateTicket()
	}
	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	fmt.Println("route = ", route, " returned response = ", response, " time taken = ", float64(duration))
	metrics.Metrics("care", "Tickets", route, response, duration)
}

func main() {

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		http.ListenAndServe(":9401", nil)
	}()

	server := NewServer(10000)

	go server.AddRoute()

	go server.ProcessRoutes()

	server.wg.Wait()

	c := make(chan string)

	<-c
}
