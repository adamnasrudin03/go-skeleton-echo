package main

import (
	"fmt"
	"time"

	help "github.com/adamnasrudin03/go-helpers"
	"github.com/adamnasrudin03/go-skeleton-echo/app"
	"github.com/adamnasrudin03/go-skeleton-echo/app/configs"
	"github.com/adamnasrudin03/go-skeleton-echo/app/router"
	"github.com/adamnasrudin03/go-skeleton-echo/pkg/database"
	"github.com/adamnasrudin03/go-skeleton-echo/pkg/driver"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func init() {
	// set timezone local
	loc, _ := time.LoadLocation(help.AsiaJakarta)
	time.Local = loc

}

func main() {
	var (
		cfg                  = configs.GetInstance()
		logger               = driver.Logger(cfg)
		cache                = driver.Redis(cfg)
		validate             = validator.New()
		db          *gorm.DB = database.SetupDbConnection(cfg, logger)
		repo                 = app.WiringRepository(db, &cache, cfg, logger)
		services             = app.WiringService(repo, cfg, logger)
		controllers          = app.WiringController(services, cfg, logger, validate)
	)

	defer database.CloseDbConnection(db, logger)

	r := router.NewRoutes()

	controllers.TeamMember.Mount(r.HttpServer.Group("/v1/team-members"))

	listen := fmt.Sprintf(":%v", cfg.App.Port)
	r.Run(listen)
}
