package models

import (
	"time"

	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/structs"
	"github.com/doug-martin/goqu/v9"
)

type ReviewInfo struct {
	Id                    int        `json:"id" db:"id"`
	AppName               string     `json:"app" csv:"app_name" db:"app_name"`
	TranslatedReview      string     `json:"translatedReview" csv:"translated_review" db:"translated_review"`
	Sentiment             string     `json:"sentiment" csv:"sentiment" db:"sentiment"`
	SentimentPolarity     string     `json:"sentimentPolarity,omitempty" csv:"sentiment_polarity" db:"sentiment_polarity,omitempty"`
	SentimentSubjectivity string     `json:"sentimentSubjectivity,omitempty" csv:"sentiment_subjectivity" db:"sentiment_subjectivity,omitempty"`
	CreatedAt             time.Time  `json:"createdAt" db:"created_at, omitempty"`
	UpdatedAt             time.Time  `json:"updatedAt" db:"updated_at, omitempty"`
	DeletedAt             *time.Time `json:"deletedAt,omitempty" db:"deleted_at, omitempty"`
}

// represents table name in database
const ReviewsTable = "reviews"

// AppsModel implements apps related database operations
type ReviewsModel struct {
	db *goqu.Database
}

func InitReviewsModel(goqu *goqu.Database) (ReviewsModel, error) {
	return ReviewsModel{
		db: goqu,
	}, nil
}

// InsertApp insert app details
func (model *ReviewsModel) InsertReviews(appname string, review structs.ReqCreateReviews) (ReviewInfo, error) {
	var resReview ReviewInfo

	_, err := model.db.Insert(ReviewsTable).Rows(
		goqu.Record{
			"app_name":               appname,
			"translated_review":      review.TranslatedReview,
			"sentiment":              review.Sentiment,
			"sentiment_polarity":     "nan",
			"sentiment_subjectivity": "nan",
		},
	).Returning("id", "app_name", "translated_review", "sentiment", "sentiment_polarity", "sentiment_subjectivity", "created_at", "updated_at").Executor().ScanStruct(&resReview)

	return resReview, err
}

// ListReviewsDetails list review details
func (model *ReviewsModel) ListReviews(app string, per_page, skip int) ([]ReviewInfo, error) {
	var reviews []ReviewInfo

	err := model.db.From(ReviewsTable).Where(goqu.I("app_name").ILike(app)).Limit(uint(per_page)).Offset(uint(skip)).ScanStructs(&reviews)
	return reviews, err
}

// CountReviews
func (model *ReviewsModel) CountReviews(app string) (int, error) {
	count, err := model.db.From(ReviewsTable).Where(goqu.I("app_name").ILike(app)).Count()
	return int(count), err
}

func AppsReviewInfo() *ReviewInfo {
	return &ReviewInfo{}
}

// Get app review info details
func (a *ReviewInfo) GetAppReviewInfoDetail(appsList []ReviewInfo, appName string) []ReviewInfo {
	app := []ReviewInfo{}
	for _, val := range appsList {
		if val.AppName == appName {
			app = append(app, val)
		}
	}
	return app
}

// Display app with most reviews
func (a *ReviewInfo) MostReviewedApp(appsList []ReviewInfo) []string {

	// Create a map to store the count of each string
	countMap := make(map[string]int)

	// Count occurrences
	for _, str := range appsList {
		countMap[str.AppName]++
	}

	// Find the maximum occurrence count
	maxCount := 0
	for _, count := range countMap {
		if count > maxCount {
			maxCount = count
		}
	}

	// Find strings with maximum occurrence
	var maxOccurrenceStrings []string
	for str, count := range countMap {
		if count == maxCount {
			maxOccurrenceStrings = append(maxOccurrenceStrings, str)
		}
	}

	return maxOccurrenceStrings
}

// Display app with least reviews
func (a *ReviewInfo) LeastReviewedApp(appsList []ReviewInfo) []string {

	// Create a map to store the count of each string
	countMap := make(map[string]int)

	// Count occurrences
	for _, str := range appsList {
		countMap[str.AppName]++
	}

	// Find the minimum occurrence count
	minCount := len(appsList)
	for _, count := range countMap {
		if count < minCount {
			minCount = count
		}
	}

	// Find strings with maximum occurrence
	var minOccurrenceStrings []string
	for str, count := range countMap {
		if count == minCount {
			minOccurrenceStrings = append(minOccurrenceStrings, str)
		}
	}

	return minOccurrenceStrings
}
