package v1_test

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/Improwised/devempire-ridham-parmar-backend/models"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/response"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/structs"
	"github.com/Improwised/devempire-ridham-parmar-backend/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateApp(t *testing.T) {
	t.Run("create app with valid body", func(t *testing.T) {
		var actual utils.ResponseCreateApp
		req := structs.ReqCreateApp{
			AppName:        "Instagram",
			Category:       "social",
			Size:           "20M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "teen",
			Generes:        "social media",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		}
		expected := structs.ReqCreateApp{
			AppName:        "Instagram",
			Category:       "social",
			Size:           "20M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "teen",
			Generes:        "social media",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		}
		res, err := client.
			R().
			EnableTrace().
			SetBody(req).
			SetResult(&actual.Body).
			Post("/api/v1/apps")
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode())
		assert.Equal(t, expected.AppName, actual.Body.Data.AppName, "expected and actual appname is not equal")
		assert.Equal(t, expected.Category, actual.Body.Data.Category, "expected and actual category is not equal")
		assert.Equal(t, expected.Size, actual.Body.Data.Size, "expected and actual size is not equal")
		assert.Equal(t, expected.Rating, actual.Body.Data.Rating, "expected and actual rating is not equal")
		assert.Equal(t, expected.Reviews, actual.Body.Data.Reviews, "expected and actual reviews is not equal")
		assert.Equal(t, expected.Installs, actual.Body.Data.Installs, "expected and actual installs is not equal")
		assert.Equal(t, expected.Price, actual.Body.Data.Price, "expected and actual price is not equal")
		assert.Equal(t, expected.ContentRating, actual.Body.Data.ContentRating, "expected and content rating appname is not equal")
		assert.Equal(t, expected.Generes, actual.Body.Data.Generes, "expected and actual genres is not equal")
		assert.Equal(t, expected.CurrentVersion, actual.Body.Data.CurrentVersion, "expected and actual current version is not equal")
		assert.Equal(t, expected.AndroidVersion, actual.Body.Data.AndroidVersion, "expected and actual android version is not equal")
		assert.Equal(t, expected.Types, actual.Body.Data.Types, "expected and actual types is not equal")
	})

	t.Run("create app with invalid fields", func(t *testing.T) {
		var actual utils.ResFailBadRequest

		reqApps := []structs.ReqCreateApp{
			{
				AppName:        "",
				Category:       "",
				Size:           "",
				Rating:         "",
				Reviews:        "",
				Installs:       "",
				Price:          "",
				ContentRating:  "",
				Generes:        "",
				CurrentVersion: "",
				AndroidVersion: "",
				Types:          "",
			},
			{
				AppName:        "",
				Category:       "social",
				Size:           "20M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "",
				Size:           "20M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "social",
				Size:           "",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "social",
				Size:           "50M",
				Rating:         "",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "social",
				Size:           "50M",
				Rating:         "0",
				Reviews:        "",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "social",
				Size:           "50M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "social",
				Size:           "50M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "social",
				Size:           "50M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "social",
				Size:           "50M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "social",
				Size:           "50M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "social",
				Size:           "50M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "",
				Types:          "free",
			},
			{
				AppName:        "instagram",
				Category:       "social",
				Size:           "50M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "5.0.0",
				AndroidVersion: "5.1",
				Types:          "",
			},
		}

		expected := []string{
			"appname,category,size,rating,reviews,types,installs,price,contentrating,generes,currentversion,androidversion fields are invalid.",
			"appname fields are invalid.",
			"category fields are invalid.",
			"size fields are invalid.",
			"rating fields are invalid.",
			"reviews fields are invalid.",
			"installs fields are invalid.",
			"price fields are invalid.",
			"contentrating fields are invalid.",
			"generes fields are invalid.",
			"currentversion fields are invalid.",
			"androidversion fields are invalid.",
			"types fields are invalid.",
		}

		for index, val := range reqApps {
			res, err := client.
				R().
				EnableTrace().
				SetBody(val).
				Post("/api/v1/apps")

			assert.Nil(t, err)
			err = json.Unmarshal(res.Body(), &actual.Body)
			assert.Nil(t, err)
			assert.Equal(t, http.StatusBadRequest, res.StatusCode())
			assert.Equal(t, expected[index], actual.Body.Data)
		}
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from apps")
		assert.Nil(t, err)
	})
}

