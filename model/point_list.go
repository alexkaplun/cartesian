package model

import "sort"

type PointList struct {
	List []Point `json:"points"`
}

func (p *PointList) Points() []Point {
	return p.List
}

// returns sorted by distance list of points within 'dist' to 'from'
func (p *PointList) GetSortedWithinDistance(from Point, dist float64) *PointList {
	res := make([]PointWithDistance, 0)
	for _, point := range p.List {
		calculatedDistance := point.Distance(from)
		if calculatedDistance <= dist {
			res = append(res,
				PointWithDistance{
					Point:    point,
					distance: calculatedDistance,
				})
		}
	}

	// sort the results by distance
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].distance < res[j].distance
	})

	// copy the sorted results prepared for API
	list := make([]Point, len(res))
	for i, v := range res {
		list[i] = Point{v.X, v.Y}
	}

	return &PointList{
		List: list,
	}
}
