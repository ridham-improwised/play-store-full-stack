package v1

import (
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

type ReviewsController struct {
	reviewModel *models.ReviewsModel
	logger      *zap.Logger
	pMetrics    *pMetrics.PrometheusMetrics
}

func InitReviewsController(db *goqu.Database, logger *zap.Logger, pMetrics *pMetrics.PrometheusMetrics) (*ReviewsController, error) {
	reviewModel, err := models.InitReviewsModel(db)
	if err != nil {
		return nil, err
	}
	return &ReviewsController{
		reviewModel: &reviewModel,
		logger:      logger,
		pMetrics:    pMetrics,
	}, nil
}

// CreateReviews insert review
// swagger:route POST /app/{appName}/reviews Reviews RequestCreateReview
//
// Insert a review.
//
//		Consumes:
//		- application/json
//
//		Schemes: http, https
//
//		Responses:
//		  201: ResponseCreateReview
//	    400: GenericResFailBadRequest
//	    500: GenericResError
func (rc *ReviewsController) CreateReviews(c *fiber.Ctx) error {
	encodedAppName := c.Params("appName")

	decodedAppName, err := utils.DecodeString(encodedAppName)
	if err != nil {
		rc.logger.Error("error while decoding appname", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, err.Error())
	}

	var reviewsReq structs.ReqCreateReviews

	err = c.BodyParser(&reviewsReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(reviewsReq)
	if err != nil {
		return utils.JSONFail(c, http.StatusBadRequest, utils.ValidatorErrorString(err))
	}

	review, err := rc.reviewModel.InsertReviews(decodedAppName, reviewsReq)

	if err != nil {
		rc.logger.Error("error while inserting reviews", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.JSONSuccess(c, http.StatusCreated, review)
}

// ListReview lists reviews
// swagger:route GET /app/{appName}/reviews Reviews RequestListReviews
//
// List review details with pagination when query page and per_page is passed, otherwise displays all reviews associated with app.
//
//			Schemes: http, https
//
//			Responses:
//			  200: ResponseListReviews
//	       400: GenericResFailBadRequest
//		      500: GenericResError
func (rc *ReviewsController) ListReview(c *fiber.Ctx) error {
	var pagination response.ResPagination
	var skip int
	var reviews []models.ReviewInfo
	encodedAppName := c.Params("appName")

	decodedAppName, err := utils.DecodeString(encodedAppName)
	if err != nil {
		rc.logger.Error("error while decoding appname", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, err.Error())
	}

	page := c.QueryInt(constants.QueryPage)
	per_page := c.QueryInt(constants.QueryPerPage)

	count, err := rc.reviewModel.CountReviews(decodedAppName)
	if err != nil {
		rc.logger.Error("error while counting apps", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, err.Error())
	}

	if page == 0 && per_page == 0 {
		reviews, err = rc.reviewModel.ListReviews(decodedAppName, count, skip)
		pagination.TotalRecords = count
	} else {
		pagination, skip = utils.CalculatePagination(page, per_page, count)
		reviews, err = rc.reviewModel.ListReviews(decodedAppName, pagination.PerPage, skip)
	}

	if err != nil {
		rc.logger.Error("error while getting reviews details", zap.Error(err))
		return utils.JSONError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.JSONSuccess(c, http.StatusOK, response.ResReview{Reviews: reviews, Pagination: &pagination})
}