func TestListApp(t *testing.T) {
	var actual utils.ResponseListApps

	postData := []structs.ReqCreateApp{
		{
			AppName:        "Instagram",
			Category:       "social",
			Size:           "20M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "teen",
			Generes:        "social media",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Google Chrome",
			Category:       "social",
			Size:           "250M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "social media",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "paid",
		},
		{
			AppName:        "Snapchats",
			Category:       "social",
			Size:           "250M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "$2",
			ContentRating:  "everyone",
			Generes:        "social media",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Facebookd",
			Category:       "social",
			Size:           "250M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "social media",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "paid",
		},
		{
			AppName:        "Whatsappss",
			Category:       "social",
			Size:           "250M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "social media",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Photo Studio",
			Category:       "Photography",
			Size:           "10M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "social media",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "paid",
		},
		{
			AppName:        "Skype",
			Category:       "communication",
			Size:           "20M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "communication",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Linkedin",
			Category:       "learning",
			Size:           "25M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "learning",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "paid",
		},
		{
			AppName:        "Hotstar",
			Category:       "movies",
			Size:           "30M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "movies",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Netflix",
			Category:       "movies",
			Size:           "40M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "movies",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "paid",
		},
		{
			AppName:        "Mattermost",
			Category:       "communication",
			Size:           "30M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "communication",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Telegram",
			Category:       "social media",
			Size:           "250M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "social media",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "paid",
		},
	}
	for _, val := range postData {
		_, err := client.R().SetBody(val).Post("/api/v1/apps")
		if err != nil {
			log.Fatal("error while inserting apps", zap.Error(err))
		}
	}

	expected := response.ResData{
		Apps: []models.AppsInfo{
			{
				AppName:        "Instagram",
				Category:       "social",
				Size:           "20M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "teen",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "Google Chrome",
				Category:       "social",
				Size:           "250M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "everyone",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "paid",
			},
			{
				AppName:        "Snapchats",
				Category:       "social",
				Size:           "250M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "$2",
				ContentRating:  "everyone",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "Facebookd",
				Category:       "social",
				Size:           "250M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "everyone",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "paid",
			},
			{
				AppName:        "Whatsappss",
				Category:       "social",
				Size:           "250M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "everyone",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "Photo Studio",
				Category:       "Photography",
				Size:           "10M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "everyone",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "paid",
			},
			{
				AppName:        "Skype",
				Category:       "communication",
				Size:           "20M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "everyone",
				Generes:        "communication",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "Linkedin",
				Category:       "learning",
				Size:           "25M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "everyone",
				Generes:        "learning",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "paid",
			},
			{
				AppName:        "Hotstar",
				Category:       "movies",
				Size:           "30M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "everyone",
				Generes:        "movies",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "Netflix",
				Category:       "movies",
				Size:           "40M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "everyone",
				Generes:        "movies",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "paid",
			},
			{
				AppName:        "Mattermost",
				Category:       "communication",
				Size:           "30M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "everyone",
				Generes:        "communication",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "free",
			},
			{
				AppName:        "Telegram",
				Category:       "social media",
				Size:           "250M",
				Rating:         "0",
				Reviews:        "0",
				Installs:       "0",
				Price:          "0",
				ContentRating:  "everyone",
				Generes:        "social media",
				CurrentVersion: "2.0.0",
				AndroidVersion: "5.1",
				Types:          "paid",
			},
		},
		Pagination: &response.ResPagination{
			TotalRecords: 12,
			PerPage:      10,
			CurrentPage:  1,
			TotalPages:   2,
		},
	}

	t.Run("list apps without query parameters", func(t *testing.T) {
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/apps")

		assert.Nil(t, err)
		assert.Equal(t, expected.Pagination, actual.Body.Data.Pagination)
		assert.Equal(t, 10, len(actual.Body.Data.Apps))
		assert.Equal(t, http.StatusOK, res.StatusCode())

		for i, val := range actual.Body.Data.Apps {
			assert.Equal(t, expected.Apps[i].AppName, val.AppName, "expected and actual appname is not equal")
			assert.Equal(t, expected.Apps[i].Category, val.Category, "expected and actual category is not equal")
			assert.Equal(t, expected.Apps[i].Size, val.Size, "expected and actual size is not equal")
			assert.Equal(t, expected.Apps[i].Rating, val.Rating, "expected and actual rating is not equal")
			assert.Equal(t, expected.Apps[i].Reviews, val.Reviews, "expected and actual reviews is not equal")
			assert.Equal(t, expected.Apps[i].Installs, val.Installs, "expected and actual installs is not equal")
			assert.Equal(t, expected.Apps[i].Price, val.Price, "expected and actual price is not equal")
			assert.Equal(t, expected.Apps[i].ContentRating, val.ContentRating, "expected and content rating appname is not equal")
			assert.Equal(t, expected.Apps[i].Generes, val.Generes, "expected and actual genres is not equal")
			assert.Equal(t, expected.Apps[i].CurrentVersion, val.CurrentVersion, "expected and actual current version is not equal")
			assert.Equal(t, expected.Apps[i].AndroidVersion, val.AndroidVersion, "expected and actual android version is not equal")
			assert.Equal(t, expected.Apps[i].Types, val.Types, "expected and actual types is not equal")
		}
	})

	t.Run("list apps with page query parameters", func(t *testing.T) {
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/apps?page=2")

		assert.Nil(t, err)
		assert.NotEqual(t, expected.Pagination.CurrentPage, actual.Body.Data.Pagination.CurrentPage)
		assert.Equal(t, http.StatusOK, res.StatusCode())
	})

	t.Run("list apps with page and per_page query parameters", func(t *testing.T) {
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/apps?page=1&per_page=5")

		assert.Nil(t, err)
		assert.Equal(t, 5, len(actual.Body.Data.Apps))
		assert.Equal(t, http.StatusOK, res.StatusCode())
	})

	t.Run("calculate pagination when page and per page is zero", func(t *testing.T) {
		expectedPagination := response.ResPagination{
			TotalRecords: 0,
			PerPage:      10,
			CurrentPage:  1,
			TotalPages:   0,
		}
		expectedSkip := 0
		actualPagination, actualSkip := utils.CalculatePagination(0, 0, 0)
		assert.Equal(t, expectedPagination, actualPagination)
		assert.Equal(t, expectedSkip, actualSkip)
	})

	t.Run("calculate pagination when page and per page is non zero", func(t *testing.T) {
		expectedPagination := response.ResPagination{
			TotalRecords: 50,
			PerPage:      15,
			CurrentPage:  5,
			TotalPages:   4,
		}
		expectedSkip := 60
		actualPagination, actualSkip := utils.CalculatePagination(5, 15, 50)
		assert.Equal(t, expectedPagination, actualPagination)
		assert.Equal(t, expectedSkip, actualSkip)
	})

	t.Run("test list apps when passed type is free", func(t *testing.T) {
		expected := response.ResData{
			Apps: []models.AppsInfo{
				{
					AppName:        "Instagram",
					Category:       "social",
					Size:           "20M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "0",
					ContentRating:  "teen",
					Generes:        "social media",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "free",
				},
				{
					AppName:        "Snapchats",
					Category:       "social",
					Size:           "250M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "$2",
					ContentRating:  "everyone",
					Generes:        "social media",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "free",
				},
				{
					AppName:        "Whatsappss",
					Category:       "social",
					Size:           "250M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "0",
					ContentRating:  "everyone",
					Generes:        "social media",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "free",
				},
				{
					AppName:        "Skype",
					Category:       "communication",
					Size:           "20M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "0",
					ContentRating:  "everyone",
					Generes:        "communication",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "free",
				},
				{
					AppName:        "Hotstar",
					Category:       "movies",
					Size:           "30M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "0",
					ContentRating:  "everyone",
					Generes:        "movies",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "free",
				},
				{
					AppName:        "Mattermost",
					Category:       "communication",
					Size:           "30M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "0",
					ContentRating:  "everyone",
					Generes:        "communication",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "free",
				},
			},
			Pagination: &response.ResPagination{
				TotalRecords: 6,
				PerPage:      10,
				CurrentPage:  1,
				TotalPages:   1,
			},
		}
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/apps?type=free")

		assert.Nil(t, err)
		assert.Equal(t, expected.Pagination, actual.Body.Data.Pagination)
		assert.Equal(t, 6, len(actual.Body.Data.Apps))
		assert.Equal(t, http.StatusOK, res.StatusCode())

		for i, val := range actual.Body.Data.Apps {
			assert.Equal(t, expected.Apps[i].AppName, val.AppName, "expected and actual appname is not equal")
			assert.Equal(t, expected.Apps[i].Category, val.Category, "expected and actual category is not equal")
			assert.Equal(t, expected.Apps[i].Size, val.Size, "expected and actual size is not equal")
			assert.Equal(t, expected.Apps[i].Rating, val.Rating, "expected and actual rating is not equal")
			assert.Equal(t, expected.Apps[i].Reviews, val.Reviews, "expected and actual reviews is not equal")
			assert.Equal(t, expected.Apps[i].Installs, val.Installs, "expected and actual installs is not equal")
			assert.Equal(t, expected.Apps[i].Price, val.Price, "expected and actual price is not equal")
			assert.Equal(t, expected.Apps[i].ContentRating, val.ContentRating, "expected and content rating appname is not equal")
			assert.Equal(t, expected.Apps[i].Generes, val.Generes, "expected and actual genres is not equal")
			assert.Equal(t, expected.Apps[i].CurrentVersion, val.CurrentVersion, "expected and actual current version is not equal")
			assert.Equal(t, expected.Apps[i].AndroidVersion, val.AndroidVersion, "expected and actual android version is not equal")
			assert.Equal(t, expected.Apps[i].Types, val.Types, "expected and actual types is not equal")
		}
	})

	t.Run("test list apps when query category social and type is paid", func(t *testing.T) {
		expected := response.ResData{
			Apps: []models.AppsInfo{
				{
					AppName:        "Google Chrome",
					Category:       "social",
					Size:           "250M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "0",
					ContentRating:  "everyone",
					Generes:        "social media",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "paid",
				},
				{
					AppName:        "Facebookd",
					Category:       "social",
					Size:           "250M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "0",
					ContentRating:  "everyone",
					Generes:        "social media",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "paid",
				},
				{
					AppName:        "Telegram",
					Category:       "social media",
					Size:           "250M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "0",
					ContentRating:  "everyone",
					Generes:        "social media",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "paid",
				},
			},
			Pagination: &response.ResPagination{
				TotalRecords: 3,
				PerPage:      10,
				CurrentPage:  1,
				TotalPages:   1,
			},
		}
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/apps?type=paid&category=social")

		assert.Nil(t, err)
		assert.Equal(t, expected.Pagination, actual.Body.Data.Pagination)
		assert.Equal(t, 3, len(actual.Body.Data.Apps))
		assert.Equal(t, http.StatusOK, res.StatusCode())

		for i, val := range actual.Body.Data.Apps {
			assert.Equal(t, expected.Apps[i].AppName, val.AppName, "expected and actual appname is not equal")
			assert.Equal(t, expected.Apps[i].Category, val.Category, "expected and actual category is not equal")
			assert.Equal(t, expected.Apps[i].Size, val.Size, "expected and actual size is not equal")
			assert.Equal(t, expected.Apps[i].Rating, val.Rating, "expected and actual rating is not equal")
			assert.Equal(t, expected.Apps[i].Reviews, val.Reviews, "expected and actual reviews is not equal")
			assert.Equal(t, expected.Apps[i].Installs, val.Installs, "expected and actual installs is not equal")
			assert.Equal(t, expected.Apps[i].Price, val.Price, "expected and actual price is not equal")
			assert.Equal(t, expected.Apps[i].ContentRating, val.ContentRating, "expected and content rating appname is not equal")
			assert.Equal(t, expected.Apps[i].Generes, val.Generes, "expected and actual genres is not equal")
			assert.Equal(t, expected.Apps[i].CurrentVersion, val.CurrentVersion, "expected and actual current version is not equal")
			assert.Equal(t, expected.Apps[i].AndroidVersion, val.AndroidVersion, "expected and actual android version is not equal")
			assert.Equal(t, expected.Apps[i].Types, val.Types, "expected and actual types is not equal")
		}
	})

	t.Run("test list apps when query search starts with 's' and type is free", func(t *testing.T) {
		expected := response.ResData{
			Apps: []models.AppsInfo{
				{
					AppName:        "Snapchats",
					Category:       "social",
					Size:           "250M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "$2",
					ContentRating:  "everyone",
					Generes:        "social media",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "free",
				},
				{
					AppName:        "Skype",
					Category:       "communication",
					Size:           "20M",
					Rating:         "0",
					Reviews:        "0",
					Installs:       "0",
					Price:          "0",
					ContentRating:  "everyone",
					Generes:        "communication",
					CurrentVersion: "2.0.0",
					AndroidVersion: "5.1",
					Types:          "free",
				},
			},
			Pagination: &response.ResPagination{
				TotalRecords: 2,
				PerPage:      10,
				CurrentPage:  1,
				TotalPages:   1,
			},
		}
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/apps?type=free&search=s")

		assert.Nil(t, err)
		assert.Equal(t, expected.Pagination, actual.Body.Data.Pagination)
		assert.Equal(t, 2, len(actual.Body.Data.Apps))
		assert.Equal(t, http.StatusOK, res.StatusCode())

		for i, val := range actual.Body.Data.Apps {
			assert.Equal(t, expected.Apps[i].AppName, val.AppName, "expected and actual appname is not equal")
			assert.Equal(t, expected.Apps[i].Category, val.Category, "expected and actual category is not equal")
			assert.Equal(t, expected.Apps[i].Size, val.Size, "expected and actual size is not equal")
			assert.Equal(t, expected.Apps[i].Rating, val.Rating, "expected and actual rating is not equal")
			assert.Equal(t, expected.Apps[i].Reviews, val.Reviews, "expected and actual reviews is not equal")
			assert.Equal(t, expected.Apps[i].Installs, val.Installs, "expected and actual installs is not equal")
			assert.Equal(t, expected.Apps[i].Price, val.Price, "expected and actual price is not equal")
			assert.Equal(t, expected.Apps[i].ContentRating, val.ContentRating, "expected and content rating appname is not equal")
			assert.Equal(t, expected.Apps[i].Generes, val.Generes, "expected and actual genres is not equal")
			assert.Equal(t, expected.Apps[i].CurrentVersion, val.CurrentVersion, "expected and actual current version is not equal")
			assert.Equal(t, expected.Apps[i].AndroidVersion, val.AndroidVersion, "expected and actual android version is not equal")
			assert.Equal(t, expected.Apps[i].Types, val.Types, "expected and actual types is not equal")
		}
	})

	t.Run("test list apps when passed type is other than free or paid", func(t *testing.T) {
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/apps?type=paaaa")

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode())
	})

	t.Run("test list apps search query passed is not matched", func(t *testing.T) {
		res, err := client.
			R().
			Get("/api/v1/apps?search=qqqqqq")

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode())
	})

	t.Run("test list apps category query passed is not matched", func(t *testing.T) {
		res, err := client.
			R().
			Get("/api/v1/apps?category=qqqqqq")

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode())
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from apps")
		assert.Nil(t, err)
	})
}

