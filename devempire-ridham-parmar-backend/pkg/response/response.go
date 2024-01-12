package response

import "github.com/Improwised/devempire-ridham-parmar-backend/models"

// All response sturcts
// Response struct have Res prefix
type ResPagination struct {
	TotalRecords int `json:"totalRecords,omitempty"`
	PerPage      int `json:"perPage,omitempty"`
	CurrentPage  int `json:"currentPage,omitempty"`
	TotalPages   int `json:"totalPages,omitempty"`
}
type ResData struct {
	Apps       []models.AppsInfo `json:"apps"`
	Pagination *ResPagination    `json:"pagination,omitempty"`
}

type ResCategory struct {
	Category []string `json:"category"`
}

type ResReview struct {
	Reviews    []models.ReviewInfo `json:"reviews"`
	Pagination *ResPagination      `json:"pagination,omitempty"`
}
