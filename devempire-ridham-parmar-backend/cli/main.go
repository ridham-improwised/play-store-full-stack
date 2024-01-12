package cli

import (
	"fmt"

	"github.com/Improwised/devempire-ridham-parmar-backend/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Init app initialization
func Init(cfg config.AppConfig, logger *zap.Logger) error {
	migrationCmd := GetMigrationCommandDef(cfg)
	apiCmd := GetAPICommandDef(cfg, logger)
	appsCmd := ListAppsCmd()
	searchAppsCmd := ListAppsBySearchCmd()
	listAppByPriceCmd := ListByPriceCmd()
	analyticsCmd := GetAnalyticsCmd()

	rootCmd := &cobra.Command{
		Use:   "apps-store",
		Short: "This is root command",
		Long:  "This is root command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Apps Store !!!")
		},
	}
	rootCmd.AddCommand(&migrationCmd, &apiCmd, &appsCmd, &searchAppsCmd, &listAppByPriceCmd)
	rootCmd.AddCommand(&migrationCmd, &apiCmd, &appsCmd, &searchAppsCmd, &analyticsCmd)
	appsCmd.Flags().StringVarP(&appName, "appName", "a", "", "mention app name (required)")
	searchAppsCmd.Flags().StringVarP(&generes, "generes", "g", "", "mention generes")
	listAppByPriceCmd.Flags().BoolVarP(&price, "price", "p", false, "set price flag (required)")
	err := listAppByPriceCmd.MarkFlagRequired("price")
	if err != nil {
		panic("price flag is not set")
	}
	return rootCmd.Execute()
}