func TestListCategory(t *testing.T) {
	var actual utils.ResponseListCategory

	postData := []structs.ReqCreateApp{
		{
			AppName:        "Instagram",
			Category:       "social",
			Size:           "20M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "teen",
			Generes:        "social media",
			LastUpdated:    "wednesday, 20 Aug 2023",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Google Chrome",
			Category:       "social",
			Size:           "250M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "social media",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Snapchats",
			Category:       "social",
			Size:           "250M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "$2",
			ContentRating:  "everyone",
			Generes:        "social media",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "paid",
		},
		{
			AppName:        "Facebookd",
			Category:       "social",
			Size:           "250M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "social media",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Whatsappss",
			Category:       "social",
			Size:           "250M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "social media",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Photo Studio",
			Category:       "Photography",
			Size:           "10M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "social media",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Skype",
			Category:       "communication",
			Size:           "20M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "communication",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Linkedin",
			Category:       "learning",
			Size:           "25M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "learning",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Hotstar",
			Category:       "movies",
			Size:           "30M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "movies",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Netflix",
			Category:       "movies",
			Size:           "40M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "movies",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Mattermost",
			Category:       "communication",
			Size:           "30M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "communication",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
		{
			AppName:        "Telegram",
			Category:       "social media",
			Size:           "250M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "everyone",
			Generes:        "social media",
			LastUpdated:    "Friday, 5 Jan 2024",
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		},
	}
	for _, val := range postData {
		_, err := client.R().SetBody(val).Post("/api/v1/apps")
		if err != nil {
			log.Fatal("error while inserting apps", zap.Error(err))
		}
	}

	t.Run("test list categories", func(t *testing.T) {
		expected := []string{"Photography", "Communication", "Learning", "Movies", "Social", "Social media"}
		res, err := client.
			R().
			SetResult(&actual.Body).
			Get("/api/v1/apps/category")

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.Equal(t, expected, actual.Body.Data.Category)
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from apps")
		assert.Nil(t, err)
	})
}

func TestDeleteApp(t *testing.T) {
	var actual utils.ResponseDeleteApp

	postData := structs.ReqCreateApp{
		AppName:        "spiderman",
		Category:       "social",
		Size:           "20M",
		Rating:         "0",
		Reviews:        "0",
		Installs:       "0",
		Price:          "0",
		ContentRating:  "teen",
		Generes:        "social media",
		CurrentVersion: "2.0.0",
		AndroidVersion: "5.1",
		Types:          "free",
	}

	_, err := client.R().SetBody(postData).SetResult(&actual.Body).Post("/api/v1/apps")
	if err != nil {
		log.Fatal("error while deleting apps", zap.Error(err))
	}

	t.Run("test delete app", func(t *testing.T) {
		res, err := client.
			R().
			SetResult(&actual).
			Delete("/api/v1/apps/" + strconv.Itoa(actual.Body.Data.Id))

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.NotEqual(t, nil, actual.Body.Data.DeletedAt)
	})

	t.Run("test delete app with 404 status code", func(t *testing.T) {
		res, err := client.
			R().
			Get("/api/v1/apps/" + strconv.Itoa(actual.Body.Data.Id))

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode())
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from apps")
		assert.Nil(t, err)
	})
}

