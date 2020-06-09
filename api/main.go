package api

import (
	"log"
	"net/http"

	"github.com/alexkaplun/cartesian/model"

	"github.com/go-chi/chi"
)

type Service struct {
	log       *log.Logger
	pointList *model.PointList
}

// creates a new Service instance with a logger and list of points
func New(log *log.Logger, pointList *model.PointList) *Service {
	return &Service{
		log:       log,
		pointList: pointList,
	}
}

// runs the http api server
func (s *Service) Run(port string) {
	s.log.Println("Starting...")
	r := s.prepareRouter()

	if err := http.ListenAndServe(port, r); err != nil {
		panic(err)
	}
}

// prepares the router which will serve queries
func (s *Service) prepareRouter() chi.Router {
	router := chi.NewRouter()
	router.Get("/api/points", s.pointsHandler)
	return router
}
