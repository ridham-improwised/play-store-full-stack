package utils

import (
	"log"
	"os"
	"regexp"
	"time"

	"github.com/Improwised/devempire-ridham-parmar-backend/models"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/structs"
	"github.com/doug-martin/goqu/v9"
	"github.com/jszwec/csvutil"
)

// Read file for now it is reading only AppsInfo
func ReadFile[T models.AppsInfo | models.ReviewInfo](filename string) []T {
	apps := []T{}
	file, readFileErr := os.ReadFile(filename)
	if readFileErr != nil {
		log.Fatal(readFileErr)
	}
	unmarshalErr := csvutil.Unmarshal(file, &apps)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return apps
}

func MatchString(str string) bool {
	pattern := "^[A-Za-z0-9 ]+$"
	regExp := regexp.MustCompile(pattern)
	return regExp.MatchString(str)
}

// creates map of required field
func ReturnValidFields(fields structs.ReqCreateApp) map[string]interface{} {
	presentTime := time.Now()
	validFields := make(map[string]interface{})

	defaultFields := map[string]interface{}{
		"app_name":        fields.AppName,
		"category":        fields.Category,
		"size":            fields.Size,
		"rating":          fields.Rating,
		"reviews":         fields.Reviews,
		"type":            fields.Types,
		"installs":        fields.Installs,
		"price":           fields.Price,
		"content_rating":  fields.ContentRating,
		"generes":         fields.Generes,
		"current_version": fields.CurrentVersion,
		"android_version": fields.AndroidVersion,
	}

	for key, val := range defaultFields {
		if defaultFields[key] != "" {
			validFields[key] = val
		}
	}

	if len(validFields) == 0 {
		return nil
	}
	validFields["last_updated"] = presentTime.Format("Monday, 02 Jan 2006")
	validFields["updated_at"] = time.Now()
	return validFields
}

// return validCondition for db
func ReturnValidCondition(queryParams map[string]string) goqu.Expression {
	conditions := []goqu.Expression{goqu.I("deleted_at").IsNull()}

	columnMapping := map[string]string{
		"type":     "type",
		"category": "category",
		"search":   "app_name",
	}

	for key, value := range queryParams {
		if _, ok := columnMapping[key]; ok {
			conditions = append(conditions, goqu.I(columnMapping[key]).ILike(value+"%"))
		}
	}
	return goqu.And(conditions...)
}
