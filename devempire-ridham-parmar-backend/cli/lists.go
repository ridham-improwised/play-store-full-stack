package cli

import (
	"os"
	"strconv"

	"github.com/Improwised/devempire-ridham-parmar-backend/constants"
	"github.com/Improwised/devempire-ridham-parmar-backend/models"
	"github.com/Improwised/devempire-ridham-parmar-backend/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var price bool

func ListByPriceCmd() cobra.Command {
	listByPrice := cobra.Command{
		Use:     "lists",
		Short:   "displays list of paid apps",
		Long:    "displays list of paid apps",
		Example: "lists --price",
		Run: func(cmd *cobra.Command, args []string) {
			appsInstance := models.NewAppInfo()

			appList := utils.ReadFile[models.AppsInfo](constants.AppsInfoFile)

			list := appsInstance.ListAppByPrice(appList)

			sortedApps := appsInstance.SortByInstalls(list)

			PrintPriceWiseMostLikedApps(sortedApps, skip, limit)
		},
	}
	listByPrice.Flags().IntVarP(&skip, "skip", "s", 0, "skips specified records")
	listByPrice.Flags().IntVarP(&limit, "limit", "l", 10, "displays specified records")
	return listByPrice
}

func PrintPriceWiseMostLikedApps(apps []models.AppsInfo, skip, limit int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	printLimit := limit
	if len(apps) > 0 && len(apps) > skip && skip >= 0 {
		table.SetRowLine(true)
		table.SetHeader([]string{"App Name", "Price"})
		for index, apps := range apps {
			if index < skip {
				continue
			}
			if limit == 0 {
				break
			}
			table.Append([]string{apps.AppName, apps.Price})
			limit--
		}
		table.Append([]string{"Total : ", strconv.Itoa(len(apps))})
		table.Append([]string{"skip : ", strconv.Itoa(skip)})
		table.Append([]string{"limit : ", strconv.Itoa(printLimit)})
	} else {
		table.SetFooter([]string{"No Records Founds"})
		return
	}

	table.Render()
}
