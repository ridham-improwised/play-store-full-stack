package models

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/structs"
	"github.com/doug-martin/goqu/v9"
	"github.com/samber/lo"
)

type AppsInfo struct {
	Id             int        `json:"id,omitempty" db:"id"`
	AppName        string     `json:"app,omitempty" csv:"app_name" db:"app_name"`
	Category       string     `json:"category,omitempty" csv:"category" db:"category, omitempty"`
	Rating         string     `json:"rating,omitempty" csv:"rating" db:"rating"`
	Reviews        string     `json:"reviews,omitempty" csv:"reviews" db:"reviews"`
	Size           string     `json:"size,omitempty" csv:"size" db:"size"`
	Installs       string     `json:"installs,omitempty" csv:"installs" db:"installs"`
	Types          string     `json:"type,omitempty" csv:"type" db:"type"`
	Price          string     `json:"price,omitempty" csv:"price" db:"price"`
	ContentRating  string     `json:"contentRating,omitempty" csv:"content_rating" db:"content_rating"`
	Generes        string     `json:"generes,omitempty" csv:"generes" db:"generes"`
	LastUpdated    string     `json:"lastUpdated,omitempty" csv:"last_updated" db:"last_updated"`
	CurrentVersion string     `json:"currentVer,omitempty" csv:"current_version" db:"current_version"`
	AndroidVersion string     `json:"androidVer,omitempty" csv:"android_version" db:"android_version"`
	CreatedAt      time.Time  `json:"createdAt,omitempty" db:"created_at, omitempty"`
	UpdatedAt      time.Time  `json:"updatedAt,omitempty" db:"updated_at, omitempty"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty" db:"deleted_at, omitempty"`
}

// represents table name in database
const AppsTable = "apps"

// AppsModel implements apps related database operations
type AppsModel struct {
	db *goqu.Database
}

func InitAppsModel(goqu *goqu.Database) (AppsModel, error) {
	return AppsModel{
		db: goqu,
	}, nil
}

// ListApps list app details with default records of 10
func (model *AppsModel) ListAppsDetails(validCondition goqu.Expression, perPage int, skip int) ([]AppsInfo, error) {
	var apps []AppsInfo

	err := model.db.From(AppsTable).Where(validCondition).Limit(uint(perPage)).Offset(uint(skip)).ScanStructs(&apps)
	return apps, err
}

// CountApps
func (model *AppsModel) CountApps(validCondition goqu.Expression) (int, error) {
	count, err := model.db.From(AppsTable).Where(validCondition).Count()
	return int(count), err
}

// ListCategories lists category of apps
func (model *AppsModel) ListCategories() ([]string, error) {
	var categories []string

	if err := model.db.From(AppsTable).Select("category").Distinct().Order(goqu.I("category").Asc()).ScanVals(&categories); err != nil {
		return nil, err
	}
	for index, val := range categories {
		str := strings.ToLower(strings.Join(strings.Split(val, "_"), " "))
		categories[index] = strings.ToUpper(string(str[0])) + str[1:]
	}
	return categories, nil
}

// InsertApp insert app details
func (model *AppsModel) InsertApp(app structs.ReqCreateApp) (AppsInfo, error) {
	var resApp AppsInfo
	presentTime := time.Now()

	_, err := model.db.Insert(AppsTable).Rows(
		goqu.Record{
			"app_name":        app.AppName,
			"category":        app.Category,
			"rating":          app.Rating,
			"reviews":         app.Reviews,
			"size":            app.Size,
			"installs":        app.Installs,
			"type":            app.Types,
			"price":           app.Price,
			"content_rating":  app.ContentRating,
			"generes":         app.Generes,
			"last_updated":    presentTime.Format("Monday, 02 Jan 2006"),
			"current_version": app.CurrentVersion,
			"android_version": app.AndroidVersion,
		},
	).Returning("id", "app_name", "category", "rating", "reviews", "size", "installs", "type", "price", "content_rating", "generes", "last_updated", "current_version", "android_version", "created_at", "updated_at", "deleted_at").Executor().ScanStruct(&resApp)

	return resApp, err
}

