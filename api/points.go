package api

import (
	"net/http"
	"strconv"

	"github.com/alexkaplun/cartesian/model"

	"github.com/go-chi/render"
)

func (s *Service) pointsHandler(w http.ResponseWriter, r *http.Request) {
	// get and validate the 'x' param
	x, err := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	if err != nil {
		s.log.Printf("Error validating the 'x' value %v\n", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, http.StatusText(http.StatusBadRequest))
		return
	}

	// get and validate the 'y' param
	y, err := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
	if err != nil {
		s.log.Printf("Error validating the 'y' value %v\n", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, http.StatusText(http.StatusBadRequest))
		return
	}

	// get and validate the 'distance' param
	distance, err := strconv.ParseFloat(r.URL.Query().Get("distance"), 64)
	if err != nil {
		s.log.Printf("Error validating the 'distance' value %v\n", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, http.StatusText(http.StatusBadRequest))
		return
	}

	from := model.Point{
		X: x,
		Y: y,
	}
	pointsInRange := s.pointList.GetSortedWithinDistance(from, distance)
	s.log.Printf("Served a request, total %v points returned", len(pointsInRange.Points()))
	render.JSON(w, r, pointsInRange)
}
