package main

import (
	database "akozadaev/go_reports/db"
	"akozadaev/go_reports/db/authentication"
	"akozadaev/go_reports/db/report"
	"akozadaev/go_reports/middleware"
	config "akozadaev/go_reports/pkg"
	"akozadaev/go_reports/server"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"log"
	"net/http"
	"os"
	"time"
)

var serverCmd = &cobra.Command{
	Use: "server:start",
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func init() {
	envPath := "."
	envFileName := ".env"

	fullPath := envPath + "/" + envFileName

	if err := godotenv.Overload(fullPath); err != nil {
		log.Printf("[ERROR] failed with %+v", "No .env file found")
	}
}

func main() {
	if err := serverCmd.Execute(); err != nil {
		log.Printf("failed to execute command. err: %v", err)
		os.Exit(1)
	}
}

/*
	func main() {
		dbName := os.Getenv("DB_NAME")
		db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		database.Migrate(db)
		server.Start(db)
	}
*/
func runApplication() {
	serverConfig, err := config.Load()
	if err != nil {
	}

	app := fx.New(
		fx.Supply(serverConfig),
		fx.StopTimeout(serverConfig.ServerConfig.GracefulShutdown+time.Second),
		fx.Provide(
			// setup database
			database.NewDatabase,
			// setup auth
			authentication.NewAuthRepository,
			server.NewHandler,
			report.NewReportRepository,
			newServer,
		),
		fx.Invoke(
			server.RouteV1,
			//datagenerator.RouteV1, //TODO
			func(r *gin.Engine) {},
		),
	)
	app.Run()
}

func newServer(lc fx.Lifecycle, cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	r.Use(middleware.TimeoutMiddleware(cfg.ServerConfig.WriteTimeout))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		Handler:      r,
		ReadTimeout:  cfg.ServerConfig.ReadTimeout,
		WriteTimeout: cfg.ServerConfig.WriteTimeout,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return r
}