// DeleteApp delete app details
func (model *AppsModel) DeleteAppDetail(appId int) (AppsInfo, error) {
	var resApp AppsInfo
	currentTime := time.Now()
	_, err := model.db.Update(AppsTable).Set(
		goqu.Record{
			"deleted_at": &currentTime,
		},
	).Where(goqu.Ex{"id": appId}).Returning("id", "app_name", "category", "rating", "reviews", "size", "installs", "type", "price", "content_rating", "generes", "last_updated", "current_version", "android_version", "created_at", "updated_at", "deleted_at").Executor().ScanStruct(&resApp)

	return resApp, err
}

// GetAppDetails displays app details
func (model *AppsModel) GetAppDetails(appId int) (AppsInfo, error) {
	var apps AppsInfo

	_, err := model.db.From(AppsTable).Where(goqu.C("id").Eq(appId), goqu.C("deleted_at").IsNull()).ScanStruct(&apps)
	return apps, err
}

// UpdateAppDetails updates details of app
func (model *AppsModel) UpdateAppDetails(appId int, validFields map[string]interface{}) (AppsInfo, error) {
	var resApp AppsInfo
	_, err := model.db.Update(AppsTable).Set(
		validFields,
	).Where(goqu.C("id").Eq(appId), goqu.I("deleted_at").IsNull()).Returning("id", "app_name", "category", "rating", "reviews", "size", "installs", "type", "price", "content_rating", "generes", "last_updated", "current_version", "android_version", "created_at", "updated_at", "deleted_at").Executor().ScanStruct(&resApp)

	return resApp, err
}

func NewAppInfo() *AppsInfo {
	return &AppsInfo{}
}

// List Apps Name
func (a *AppsInfo) ListNames(appsList []AppsInfo) []string {
	var app []string
	for _, val := range appsList {
		app = append(app, val.AppName)
	}
	return app
}

// Get single app details
func (a *AppsInfo) GetApp(appsList []AppsInfo, appName string) (AppsInfo, bool) {
	for _, val := range appsList {
		if val.AppName == appName {
			return val, true
		}
	}
	return AppsInfo{}, false
}

// List Apps Name by search argument
func (a *AppsInfo) ListAppsByArg(appsList []string, appName string) []string {
	var app []string
	appLower := strings.ToLower(appName)
	pattern := `^` + appLower
	regex := regexp.MustCompile(pattern)

	for _, val := range appsList {
		lowerStr := strings.ToLower(val)
		isMatch := regex.MatchString(lowerStr)
		if isMatch {
			app = append(app, val)
		}
	}
	return app
}

// Get Apps details by generes and those which are paid
func (a *AppsInfo) GetAppByGeneresPrice(appsList []AppsInfo, generes string) []AppsInfo {
	generes = strings.ToLower(generes)
	list := lo.Filter(appsList, func(val AppsInfo, index int) bool {
		return strings.ToLower(val.Generes) == generes && val.Types == "Paid"
	})
	return list
}

// Sort apps based on number of installs
func (a *AppsInfo) SortByInstalls(list []AppsInfo) []AppsInfo {
	regex := regexp.MustCompile(`[^0-9 ]+`)
	sort.SliceStable(list, func(i, j int) bool {
		m, err1 := strconv.Atoi(regex.ReplaceAllString(list[i].Installs, ""))
		if err1 != nil {
			panic(err1)
		}
		n, err2 := strconv.Atoi(regex.ReplaceAllString(list[j].Installs, ""))
		if err2 != nil {
			panic(err2)
		}
		return m > n
	})
	return list
}

// Get apps details by paid price
func (a *AppsInfo) ListAppByPrice(appsList []AppsInfo) []AppsInfo {
	list := lo.Filter(appsList, func(val AppsInfo, index int) bool {
		return val.Types == "Paid"
	})
	return list
}

