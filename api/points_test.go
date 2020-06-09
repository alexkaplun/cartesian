package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-chi/chi"

	"github.com/stretchr/testify/require"

	"github.com/alexkaplun/cartesian/model"
)

var service *Service

// contains test cases for empty and malformed request values
func TestPointsHandler_Negative(t *testing.T) {
	// prepare router
	router := chi.NewRouter()
	router.Get("/api/points", service.pointsHandler)

	cases := map[string]struct {
		x                string
		y                string
		distance         string
		expectedCode     int
		expectedResponse *model.PointList
	}{
		"empty x": {
			x: "", y: "1", distance: "2",
			expectedCode:     http.StatusBadRequest,
			expectedResponse: nil,
		},
		"empty y": {
			x: "1", y: "", distance: "2",
			expectedCode:     http.StatusBadRequest,
			expectedResponse: nil,
		},
		"empty distance": {
			x: "1", y: "1", distance: "",
			expectedCode:     http.StatusBadRequest,
			expectedResponse: nil,
		},
		"bad x": {
			x: "abc", y: "1", distance: "2",
			expectedCode:     http.StatusBadRequest,
			expectedResponse: nil,
		},
		"bad y": {
			x: "0", y: "abc", distance: "2",
			expectedCode:     http.StatusBadRequest,
			expectedResponse: nil,
		},
		"bad distance": {
			x: "-1", y: "1", distance: "bad",
			expectedCode:     http.StatusBadRequest,
			expectedResponse: nil,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			queryPath := fmt.Sprintf("/api/points?x=%v&y=%v&distance=%v", tc.x, tc.y, tc.distance)
			makeReq := httptest.NewRequest(http.MethodGet, queryPath, nil)
			makeReq.URL.Query()

			out := httptest.NewRecorder()
			router.ServeHTTP(out, makeReq)
			require.Equal(t, tc.expectedCode, out.Code)
		})
	}
}

// contains regular testcases
func TestPointsHandler(t *testing.T) {
	// prepare router
	router := chi.NewRouter()
	router.Get("/api/points", service.pointsHandler)

	cases := map[string]struct {
		x                string
		y                string
		distance         string
		expectedCode     int
		expectedResponse *model.PointList
	}{
		"far away": {
			x: "50000", y: "50000", distance: "1",
			expectedCode:     http.StatusOK,
			expectedResponse: &model.PointList{[]model.Point{}},
		},
		"at a point": {
			x: "0", y: "0", distance: "0",
			expectedCode: http.StatusOK,
			expectedResponse: &model.PointList{
				List: []model.Point{
					{0, 0},
				},
			},
		},
		"duplicate at a point": {
			x: "1.01", y: "1.01", distance: "0.3",
			expectedCode: http.StatusOK,
			expectedResponse: &model.PointList{
				List: []model.Point{
					{1, 1},
					{1, 1},
				},
			},
		},
		"huge distance, all points": {
			x: "0", y: "0", distance: "100000",
			expectedCode: http.StatusOK,
			expectedResponse: &model.PointList{
				List: []model.Point{
					{0, 0}, {0.1, 0.1}, {1, 1}, {1, 1}, {-1, 2}, {5, 0},
					{0, -5}, {1000, -3000},
				},
			},
		},
		"limited distance, order matters": {
			x: "2", y: "3", distance: "5.1",
			expectedCode: http.StatusOK,
			expectedResponse: &model.PointList{
				List: []model.Point{
					{1, 1}, {1, 1}, {-1, 2}, {0.1, 0.1}, {0, 0},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			queryPath := fmt.Sprintf("/api/points?x=%v&y=%v&distance=%v", tc.x, tc.y, tc.distance)
			makeReq := httptest.NewRequest(http.MethodGet, queryPath, nil)

			out := httptest.NewRecorder()
			router.ServeHTTP(out, makeReq)
			require.Equal(t, tc.expectedCode, out.Code)

			response := &model.PointList{}
			err := json.Unmarshal(out.Body.Bytes(), response)
			require.NoError(t, err)

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}

// initializes the service with prepared data for tests
func init() {
	pointList := &model.PointList{
		List: []model.Point{
			{0, 0}, {1, 1}, {1, 1}, {0.1, 0.1}, {-1, 2},
			{1000, -3000}, {5, 0}, {0, -5},
		},
	}

	service = New(
		log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		pointList,
	)
}