func TestGetApp(t *testing.T) {
	var actual utils.ResponseGetApp
	postData := structs.ReqCreateApp{
		AppName:        "spiderman",
		Category:       "social",
		Size:           "20M",
		Rating:         "0",
		Reviews:        "0",
		Installs:       "0",
		Price:          "0",
		ContentRating:  "teen",
		Generes:        "social media",
		CurrentVersion: "2.0.0",
		AndroidVersion: "5.1",
		Types:          "free",
	}

	_, err := client.R().SetBody(postData).SetResult(&actual.Body).Post("/api/v1/apps")
	if err != nil {
		log.Fatal("error while posting apps", zap.Error(err))
	}

	t.Run("test get app", func(t *testing.T) {
		var actualApp utils.ResponseGetApp
		res, err := client.
			R().
			SetResult(&actualApp.Body).
			Get("/api/v1/apps/" + strconv.Itoa(actual.Body.Data.Id))

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.Equal(t, actual.Body.Data.AppName, actualApp.Body.Data.AppName)
		assert.Equal(t, actual.Body.Data.Category, actualApp.Body.Data.Category)
		assert.Equal(t, actual.Body.Data.Rating, actualApp.Body.Data.Rating)
		assert.Equal(t, actual.Body.Data.Reviews, actualApp.Body.Data.Reviews)
		assert.Equal(t, actual.Body.Data.Size, actualApp.Body.Data.Size)
		assert.Equal(t, actual.Body.Data.ContentRating, actualApp.Body.Data.ContentRating)
		assert.Equal(t, actual.Body.Data.Generes, actualApp.Body.Data.Generes)
		assert.Equal(t, actual.Body.Data.Types, actualApp.Body.Data.Types)
		assert.Equal(t, actual.Body.Data.CurrentVersion, actualApp.Body.Data.CurrentVersion)
		assert.Equal(t, actual.Body.Data.AndroidVersion, actualApp.Body.Data.AndroidVersion)
	})
}

