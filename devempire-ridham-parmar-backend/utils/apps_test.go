package utils

import (
	"testing"

	"github.com/Improwised/devempire-ridham-parmar-backend/constants"
	"github.com/Improwised/devempire-ridham-parmar-backend/models"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/structs"
	"github.com/stretchr/testify/assert"
)

// Test for reading file
func TestCanReadFile(t *testing.T) {
	assert := assert.New(t)
	expected := []models.AppsInfo{
		{AppName: "Sketch - Draw & Paint", Category: "ART_AND_DESIGN", Rating: "4.5", Reviews: "215644", Size: "25M", Installs: "50,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Art & Design", LastUpdated: "June 8, 2018", CurrentVersion: "Varies with device", AndroidVersion: "4.2 and up"},
		{AppName: "Pixel Draw - Number Art Coloring Book", Category: "ART_AND_DESIGN", Rating: "4.3", Reviews: "967", Size: "2.8M", Installs: "100,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "Creativity", CurrentVersion: "June 20, 2018", AndroidVersion: "1.1"},
		{AppName: "Paper flowers instructions", Category: "ART_AND_DESIGN", Rating: "4.4", Reviews: "167", Size: "5.6M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "March 26, 2017", CurrentVersion: "1", AndroidVersion: "2.3 and up"},
		{AppName: "Smoke Effect Photo Maker - Smoke Editor", Category: "ART_AND_DESIGN", Rating: "3.8", Reviews: "178", Size: "19M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "April 26, 2018", CurrentVersion: "1.1", AndroidVersion: "4.0.3 and up"},
	}
	actual := ReadFile[models.AppsInfo]("../" + constants.AppsInfoTestFile)
	assert.Equal(len(expected), len(actual), "length of actual and expected output is not same")
	for i := range actual {
		assert.Equal(expected[i], actual[i], "actual and expected output is not same")
	}
	assert.ElementsMatchf(expected, actual, "expected and actual output is not same")
}

func TestReturnValidFields(t *testing.T) {
	t.Run("test when valid field", func(t *testing.T) {
		req := structs.ReqCreateApp{
			AppName:        "spiderman app",
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

		expected := map[string]interface{}{
			"app_name":        "spiderman app",
			"category":        "social",
			"size":            "20M",
			"rating":          "0",
			"reviews":         "0",
			"type":            "free",
			"installs":        "0",
			"price":           "0",
			"content_rating":  "teen",
			"generes":         "social media",
			"current_version": "2.0.0",
			"android_version": "5.1",
		}

		actual := ReturnValidFields(req)

		assert.Equal(t, expected["app_name"], actual["app_name"])
		assert.Equal(t, expected["category"], actual["category"])
		assert.Equal(t, expected["size"], actual["size"])
		assert.Equal(t, expected["rating"], actual["rating"])
		assert.Equal(t, expected["reviews"], actual["reviews"])
		assert.Equal(t, expected["type"], actual["type"])
		assert.Equal(t, expected["installs"], actual["installs"])
		assert.Equal(t, expected["price"], actual["price"])
		assert.Equal(t, expected["content_rating"], actual["content_rating"])
		assert.Equal(t, expected["generes"], actual["generes"])
		assert.Equal(t, expected["current_version"], actual["current_version"])
		assert.Equal(t, expected["android_version"], actual["android_version"])

	})

	t.Run("test when invalid field", func(t *testing.T) {
		req := structs.ReqCreateApp{
			AppName: "",
		}

		expected := 0

		actual := ReturnValidFields(req)

		assert.Equal(t, expected, len(actual))

	})

}
