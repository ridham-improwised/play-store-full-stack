package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Improwised/devempire-ridham-parmar-backend/constants"
	"github.com/Improwised/devempire-ridham-parmar-backend/models"
	"github.com/Improwised/devempire-ridham-parmar-backend/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func GetAnalyticsCmd() cobra.Command {
	analyticsCmd := cobra.Command{
		Use:   "analytics",
		Short: "displays analytics",
		Long:  "displays analytics",
		Run: func(cmd *cobra.Command, args []string) {
			appInstance := models.NewAppInfo()
			reviewInstance := models.AppsReviewInfo()
			appList := utils.ReadFile[models.AppsInfo](constants.AppsInfoFile)
			appReviewList := utils.ReadFile[models.ReviewInfo](constants.AppsReviewInfoFile)
			// total apps
			totalApps := len(appList)
			// fmt.Println(totalApps)

			// app size
			appsSize := appInstance.ListSize(appList)
			// PrintAppsSize(appsSize)

			// no of generes
			generes := appInstance.UniqueGeneresNumber(appList)
			// fmt.Println(generes)

			//  no of category
			categories := appInstance.UniqueCategoriesNumber(appList)
			// fmt.Println(categories)

			// most reviewed apps
			mostReviewedApp := reviewInstance.MostReviewedApp(appReviewList)
			// fmt.Println(mostReviewedApp)

			// least reviewed apps
			leastReviwedApps := reviewInstance.LeastReviewedApp(appReviewList)
			// fmt.Println(leastReviwedApps)

			// most installed apps
			mostInstalledApp := appInstance.MostInstalledApp(appList)
			// fmt.Println(mostInstalledApp)

			// least installed apps
			leastInstalledApps := appInstance.LeastInstalledApp(appList)
			// fmt.Println(leastInstalledApps)

			// Total size of store
			totalSize := appInstance.TotalPlayStoreSize(appList)
			// fmt.Println(totalSize)

			// Total installs
			totalInstalls := appInstance.TotalInstalls(appList)
			// fmt.Println(totalInstalls)

			PrintAnalytics(totalApps, appsSize, generes, categories, mostReviewedApp, leastReviwedApps, mostInstalledApp, leastInstalledApps, totalSize, totalInstalls)
		},
	}
	return analyticsCmd
}

// Priting analytics
func PrintAnalytics(totalApps int, appsSize []float64, generes int, categories int, mostReviewedApp []string, leastReviwedApps []string, mostInstalledApp []string, leastInstalledApps []string, totalSize string, totalInstalls int) {
	p := message.NewPrinter(language.English)

	fmt.Println("Total Apps : ", totalApps)
	fmt.Println("Size of apps with count : ")
	PrintAppsSize(appsSize)
	fmt.Println("No. of Generes :", generes)
	fmt.Println("No. of Categories:", categories)
	fmt.Println("Most Reviwed Apps :", mostReviewedApp[0])
	fmt.Println("Least Reviwed Apps :", leastReviwedApps[0])
	fmt.Println("Most Installed Apps :", mostInstalledApp[0])
	fmt.Println("Least Installed Apps :", leastInstalledApps[0])
	fmt.Println("Total Size of Store :", totalSize, "GB")
	p.Println("Total Installs :", totalInstalls, "times")
}

// printing apps size
func PrintAppsSize(appsSize []float64) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"size > 50MB (%)", "size < 10MB (%)", "Varies with device (%)", "10MB < size <= 50MB(%)", "size (%) (KB)"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Append([]string{strconv.FormatFloat(appsSize[0], 'f', 2, 32), strconv.FormatFloat(appsSize[1], 'f', 2, 32), strconv.FormatFloat(appsSize[2], 'f', 2, 32), strconv.FormatFloat(appsSize[3], 'f', 2, 32), strconv.FormatFloat(appsSize[4], 'f', 2, 32)})
	table.Render()
}
