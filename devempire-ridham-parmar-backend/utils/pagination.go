package utils

import (
	"math"

	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/response"
)

// validate pagination
func CalculatePagination(page, per_page, count int) (response.ResPagination, int) {
	if per_page <= 0 {
		per_page = 10
	}
	if page <= 0 {
		page = 1
	}
	skip := (per_page * page) - per_page

	totalPages := int(math.Ceil(float64(count) / float64(per_page)))

	return response.ResPagination{
		TotalRecords: count,
		PerPage:      per_page,
		CurrentPage:  page,
		TotalPages:   totalPages,
	}, skip
}
