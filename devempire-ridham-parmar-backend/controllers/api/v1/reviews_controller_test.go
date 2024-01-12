package v1_test

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/Improwised/devempire-ridham-parmar-backend/models"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/response"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/structs"
	"github.com/Improwised/devempire-ridham-parmar-backend/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateReview(t *testing.T) {
	t.Run("create review with valid body", func(t *testing.T) {
		var actual utils.ResponseCreateReview
		req := structs.ReqCreateReviews{
			TranslatedReview: "Great app",
			Sentiment:        "Positive",
		}
		expected := structs.ReqCreateReviews{
			TranslatedReview:      "Great app",
			Sentiment:             "Positive",
			SentimentPolarity:     "nan",
			SentimentSubjectivity: "nan",
		}
		res, err := client.
			R().
			EnableTrace().
			SetBody(req).
			SetResult(&actual.Body).
			Post("/api/v1/app/facebook/reviews")
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode())
		assert.Equal(t, expected.TranslatedReview, actual.Body.Data.TranslatedReview, "expected and actual translated review is not equal")
		assert.Equal(t, expected.Sentiment, actual.Body.Data.Sentiment, "expected and actual sentiment is not equal")
		assert.Equal(t, expected.SentimentPolarity, actual.Body.Data.SentimentPolarity, "expected and actual sentiment polarity is not equal")
		assert.Equal(t, expected.SentimentSubjectivity, actual.Body.Data.SentimentSubjectivity, "expected and actual sentiment subjectivity is not equal")
	})

	t.Run("create app with invalid fields", func(t *testing.T) {
		var actual utils.ResFailBadRequest

		reqReviews := []structs.ReqCreateReviews{
			{
				TranslatedReview: "",
				Sentiment:        "",
			},
			{
				TranslatedReview: "",
				Sentiment:        "Negative",
			},
			{
				TranslatedReview: "not so cool app",
				Sentiment:        "",
			},
		}

		expected := []string{
			"translatedreview,sentiment fields are invalid.",
			"translatedreview fields are invalid.",
			"sentiment fields are invalid.",
		}

		for index, val := range reqReviews {
			res, err := client.
				R().
				EnableTrace().
				SetBody(val).
				Post("/api/v1/app/instagram beta/reviews")

			assert.Nil(t, err)
			err = json.Unmarshal(res.Body(), &actual.Body)
			assert.Nil(t, err)
			assert.Equal(t, http.StatusBadRequest, res.StatusCode())
			assert.Equal(t, expected[index], actual.Body.Data)
		}
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from reviews")
		assert.Nil(t, err)
	})
}

func TestListReview(t *testing.T) {
	var actual utils.ResponseListReviews

	postData := []structs.ReqCreateReviews{
		{
			TranslatedReview: "good app for beginners",
			Sentiment:        "Positive",
		},
		{
			TranslatedReview: "better than any other apps",
			Sentiment:        "Positive",
		},
		{
			TranslatedReview: "good app",
			Sentiment:        "Positive",
		},
		{
			TranslatedReview: "best ever app",
			Sentiment:        "Positive",
		},
		{
			TranslatedReview: "worst app",
			Sentiment:        "Negative",
		},
		{
			TranslatedReview: "good app for beginners",
			Sentiment:        "Positive",
		},
		{
			TranslatedReview: "better than any other apps",
			Sentiment:        "Positive",
		},
		{
			TranslatedReview: "good app",
			Sentiment:        "Positive",
		},
		{
			TranslatedReview: "best ever app",
			Sentiment:        "Positive",
		},
		{
			TranslatedReview: "worst app",
			Sentiment:        "Negative",
		},
		{
			TranslatedReview: "good app",
			Sentiment:        "Positive",
		},
		{
			TranslatedReview: "good app",
			Sentiment:        "Positive",
		},
	}
	for _, val := range postData {
		_, err := client.R().SetBody(val).Post("/api/v1/app/instagram/reviews")
		if err != nil {
			log.Fatal("error while inserting reviews", zap.Error(err))
		}
	}

	expected := response.ResReview{
		Reviews: []models.ReviewInfo{
			{
				TranslatedReview:      "good app for beginners",
				Sentiment:             "Positive",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "better than any other apps",
				Sentiment:             "Positive",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "good app",
				Sentiment:             "Positive",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "best ever app",
				Sentiment:             "Positive",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "worst app",
				Sentiment:             "Negative",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "good app for beginners",
				Sentiment:             "Positive",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "better than any other apps",
				Sentiment:             "Positive",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "good app",
				Sentiment:             "Positive",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "best ever app",
				Sentiment:             "Positive",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "worst app",
				Sentiment:             "Negative",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "good app",
				Sentiment:             "Positive",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
			{
				TranslatedReview:      "good app",
				Sentiment:             "Positive",
				SentimentPolarity:     "nan",
				SentimentSubjectivity: "nan",
			},
		},
		Pagination: &response.ResPagination{
			TotalRecords: 12,
		},
	}

	t.Run("list reviews without page and per page query parameters", func(t *testing.T) {
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/app/instagram/reviews")

		assert.Nil(t, err)
		assert.Equal(t, expected.Pagination.TotalRecords, actual.Body.Data.Pagination.TotalRecords)
		assert.Equal(t, http.StatusOK, res.StatusCode())

		for i, val := range actual.Body.Data.Reviews {
			assert.Equal(t, expected.Reviews[i].TranslatedReview, val.TranslatedReview, "expected and actual translated review is not equal")
			assert.Equal(t, expected.Reviews[i].Sentiment, val.Sentiment, "expected and actual sentiment is not equal")
			assert.Equal(t, expected.Reviews[i].SentimentPolarity, val.SentimentPolarity, "expected and actual sentiment polarity is not equal")
			assert.Equal(t, expected.Reviews[i].SentimentSubjectivity, val.SentimentSubjectivity, "expected and actual sentiment subjectivity is not equal")
		}
	})

	t.Run("list reviews when app is not found", func(t *testing.T) {
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/app/instagr/reviews")

		assert.Nil(t, err)
		assert.Equal(t, expected.Pagination, actual.Body.Data.Pagination)
		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.Nil(t, actual.Body.Data.Reviews)
	})

	t.Run("list reviews with page and per_page query parameters", func(t *testing.T) {
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/app/instagram/reviews?page=1&per_page=5")

		assert.Nil(t, err)
		assert.Equal(t, 5, len(actual.Body.Data.Reviews))
		assert.Equal(t, http.StatusOK, res.StatusCode())
	})

	t.Run("calculate pagination when page and per page is non zero and app is passed", func(t *testing.T) {
		expectedPagination := &response.ResPagination{
			TotalRecords: 12,
			PerPage:      2,
			CurrentPage:  2,
			TotalPages:   6,
		}
		_, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/app/instagram/reviews?page=2&per_page=2")

		assert.Nil(t, err)
		assert.Equal(t, expectedPagination, actual.Body.Data.Pagination)
	})

	t.Run("calculate pagination when current page is greater than total pages", func(t *testing.T) {
		expectedPagination := &response.ResPagination{
			TotalRecords: 12,
			PerPage:      2,
			CurrentPage:  7,
			TotalPages:   6,
		}
		_, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/app/instagram/reviews?page=7&per_page=2")

		assert.Nil(t, err)
		assert.Equal(t, expectedPagination, actual.Body.Data.Pagination)
		assert.Nil(t, actual.Body.Data.Reviews)
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from reviews")
		assert.Nil(t, err)
	})
}
