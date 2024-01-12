package utils

import (
	"github.com/Improwised/devempire-ridham-parmar-backend/models"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/response"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/structs"
)

// swagger:parameters RequestCreateUser
type RequestCreateUser struct {
	// in:body
	// required: true
	Body struct {
		structs.ReqRegisterUser
	}
}

// swagger:response ResponseCreateUser
type ResponseCreateUser struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   struct {
			models.User
		} `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestGetUser
type RequestGetUser struct {
	// in:path
	UserId string `json:"userId"`
}

// swagger:response ResponseGetUser
type ResponseGetUser struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   struct {
			models.User
		} `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestAuthnUser
type RequestAuthnUser struct {
	// in:body
	// required: true
	Body struct {
		structs.ReqLoginUser
	}
}

// swagger:response ResponseAuthnUser
type ResponseAuthnUser struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   struct {
			models.User
		} `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseCreateApp
type ResponseCreateApp struct {
	// in:body
	Body struct {
		// enum: success
		Status string          `json:"status"`
		Data   models.AppsInfo `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestCreateApp
type RequestCreateApp struct {
	// in:body
	// required: true
	Body struct {
		structs.ReqCreateApp
	} `json:"body"`
}

// swagger:response ResponseListApps
type ResponseListApps struct {
	// in:body
	Body struct {
		// enum: success
		Status string           `json:"status"`
		Data   response.ResData `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestListApps
type RequestListApps struct {
	// in:query
	Page     int    `json:"page"`
	Per_Page int    `json:"per_page"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Search   string `json:"search"`
}

// // swagger:response ResponseListCategory
type ResponseListCategory struct {
	// in:body
	Body struct {
		// enum: success
		Status string               `json:"status"`
		Data   response.ResCategory `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestDeleteApp
type RequestDeleteApp struct {
	// in:path
	AppId string `json:"appId"`
}

// // swagger:response ResponseDeleteApp
type ResponseDeleteApp struct {
	// in:body
	Body struct {
		// enum: success
		Status string          `json:"status"`
		Data   models.AppsInfo `json:"data"`
	} `json:"body"`
}

// // swagger:parameters RequestGetApp
type RequestGetApp struct {
	// in:path
	AppId string `json:"appId"`
}

// // swagger:response ResponseGetApp
type ResponseGetApp struct {
	// in:body
	Body struct {
		// enum: success
		Status string          `json:"status"`
		Data   models.AppsInfo `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestUpdateApp
type RequestUpdateApp struct {
	// in:path
	AppId string `json:"appId"`
	// in:body
	// required: true
	Body struct {
		structs.ReqCreateApp
	} `json:"body"`
}

// swagger:response ResponseUpdateApp
type ResponseUpdateApp struct {
	// in:body
	Body struct {
		// enum: success
		Status string          `json:"status"`
		Data   models.AppsInfo `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseCreateReview
type ResponseCreateReview struct {
	// in:body
	Body struct {
		// enum: success
		Status string            `json:"status"`
		Data   models.ReviewInfo `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestCreateReview
type RequestCreateReview struct {
	// in:path
	AppName string `json:"appName"`
	// in:body
	// required: true
	Body struct {
		structs.ReqCreateReviews
	} `json:"body"`
}

// swagger:response ResponseListReviews
type ResponseListReviews struct {
	// in:body
	Body struct {
		// enum: success
		Status string             `json:"status"`
		Data   response.ResReview `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestListReviews
type RequestListReviews struct {
	// in:path
	// required: true
	AppName string `json:"appName"`
	// in:query
	Page     int `json:"page"`
	Per_Page int `json:"per_page"`
}

////////////////////
// --- GENERIC ---//
////////////////////

// Response is okay
// swagger:response GenericResOk
type ResOK struct {
	// in:body
	Body struct {
		// enum:success
		Status string `json:"status"`
	}
}

// Fail due to user invalid input
// swagger:response GenericResFailBadRequest
type ResFailBadRequest struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Fail due to user invalid input
// swagger:response ResForbiddenRequest
type ResForbiddenRequest struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Server understand request but refuse to authorize it
// swagger:response GenericResFailConflict
type ResFailConflict struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Fail due to server understand request but unable to process
// swagger:response GenericResFailUnprocessableEntity
type ResFailUnprocessableEntity struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Fail due to resource not exists
// swagger:response GenericResFailNotFound
type ResFailNotFound struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Unexpected error occurred
// swagger:response GenericResError
type ResError struct {
	// in: body
	Body struct {
		// enum: error
		Status  string      `json:"status"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	} `json:"body"`
}