func TestUpdateApp(t *testing.T) {
	presetTime := time.Now()
	var actual utils.ResponseUpdateApp
	postData := structs.ReqCreateApp{
		AppName:        "Spiderman",
		Category:       "social",
		Size:           "20M",
		Rating:         "0",
		Reviews:        "0",
		Installs:       "0",
		Price:          "0",
		ContentRating:  "teen",
		Generes:        "social media",
		LastUpdated:    presetTime.Format("Monday, 02 Jan 2006"),
		CurrentVersion: "2.0.0",
		AndroidVersion: "5.1",
		Types:          "free",
	}

	_, err := client.R().SetBody(postData).SetResult(&actual.Body).Post("/api/v1/apps")
	if err != nil {
		log.Fatal("error while posting app", zap.Error(err))
	}
	t.Run("test updated app with valid body", func(t *testing.T) {
		var actualUpdated utils.ResponseCreateApp
		currentTime := time.Now()
		req := structs.ReqCreateApp{
			AppName: "Spiderman - amazing",
		}
		expected := structs.ReqCreateApp{
			AppName:        "Spiderman - amazing",
			Category:       "social",
			Size:           "20M",
			Rating:         "0",
			Reviews:        "0",
			Installs:       "0",
			Price:          "0",
			ContentRating:  "teen",
			Generes:        "social media",
			LastUpdated:    currentTime.Format("Monday, 02 Jan 2006"),
			CurrentVersion: "2.0.0",
			AndroidVersion: "5.1",
			Types:          "free",
		}
		res, err := client.
			R().
			EnableTrace().
			SetBody(req).
			SetResult(&actualUpdated.Body).
			Put("/api/v1/apps/" + strconv.Itoa(actual.Body.Data.Id))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode())
		assert.Equal(t, expected.AppName, actualUpdated.Body.Data.AppName, "expected and actual appname is not equal")
		assert.Equal(t, expected.Category, actualUpdated.Body.Data.Category, "expected and actual category is not equal")
		assert.Equal(t, expected.Size, actualUpdated.Body.Data.Size, "expected and actual size is not equal")
		assert.Equal(t, expected.Rating, actualUpdated.Body.Data.Rating, "expected and actual rating is not equal")
		assert.Equal(t, expected.Reviews, actualUpdated.Body.Data.Reviews, "expected and actual reviews is not equal")
		assert.Equal(t, expected.Installs, actualUpdated.Body.Data.Installs, "expected and actual installs is not equal")
		assert.Equal(t, expected.Price, actualUpdated.Body.Data.Price, "expected and actual price is not equal")
		assert.Equal(t, expected.ContentRating, actualUpdated.Body.Data.ContentRating, "expected and content rating appname is not equal")
		assert.Equal(t, expected.Generes, actualUpdated.Body.Data.Generes, "expected and actual genres is not equal")
		assert.Equal(t, expected.LastUpdated, actualUpdated.Body.Data.LastUpdated, "expected and actual last updated is not equal")
		assert.Equal(t, expected.CurrentVersion, actualUpdated.Body.Data.CurrentVersion, "expected and actual current version is not equal")
		assert.Equal(t, expected.AndroidVersion, actualUpdated.Body.Data.AndroidVersion, "expected and actual android version is not equal")
		assert.Equal(t, expected.Types, actualUpdated.Body.Data.Types, "expected and actual types is not equal")
	})

	t.Run("test update app with invalid field", func(t *testing.T) {
		req := structs.ReqCreateApp{
			AppName: "",
		}
		res, err := client.
			R().
			SetBody(req).
			Put("/api/v1/apps/" + strconv.Itoa(actual.Body.Data.Id))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode())
	})

	t.Run("test update app when app does not exits", func(t *testing.T) {
		_, err := client.R().Delete("/api/v1/apps/" + strconv.Itoa(actual.Body.Data.Id))
		if err != nil {
			log.Fatal("error while deleting app", zap.Error(err))
		}

		req := structs.ReqCreateApp{
			AppName: "Amazing Spiderman 2",
		}

		res, err := client.
			R().
			SetBody(req).
			Put("/api/v1/apps/" + strconv.Itoa(actual.Body.Data.Id))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode())
	})

	t.Cleanup(func() {
		_, err := db.Exec("delete from apps")
		assert.Nil(t, err)
	})
}
