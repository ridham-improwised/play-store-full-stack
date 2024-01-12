package cli

import (
	"errors"
	"os"
	"strconv"

	"github.com/Improwised/devempire-ridham-parmar-backend/constants"
	"github.com/Improwised/devempire-ridham-parmar-backend/models"
	"github.com/Improwised/devempire-ridham-parmar-backend/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	generes string
)

func ListAppsBySearchCmd() cobra.Command {
	searchAppsCmd := cobra.Command{
		Use:   "search",
		Short: "displays list of apps",
		Long:  "displays list of apps associated with passed argument",
		Example: `search 'instagram'
search --generes="communication`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			isSet := cmd.Flags().Lookup("generes").Changed

			switch {
			case len(args) > 0 && generes != "":
				return errors.New("expected either command or flag value")
			case generes != "":
				return nil
			case len(args) == 1 && !isSet:
				return nil
			case len(args) > 1:
				return errors.New("received two args, expected one")
			default:
				return errors.New("expected one arg")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			appsInstance := models.NewAppInfo()
			appList := utils.ReadFile[models.AppsInfo](constants.AppsInfoFile)

			if generes != "" {
				list := appsInstance.GetAppByGeneresPrice(appList, generes)
				sortedApps := appsInstance.SortByInstalls(list)
				apps := appsInstance.ListNames(sortedApps)
				PrintApps(apps, skip, limit)
				return
			}

			appsName := appsInstance.ListNames(appList)
			apps := appsInstance.ListAppsByArg(appsName, args[0])
			PrintApps(apps, skip, limit)
		},
	}
	searchAppsCmd.Flags().IntVarP(&skip, "skip", "s", 0, "skips specified records")
	searchAppsCmd.Flags().IntVarP(&limit, "limit", "l", 10, "displays specified records")
	return searchAppsCmd
}

// Prints apps by search args
func PrintApps(apps []string, skip, limit int) {
	printLimit := limit
	table := tablewriter.NewWriter(os.Stdout)
	if len(apps) > 0 && len(apps) > skip && skip >= 0 {
		table.SetHeader([]string{"Apps Name"})
		table.SetRowLine(true)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		for index, app := range apps {
			if index < skip {
				continue
			}
			if limit == 0 {
				break
			}
			table.Append([]string{app})
			limit--
		}
		table.Append([]string{"Total : " + strconv.Itoa(len(apps))})
		table.Append([]string{"skip : " + strconv.Itoa(skip)})
		table.Append([]string{"limit : " + strconv.Itoa(printLimit)})
	} else {
		table.SetFooter([]string{"No Records Founds"})
	}
	table.Render()
}
