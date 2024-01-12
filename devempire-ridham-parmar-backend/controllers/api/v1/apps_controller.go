package v1

import (
	"fmt"
	"net/http"

	"github.com/Improwised/devempire-ridham-parmar-backend/constants"
	"github.com/Improwised/devempire-ridham-parmar-backend/models"
	pMetrics "github.com/Improwised/devempire-ridham-parmar-backend/pkg/prometheus"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/response"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/structs"
	"github.com/Improwised/devempire-ridham-parmar-backend/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

type AppsController struct {
	appModel *models.AppsModel
	logger   *zap.Logger
	pMetrics *pMetrics.PrometheusMetrics
}

func InitAppsController(db *goqu.Database, logger *zap.Logger, pMetrics *pMetrics.PrometheusMetrics) (*AppsController, error) {
	appsModel, err := models.InitAppsModel(db)
	if err != nil {
		return nil, err
	}
	return &AppsController{
		appModel: &appsModel,
		logger:   logger,
		pMetrics: pMetrics,
	}, nil
}

// ListAppDetail lists app details with pagination
// swagger:route GET /apps Apps RequestListApps
//
// List app details with pagination.
//
//			Schemes: http, https
//
//			Responses:
//			  200: ResponseListApps
//	       400: GenericResFailBadRequest
//		      500: GenericResError
func (ac *AppsController) ListAppDetail(c *fiber.Ctx) error {

	queries := c.Queries()
	validCondition := utils.ReturnValidCondition(queries)

	count, err := ac.appModel.CountApps(validCondition)
	if err != nil {
		ac.logger.Error("error while counting apps", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, err.Error())
	} else if count == 0 {
		return utils.JSONFail(c, http.StatusBadRequest, response.ResData{Apps: nil})
	}

	page := c.QueryInt(constants.QueryPage)
	perPage := c.QueryInt(constants.QueryPerPage)

	pagination, skip := utils.CalculatePagination(page, perPage, count)

	apps, err := ac.appModel.ListAppsDetails(validCondition, pagination.PerPage, skip)
	if err != nil {
		ac.logger.Error("error while getting app details", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrGetApps)
	}

	return utils.JSONSuccess(c, http.StatusOK, response.ResData{Apps: apps, Pagination: &pagination})
}

// ListCategory lists categories
// swagger:route GET /apps/category Category RequestListCategory
//
// List categories.
//
//		Schemes: http, https
//
//		Responses:
//		  200: ResponseListCategory
//	      500: GenericResError
func (ac *AppsController) ListCategory(c *fiber.Ctx) error {
	categories, err := ac.appModel.ListCategories()

	if err != nil {
		ac.logger.Error("error while getting category", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, constants.ErrGetCategories)
	}

	return utils.JSONSuccess(c, http.StatusOK, response.ResCategory{Category: categories})
}

// CreateApp insert an app
// swagger:route POST /apps Apps RequestCreateApp
//
// Insert an app.
//
//		Consumes:
//		- application/json
//
//		Schemes: http, https
//
//		Responses:
//		  201: ResponseCreateApp
//	    400: GenericResFailBadRequest
//	    500: GenericResError
func (ac *AppsController) CreateApp(c *fiber.Ctx) error {
	var appsReq structs.ReqCreateApp

	err := c.BodyParser(&appsReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(appsReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, utils.ValidatorErrorString(err))
	}

	app, err := ac.appModel.InsertApp(structs.ReqCreateApp{AppName: appsReq.AppName, Category: appsReq.Category, Size: appsReq.Size, Rating: appsReq.Rating, Reviews: appsReq.Reviews, Installs: appsReq.Installs, Types: appsReq.Types, Price: appsReq.Price, ContentRating: appsReq.ContentRating, Generes: appsReq.Generes, CurrentVersion: appsReq.CurrentVersion, AndroidVersion: appsReq.AndroidVersion})

	if err != nil {
		ac.logger.Error("error while insert app", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.JSONSuccess(c, http.StatusCreated, app)
}

// DeleteApp modifies deleted_at field from null to current time
// swagger:route DELETE /apps/{appId} Apps RequestDeleteApp
//
// Delete an app.
//
//		Schemes: http, https
//
//		Responses:
//		  200: ResponseDeleteApp
//	    500: GenericResError
func (ac *AppsController) DeleteApp(c *fiber.Ctx) error {
	id, err := c.ParamsInt("appId", 0)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}

	app, err := ac.appModel.DeleteAppDetail(id)
	if err != nil {
		ac.logger.Error("error while deleting app", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.JSONSuccess(c, http.StatusOK, app)
}

// GetApp gets details of apps
// swagger:route GET /apps/{appId} Apps RequestGetApp
//
// Get an app.
//
//			Schemes: http, https
//
//			Responses:
//			  200: ResponseGetApp
//	       404: GenericResFailNotFound
//		    500: GenericResError
func (ac *AppsController) GetAppDetail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("appId", 0)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}
	fmt.Println("from GetAppDetails ", id)

	app, err := ac.appModel.GetAppDetails(id)

	if err != nil {
		ac.logger.Error("error while getting app", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, err.Error())
	} else if app.Id == 0 {
		return utils.JSONFail(c, http.StatusNotFound, constants.AppNotExists)
	}
	return utils.JSONSuccess(c, http.StatusOK, app)
}

// UpdateApp update an app
// swagger:route PUT /apps/{appId} Apps RequestUpdateApp
//
// Update an app.
//
//			Consumes:
//			- application/json
//
//			Schemes: http, https
//
//			Responses:
//			  200: ResponseUpdateApp
//		    400: GenericResFailBadRequest
//	     404: GenericResFailNotFound
//		    500: GenericResError
func (ac *AppsController) UpdateApp(c *fiber.Ctx) error {
	var appsReq structs.ReqCreateApp

	err := c.BodyParser(&appsReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}

	id, err := c.ParamsInt("appId", 0)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}

	validFields := utils.ReturnValidFields(appsReq)
	if validFields == nil {
		return utils.JSONFail(c, http.StatusBadRequest, constants.ErrInValidFields)
	}

	app, err := ac.appModel.UpdateAppDetails(id, validFields)

	if app.Id == 0 {
		return utils.JSONFail(c, http.StatusNotFound, constants.AppNotExists)
	} else if err != nil {
		ac.logger.Error("error while updating app", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.JSONSuccess(c, http.StatusOK, app)
}
