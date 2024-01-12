package cli

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/Improwised/devempire-ridham-parmar-backend/config"
	"github.com/Improwised/devempire-ridham-parmar-backend/database"
	"github.com/Improwised/devempire-ridham-parmar-backend/pkg/events"
	"github.com/Improwised/devempire-ridham-parmar-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"

	pMetrics "github.com/Improwised/devempire-ridham-parmar-backend/pkg/prometheus"
)

// GetAPICommandDef runs app
func GetAPICommandDef(cfg config.AppConfig, logger *zap.Logger) cobra.Command {
	apiCommand := cobra.Command{
		Use:   "api",
		Short: "To start api",
		Long:  `To start api`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// Create fiber app
			app := fiber.New(fiber.Config{})

			app.Use(cors.New())

			promMetrics := pMetrics.InitPrometheusMetrics()

			// Init eventbus
			events := events.NewEventBus(logger)

			db, err := database.Connect(cfg.DB)
			if err != nil {
				return err
			}

			err = events.SubscribeAll()
			if err != nil {
				return err
			}

			// setup routes
			err = routes.Setup(app, db, logger, cfg, events, promMetrics)
			if err != nil {
				return err
			}

			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
			go func() {
				if err := app.Listen(cfg.Port); err != nil {
					logger.Panic(err.Error())
				}
			}()

			<-interrupt
			logger.Info("gracefully shutting down...")
			if err := app.Shutdown(); err != nil {
				logger.Panic("error while shutdown server", zap.Error(err))
			}

			logger.Info("server stopped to receive new requests or connection.")
			return nil
		},
	}

	return apiCommand
}