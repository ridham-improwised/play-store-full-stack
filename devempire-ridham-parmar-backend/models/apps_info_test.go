package models

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for getting names of apps
func TestCanListAppsName(t *testing.T) {
	appInstance := NewAppInfo()
	assert := assert.New(t)
	expected := []string{"Sketch - Draw & Paint", "Pixel Draw - Number Art Coloring Book", "Paper flowers instructions", "Smoke Effect Photo Maker - Smoke Editor"}

	appsStruct := []AppsInfo{
		{AppName: "Sketch - Draw & Paint", Category: "ART_AND_DESIGN", Rating: "4.5", Reviews: "215644", Size: "25M", Installs: "50,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Art & Design", LastUpdated: "June 8, 2018", CurrentVersion: "Varies with device", AndroidVersion: "4.2 and up"},
		{AppName: "Pixel Draw - Number Art Coloring Book", Category: "ART_AND_DESIGN", Rating: "4.3", Reviews: "967", Size: "2.8M", Installs: "100,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "Creativity", CurrentVersion: "June 20, 2018", AndroidVersion: "1.1"},
		{AppName: "Paper flowers instructions", Category: "ART_AND_DESIGN", Rating: "4.4", Reviews: "167", Size: "5.6M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "March 26, 2017", CurrentVersion: "1", AndroidVersion: "2.3 and up"},
		{AppName: "Smoke Effect Photo Maker - Smoke Editor", Category: "ART_AND_DESIGN", Rating: "3.8", Reviews: "178", Size: "19M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "April 26, 2018", CurrentVersion: "1.1", AndroidVersion: "4.0.3 and up"},
	}

	actual := appInstance.ListNames(appsStruct)
	assert.Equalf(len(expected), len(actual), "length of actual and expected output is not same")

	assert.ElementsMatchf(expected, actual, "expected and actual output is not same")
}

// Test for getting app details
func TestCanGetAppInfoDetails(t *testing.T) {
	appInstance := NewAppInfo()
	assert := assert.New(t)
	expected := AppsInfo{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"}

	appsTruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
	}

	actual, ok := appInstance.GetApp(appsTruct, "Facebook")
	assert.Equal(true, ok, "expected and actual output is not same")
	assert.Equal(expected, actual, "expected and actual output is not same")
}

// Test for getting list of apps by search argument
func TestCanListAppsByArg(t *testing.T) {
	assert := assert.New(t)
	appInstance := NewAppInfo()
	expected := []string{"Facebook Pages Manager", "Facebook", "Facebook Lite", "Facebook Creator"}

	lists := []string{"Facebook Pages Manager", "Fingerprint Lock Screen Prank", "The FP Shield", "Chat For Strangers - Video Chat", "Stylish Fonts", "Facebook", "Facebook Lite", "MARKET FO", "Facebook Creator", "Gun Builder ELITE"}

	actual := appInstance.ListAppsByArg(lists, "Face")

	assert.Equal(expected, actual)
}

// Test for getting apps by generes wise and those which are paid
func TestCanGetAppByGeneresPrice(t *testing.T) {
	appInstance := NewAppInfo()
	assert := assert.New(t)
	appsTruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
	}

	expected := []AppsInfo{
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
	}

	actual := appInstance.GetAppByGeneresPrice(appsTruct, "Social")

	assert.ElementsMatchf(expected, actual, "expected and actual output is not same")
}

// Test for sorting apps by most installs
func TestCanSortByInstalls(t *testing.T) {
	appInstance := NewAppInfo()
	assert := assert.New(t)
	appsTruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
	}

	expected := []AppsInfo{
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
	}

	actual := appInstance.SortByInstalls(appsTruct)

	assert.Equalf(len(expected), len(actual), "length of actual and expected output is not same")

	assert.Equal(true, reflect.DeepEqual(expected, actual), "expected and actual output is not same")
}

// Test for getting price wise apps
func TestCanGetAppByPrice(t *testing.T) {
	appInstance := NewAppInfo()
	assert := assert.New(t)
	appsTruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
	}

	expected := []AppsInfo{
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "Varies with device", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
	}

	actual := appInstance.ListAppByPrice(appsTruct)

	assert.ElementsMatchf(expected, actual, "expected and actual output is not same")
}

