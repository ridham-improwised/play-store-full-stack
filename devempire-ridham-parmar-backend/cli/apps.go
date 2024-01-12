package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Improwised/devempire-ridham-parmar-backend/constants"
	"github.com/Improwised/devempire-ridham-parmar-backend/models"
	"github.com/Improwised/devempire-ridham-parmar-backend/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	appName     string
	skip, limit int
)

// ListAppsCmd shows apps name
func ListAppsCmd() cobra.Command {
	appsCommand := cobra.Command{
		Use:   "apps",
		Short: "To display list of apps name",
		Args:  cobra.ExactArgs(0),
		Long:  `To display list of apps name in tabular format with total count`,
		Example: `apps-store apps
apps-store apps -a="app name"
apps-store apps --appName="app name"`,

		PreRunE: func(cmd *cobra.Command, args []string) error {
			isSet := cmd.Flags().Lookup("appName").Changed
			isMatch := utils.MatchString(appName)
			if !isMatch && isSet {
				return errors.New("required flag --appName or -a has an empty value")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {

			appsInstance := models.NewAppInfo()
			appList := utils.ReadFile[models.AppsInfo](constants.AppsInfoFile)

			if appName != "" {
				apps, ok := appsInstance.GetApp(appList, appName)
				if !ok {
					fmt.Println("No App found")
					return
				}

				appReviewsList := utils.ReadFile[models.ReviewInfo](constants.AppsReviewInfoFile)
				appReviewInstance := models.AppsReviewInfo()
				appReviews := appReviewInstance.GetAppReviewInfoDetail(appReviewsList, appName)

				PrintDetails(apps, appReviews, skip, limit)
				return
			}

			appsName := appsInstance.ListNames(appList)
			PrintAppsName(appsName)
		},
	}
	appsCommand.Flags().IntVarP(&skip, "skip", "s", 0, "skips specified records")
	appsCommand.Flags().IntVarP(&limit, "limit", "l", 10, "displays specified records")
	return appsCommand
}

// Print Apps Name in tabular format
func PrintAppsName(appNames []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Apps Name"})
	for _, val := range appNames {
		table.Append([]string{val})
	}
	table.SetRowLine(true)
	table.SetFooter([]string{"Total : " + strconv.Itoa(len(appNames))})
	table.Render()
}

// Print App with its review
func PrintDetails(apps models.AppsInfo, reviews []models.ReviewInfo, skip, limit int) {
	// displays apps details
	table := tablewriter.NewWriter(os.Stdout)
	vertical := [][]string{
		{"APP NAME", apps.AppName},
		{"CATEGORY", apps.Category},
		{"RATING", apps.Rating},
		{"REVIEWS", apps.Reviews},
		{"SIZE", apps.Size},
		{"INSTALLS", apps.Installs},
		{"TYPES", apps.Types},
		{"PRICE", apps.Price},
		{"CONTENT RATING", apps.ContentRating},
		{"GENERES", apps.Generes},
		{"LAST UPDATED", apps.LastUpdated},
		{"CURRENT VER", apps.CurrentVersion},
		{"ANDROID VER", apps.AndroidVersion},
	}

	for _, val := range vertical {
		table.Append(val)
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()

	header := []string{"Reviews", "Sentiment", "Sentiment Polarity", "Sentiment Subjectivity"}
	PrintReviewsDetails(reviews, header, skip, limit)
}

func PrintReviewsDetails(reviews []models.ReviewInfo, header []string, skip, limit int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	if len(reviews) > 0 && len(reviews) > skip && skip >= 0 {
		table.SetHeader(header)
		table.SetFooter([]string{"Total : " + strconv.Itoa(len(reviews)), "skip : " + strconv.Itoa(skip), "limit : " + strconv.Itoa(limit)})
		for index, review := range reviews {
			if index < skip {
				continue
			}
			if limit == 0 {
				break
			}
			table.Append([]string{review.TranslatedReview, review.Sentiment, review.SentimentPolarity, review.SentimentSubjectivity})
			limit--
		}
	} else {
		table.SetFooter([]string{"No Records Founds"})
	}
	table.Render()
}
