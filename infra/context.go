package infra

import (
	"20241209/config"
	"20241209/database"
	"20241209/handler"
	"20241209/log"
	"20241209/middleware"
	"20241209/repository"
	"20241209/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cacher     database.Cacher
	Cfg        config.Config
	Ctl        handler.Handler
	Log        *zap.Logger
	Middleware middleware.Middleware
}

func NewServiceContext(migrateDb bool, seedDb bool) (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	appConfig, err := config.LoadConfig(migrateDb, seedDb)
	if err != nil {
		return handlerError(err)
	}

	// instance logger
	logger, err := log.InitZapLogger(appConfig)
	if err != nil {
		return handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(appConfig)
	if err != nil {
		return handlerError(err)
	}

	rdb := database.NewCacher(appConfig, 60*60)

	// instance repository
	repo := repository.NewRepository(db, rdb, appConfig, logger)

	// instance service
	services := service.NewService(repo)

	// instance controller
	Ctl := handler.NewHandler(services, logger)

	mw := middleware.NewMiddleware(rdb, appConfig.AppSecret)

	return &ServiceContext{Cacher: rdb, Cfg: appConfig, Ctl: *Ctl, Log: logger, Middleware: mw}, nil
}