// List apps by size
func (a *AppsInfo) ListSize(appsList []AppsInfo) []float64 {
	regex := regexp.MustCompile(`[^0-9]`)

	partitions := lo.PartitionBy(appsList, func(val AppsInfo) string {

		// comparing string type values
		switch {
		case val.Size == "Varies with device":
			return "varies with device"
		case strings.Contains(val.Size, "k"):
			return "KB"
		case strings.Contains(val.Size, "."):
			return "Less than to 10MB"
		}

		// comparing number types
		num, err := strconv.Atoi(regex.ReplaceAllString(val.Size, ""))
		if err != nil {
			panic(err)
		}
		switch {
		case num <= 50 && num >= 10:
			return "Less than 50MB & greater than 20MB"
		default:
			return "Greater than 50MB"
		}
	})

	appsSize := []float64{}

	// calculated percentage for specific app's size
	for _, val := range partitions {
		count := (float64(len(val)) * 100) / float64(len(appsList))
		appsSize = append(appsSize, count)
	}
	return appsSize
}

// number of unique generes
func (a *AppsInfo) UniqueGeneresNumber(appsList []AppsInfo) int {

	unique := lo.UniqBy(appsList, func(val AppsInfo) string {
		return val.Generes
	})
	return len(unique)
}

// number of unique categories
func (a *AppsInfo) UniqueCategoriesNumber(appsList []AppsInfo) int {

	unique := lo.UniqBy(appsList, func(val AppsInfo) string {
		return val.Category
	})
	return len(unique)
}

// Display app with most downloads
func (a *AppsInfo) MostInstalledApp(appsList []AppsInfo) []string {
	regex := regexp.MustCompile(`[^0-9]+`)

	maxInstalledSize := 0
	var maxValue string
	maxInstalledApps := make(map[string][]string)
	for _, val := range appsList {
		if regex.ReplaceAllString(val.Installs, "") == "" {
			continue
		}
		num, err := strconv.Atoi(regex.ReplaceAllString(val.Installs, ""))
		if err != nil {
			panic(err)
		}
		if num >= maxInstalledSize {
			maxInstalledSize = num
			maxValue = val.Installs
			maxInstalledApps[maxValue] = append(maxInstalledApps[maxValue], val.AppName)
		}
	}
	return maxInstalledApps[maxValue]
}

// Display app with least downloads
func (a *AppsInfo) LeastInstalledApp(appsList []AppsInfo) []string {
	regex := regexp.MustCompile(`[^0-9]+`)

	// assuming the maximum installed value got from most installed apps func
	minInstalledSize := 1000000000
	var minValue string
	minInstalledApps := make(map[string][]string)
	for _, val := range appsList {
		if regex.ReplaceAllString(val.Installs, "") == "" {
			continue
		}
		num, err := strconv.Atoi(regex.ReplaceAllString(val.Installs, ""))
		if err != nil {
			panic(err)
		}
		if num <= minInstalledSize {
			minInstalledSize = num
			minValue = val.Installs
			minInstalledApps[minValue] = append(minInstalledApps[minValue], val.AppName)
		}
	}
	return minInstalledApps[minValue]
}

// Display total play store size in GB
func (a *AppsInfo) TotalPlayStoreSize(appsList []AppsInfo) string {

	regex := regexp.MustCompile(`[^0-9\.]`)
	sumInKB := 0.0
	sumInMB := 0.0
	for _, val := range appsList {
		switch {
		case val.Size == "Varies with device":
			continue
		case strings.Contains(val.Size, "k"):
			size, err := strconv.ParseFloat(regex.ReplaceAllString(val.Size, ""), 32)
			if err != nil {
				panic(err)
			}
			sumInKB += size
		case strings.Contains(val.Size, "M"):
			size, err := strconv.ParseFloat(regex.ReplaceAllString(val.Size, ""), 32)
			if err != nil {
				panic(err)
			}
			sumInMB += size
		}
	}
	return strconv.FormatFloat(((sumInKB/1000)+sumInMB)/1000, 'f', 2, 64)
}

// Display total installs
func (a *AppsInfo) TotalInstalls(appsList []AppsInfo) int {
	regex := regexp.MustCompile(`[^0-9]+`)

	totalInstalls := 0
	for _, val := range appsList {
		if regex.ReplaceAllString(val.Installs, "") == "" {
			continue
		}
		num, err := strconv.Atoi(regex.ReplaceAllString(val.Installs, ""))
		if err != nil {
			panic(err)
		}
		totalInstalls += num
	}

	return totalInstalls
}
