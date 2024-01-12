package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for getting review details of app
func TestCanGetAppReviewDetail(t *testing.T) {
	assert := assert.New(t)
	appInstance := AppsReviewInfo()
	expected := []ReviewInfo{
		{AppName: "Facebook", TranslatedReview: "I believe lot Google need slow little. There enough recent improvements satisfy social media world time being. Let us digest latest exponential leaps brilliant design development user friendly fixes wiring. I mean, EVERY WEEK !", Sentiment: "Positive", SentimentPolarity: "0.10335648148148148", SentimentSubjectivity: "0.5337962962962963"},
		{AppName: "Facebook", TranslatedReview: "I know went wrong app. I issues facebook visit phone..the hand laggy, load content properly lot times worst. I get notifications timely, want open rgem, facebook refuses load them. Bad basically, dislike it. I hope update properly soon", Sentiment: "Negative", SentimentPolarity: "-0.36666666666666664", SentimentSubjectivity: "0.5444444444444444"},
	}
	appsStruct := []ReviewInfo{
		{AppName: "Facebook", TranslatedReview: "I believe lot Google need slow little. There enough recent improvements satisfy social media world time being. Let us digest latest exponential leaps brilliant design development user friendly fixes wiring. I mean, EVERY WEEK !", Sentiment: "Positive", SentimentPolarity: "0.10335648148148148", SentimentSubjectivity: "0.5337962962962963"},
		{AppName: "Facebook", TranslatedReview: "I know went wrong app. I issues facebook visit phone..the hand laggy, load content properly lot times worst. I get notifications timely, want open rgem, facebook refuses load them. Bad basically, dislike it. I hope update properly soon", Sentiment: "Negative", SentimentPolarity: "-0.36666666666666664", SentimentSubjectivity: "0.5444444444444444"},
		{AppName: "10 Best Foods for You", TranslatedReview: "One greatest apps.", Sentiment: "Positive", SentimentPolarity: "1.0", SentimentSubjectivity: "1.0"},
		{AppName: "10 Best Foods for You", TranslatedReview: "good nice", Sentiment: "Positive", SentimentPolarity: "0.6499999999999999", SentimentSubjectivity: "0.8"},
	}

	actual := appInstance.GetAppReviewInfoDetail(appsStruct, "Facebook")

	assert.Equalf(len(expected), len(actual), "length of actual and expected output is not same")

	assert.ElementsMatchf(expected, actual, "expected and actual output is not same")
}

// Test for most reviwed apps
func TestCanMostReviwedApps(t *testing.T) {
	assert := assert.New(t)
	appInstance := AppsReviewInfo()
	appsStruct := []ReviewInfo{
		{AppName: "Girly Lock Screen Wallpaper with Quotes", TranslatedReview: "", Sentiment: "", SentimentPolarity: "", SentimentSubjectivity: ""},
		{AppName: "Facebook", TranslatedReview: "I believe lot Google need slow little. There enough recent improvements satisfy social media world time being. Let us digest latest exponential leaps brilliant design development user friendly fixes wiring. I mean, EVERY WEEK !", Sentiment: "Positive", SentimentPolarity: "0.10335648148148148", SentimentSubjectivity: "0.5337962962962963"},
		{AppName: "Facebook", TranslatedReview: "I know went wrong app. I issues facebook visit phone..the hand laggy, load content properly lot times worst. I get notifications timely, want open rgem, facebook refuses load them. Bad basically, dislike it. I hope update properly soon", Sentiment: "Negative", SentimentPolarity: "-0.36666666666666664", SentimentSubjectivity: "0.5444444444444444"},
		{AppName: "10 Best Foods for You", TranslatedReview: "One greatest apps.", Sentiment: "Positive", SentimentPolarity: "1.0", SentimentSubjectivity: "1.0"},
		{AppName: "10 Best Foods for You", TranslatedReview: "Amazing", Sentiment: "Positive", SentimentPolarity: "1.0", SentimentSubjectivity: "1.0"},
		{AppName: "10 Best Foods for You", TranslatedReview: "Good health...... Good health first", Sentiment: "Positive", SentimentPolarity: "1.0", SentimentSubjectivity: "1.0"},
	}

	expected := []string{"10 Best Foods for You"}

	actual := appInstance.MostReviewedApp(appsStruct)

	assert.Equal(expected, actual, "expected and actual output is not same")
}

// Test for least reviwed apps
func TestCanLeastReviwedApps(t *testing.T) {
	assert := assert.New(t)
	appInstance := AppsReviewInfo()
	appsStruct := []ReviewInfo{
		{AppName: "Girly Lock Screen Wallpaper with Quotes", TranslatedReview: "", Sentiment: "", SentimentPolarity: "", SentimentSubjectivity: ""},
		{AppName: "Facebook", TranslatedReview: "I believe lot Google need slow little. There enough recent improvements satisfy social media world time being. Let us digest latest exponential leaps brilliant design development user friendly fixes wiring. I mean, EVERY WEEK !", Sentiment: "Positive", SentimentPolarity: "0.10335648148148148", SentimentSubjectivity: "0.5337962962962963"},
		{AppName: "Facebook", TranslatedReview: "I know went wrong app. I issues facebook visit phone..the hand laggy, load content properly lot times worst. I get notifications timely, want open rgem, facebook refuses load them. Bad basically, dislike it. I hope update properly soon", Sentiment: "Negative", SentimentPolarity: "-0.36666666666666664", SentimentSubjectivity: "0.5444444444444444"},
		{AppName: "10 Best Foods for You", TranslatedReview: "One greatest apps.", Sentiment: "Positive", SentimentPolarity: "1.0", SentimentSubjectivity: "1.0"},
		{AppName: "10 Best Foods for You", TranslatedReview: "Amazing", Sentiment: "Positive", SentimentPolarity: "1.0", SentimentSubjectivity: "1.0"},
		{AppName: "10 Best Foods for You", TranslatedReview: "Good health...... Good health first", Sentiment: "Positive", SentimentPolarity: "1.0", SentimentSubjectivity: "1.0"},
	}

	expected := []string{"Girly Lock Screen Wallpaper with Quotes"}

	actual := appInstance.LeastReviewedApp(appsStruct)

	assert.Equal(expected, actual, "expected and actual output is not same")
}