// Test for getting size of apps
func TestCanListSize(t *testing.T) {
	assert := assert.New(t)
	appsInstance := NewAppInfo()

	appStruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "20M", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "78M", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Sketch - Draw & Paint", Category: "ART_AND_DESIGN", Rating: "4.5", Reviews: "215644", Size: "25M", Installs: "50,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Art & Design", LastUpdated: "June 8, 2018", CurrentVersion: "Varies with device", AndroidVersion: "4.2 and up"},
		{AppName: "Paper flowers instructions", Category: "ART_AND_DESIGN", Rating: "4.4", Reviews: "167", Size: "5.6M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "March 26, 2017", CurrentVersion: "1", AndroidVersion: "2.3 and up"},
		{AppName: "Paper flowers instructions", Category: "ART_AND_DESIGN", Rating: "4.4", Reviews: "167", Size: "49M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "March 26, 2017", CurrentVersion: "1", AndroidVersion: "2.3 and up"},
	}

	expected := []float64{12.50, 25.00, 12.50, 37.50, 12.50}

	actual := appsInstance.ListSize(appStruct)

	assert.Equal(len(expected), len(actual), "length of expected and actual output is not same")
	assert.ElementsMatchf(expected, actual, "expected and actual output is not same")
}

// Test for getting number of unique generes
func TestCanUniqueGeneresNumber(t *testing.T) {
	assert := assert.New(t)
	appsInstance := NewAppInfo()

	appStruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "20M", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "78M", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Sketch - Draw & Paint", Category: "ART_AND_DESIGN", Rating: "4.5", Reviews: "215644", Size: "25M", Installs: "50,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Art & Design", LastUpdated: "June 8, 2018", CurrentVersion: "Varies with device", AndroidVersion: "4.2 and up"},
		{AppName: "Paper flowers instructions", Category: "ART_AND_DESIGN", Rating: "4.4", Reviews: "167", Size: "5.6M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "March 26, 2017", CurrentVersion: "1", AndroidVersion: "2.3 and up"},
	}

	expected := 3

	actual := appsInstance.UniqueGeneresNumber(appStruct)

	assert.Equal(expected, actual, "expected and actual output is not same")
}

// Test for getting unique categories
func TestCanUniqueCategoriesNumber(t *testing.T) {
	assert := assert.New(t)
	appsInstance := NewAppInfo()

	appStruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "20M", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "78M", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Sketch - Draw & Paint", Category: "ART_AND_DESIGN", Rating: "4.5", Reviews: "215644", Size: "25M", Installs: "50,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Art & Design", LastUpdated: "June 8, 2018", CurrentVersion: "Varies with device", AndroidVersion: "4.2 and up"},
		{AppName: "Paper flowers instructions", Category: "ART_AND_DESIGN", Rating: "4.4", Reviews: "167", Size: "5.6M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "March 26, 2017", CurrentVersion: "1", AndroidVersion: "2.3 and up"},
	}

	expected := 3

	actual := appsInstance.UniqueCategoriesNumber(appStruct)

	assert.Equal(expected, actual, "expected and actual output is not same")
}

// Test for getting most installed apps
func TestCanMostInstalledApps(t *testing.T) {
	assert := assert.New(t)
	appsInstance := NewAppInfo()

	appStruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "20M", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "78M", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Sketch - Draw & Paint", Category: "ART_AND_DESIGN", Rating: "4.5", Reviews: "215644", Size: "25M", Installs: "50,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Art & Design", LastUpdated: "June 8, 2018", CurrentVersion: "Varies with device", AndroidVersion: "4.2 and up"},
		{AppName: "Paper flowers instructions", Category: "ART_AND_DESIGN", Rating: "4.4", Reviews: "167", Size: "5.6M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "March 26, 2017", CurrentVersion: "1", AndroidVersion: "2.3 and up"},
	}

	expected := []string{"Facebook", "HTC Social Plugin - Facebook"}

	actual := appsInstance.MostInstalledApp(appStruct)

	assert.Equal(expected, actual, "expected and actual output is not same")
}

// Test for getting most installed apps
func TestCanLeastInstalledApps(t *testing.T) {
	assert := assert.New(t)
	appsInstance := NewAppInfo()

	appStruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "20M", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "78M", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Sketch - Draw & Paint", Category: "ART_AND_DESIGN", Rating: "4.5", Reviews: "215644", Size: "25M", Installs: "50,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Art & Design", LastUpdated: "June 8, 2018", CurrentVersion: "Varies with device", AndroidVersion: "4.2 and up"},
		{AppName: "Paper flowers instructions", Category: "ART_AND_DESIGN", Rating: "4.4", Reviews: "167", Size: "5.6M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "March 26, 2017", CurrentVersion: "1", AndroidVersion: "2.3 and up"},
	}

	expected := []string{"Super Ear Super Hearing", "Blood Pressure Monitor"}

	actual := appsInstance.LeastInstalledApp(appStruct)

	assert.Equal(expected, actual, "expected and actual output is not same")
}

// Test for total play store size in GB
func TestCanTotalPlayStoreSize(t *testing.T) {
	assert := assert.New(t)
	appsInstance := NewAppInfo()

	appStruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "20M", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "78M", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Sketch - Draw & Paint", Category: "ART_AND_DESIGN", Rating: "4.5", Reviews: "215644", Size: "25M", Installs: "50,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Art & Design", LastUpdated: "June 8, 2018", CurrentVersion: "Varies with device", AndroidVersion: "4.2 and up"},
		{AppName: "Paper flowers instructions", Category: "ART_AND_DESIGN", Rating: "4.4", Reviews: "167", Size: "5.6M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "March 26, 2017", CurrentVersion: "1", AndroidVersion: "2.3 and up"},
	}

	expected := "0.14"

	actual := appsInstance.TotalPlayStoreSize(appStruct)

	assert.Equal(expected, actual, "expected and actual output is not same")
}

// Test for getting total installs
func TestCanTotalInstalls(t *testing.T) {
	assert := assert.New(t)
	appsInstance := NewAppInfo()

	appStruct := []AppsInfo{
		{AppName: "Super Ear Super Hearing", Category: "MEDICAL", Rating: "4", Reviews: "122", Size: "980k", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "July 13, 2018", CurrentVersion: "1.4", AndroidVersion: "4.0 and up"},
		{AppName: "Blood Pressure Monitor", Category: "MEDICAL", Rating: "4.3", Reviews: "17", Size: "6.0M", Installs: "10,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Medical", LastUpdated: "February 25, 2017", CurrentVersion: "1.0.1", AndroidVersion: "4.4 and up"},
		{AppName: "Facebook Lite", Category: "SOCIAL", Rating: "4.3", Reviews: "8606259", Size: "Varies with device", Installs: "500,000,000+", Types: "Paid", Price: "10,55550+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 1, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "20M", Installs: "1,000,000,000+", Types: "Paid", Price: "50,0000+", ContentRating: "Teen", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "HTC Social Plugin - Facebook", Category: "SOCIAL", Rating: "4.1", Reviews: "78158306", Size: "78M", Installs: "1,000,000,000+", Types: "Free", Price: "10,000,000+", ContentRating: "Mature", Generes: "Social", LastUpdated: "August 3, 2018", CurrentVersion: "Varies with device", AndroidVersion: "Varies with device"},
		{AppName: "Sketch - Draw & Paint", Category: "ART_AND_DESIGN", Rating: "4.5", Reviews: "215644", Size: "25M", Installs: "50,000,000+", Types: "Free", Price: "0", ContentRating: "Teen", Generes: "Art & Design", LastUpdated: "June 8, 2018", CurrentVersion: "Varies with device", AndroidVersion: "4.2 and up"},
		{AppName: "Paper flowers instructions", Category: "ART_AND_DESIGN", Rating: "4.4", Reviews: "167", Size: "5.6M", Installs: "50,000+", Types: "Free", Price: "0", ContentRating: "Everyone", Generes: "Art & Design", LastUpdated: "March 26, 2017", CurrentVersion: "1", AndroidVersion: "2.3 and up"},
	}

	expected := 2550070000

	actual := appsInstance.TotalInstalls(appStruct)

	assert.Equal(expected, actual, "expected and actual output is not same")
}
