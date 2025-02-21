package run

import (
	"Library/config"
	"Library/internal/dbase"
	"Library/internal/models"
	"Library/internal/modules"
	"Library/internal/responder"
	"Library/internal/router"
	"Library/internal/server"
	"Library/internal/token"
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/ptflp/godecoder"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Runner interface {
	Run()
}

type App struct {
	conf         config.AppConfig
	logger       *zap.Logger
	srv          server.Server
	Repositories *modules.Repositories
	Services     *modules.Services
}

func NewApp(conf config.AppConfig, logger *zap.Logger) *App {
	return &App{conf: conf, logger: logger}
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		a.logger.Info("Shutting down gracefully...")
		cancel()
	}()

	if err := a.srv.Serve(ctx); err != nil {
		a.logger.Error("Server error", zap.Error(err))
		os.Exit(1)
	}

	a.logger.Info("Application stopped")
}

func (a *App) Bootstrap() Runner {
	decoder := godecoder.NewDecoder(jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		DisallowUnknownFields:  true,
	})

	db, err := dbase.NewPostgersDB(&a.conf)
	if err != nil {
		a.logger.Error("Error initializing PostgersDB: %v", zap.Error(err))
	}

	//if err = db.Migrator().DropTable(&models.Book{}, &models.Author{}, &models.User{}); err != nil {
	//	a.logger.Error("Error dropping tables: %v", zap.Error(err))
	//}

	if err = db.AutoMigrate(&models.Book{}, &models.Author{}, &models.User{}); err != nil {
		a.logger.Error("Error migrating tables: %v", zap.Error(err))
	}

	responseManager := responder.NewResponder(decoder, a.logger)

	tokenManager := token.NewJWTTokenService(a.conf.Token.Key)

	NewRepositories := modules.NewRepositories(db)
	a.Repositories = NewRepositories

	services := modules.NewServices(NewRepositories, a.logger, tokenManager)
	a.Services = services

	isEmpty, err := services.IsEmpty()
	if err != nil {
		a.logger.Error("Error checking db for emptiness: %v", zap.Error(err))
	}
	if isEmpty {
		if err = a.Services.FillBD(10, 50, 100); err != nil {
			a.logger.Error("Error filling db: %v", zap.Error(err))
		}
	}

	controllers := modules.NewControllers(services, decoder, responseManager)

	initHandlers := router.NewApiRouter(controllers, tokenManager)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.conf.Server.Port),
		Handler: initHandlers,
	}

	a.srv = server.NewHttpServer(a.conf.Server, a.logger, srv)

	return a
}
