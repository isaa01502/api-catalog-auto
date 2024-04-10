package http

import (
	"api-catalog-auto/config"
	"api-catalog-auto/internal/common/logger"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start() (serverChannel chan error)
	Stop() error
}

type server struct {
	router      *gin.Engine
	srv         *http.Server
	srvCh       chan error
	stopTimeout time.Duration
	handlers    *Handlers
}

func (s *server) Start() (serverChannel chan error) {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.srvCh <- err
		}
	}()

	return s.srvCh
}

func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.stopTimeout)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown : %v", err)
	}
	return nil
}

func New(cfg *config.Config, handlers *Handlers, log logger.AppLogger) (s Server, err error) {
	if cfg.Http.Port <= 0 {
		return nil, fmt.Errorf("bad port value %v", cfg.Http.Port)
	}
	if cfg.Http.StopTimeout < 1 {
		return nil, fmt.Errorf("bad stop timeout value %v", cfg.Http.StopTimeout)
	}

	if cfg.Http.Gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()

	if cfg.Http.Gin.UseRecovery {
		engine.Use(gin.CustomRecovery(log.HttpPanicHandler))
	}

	srv := &server{
		router: engine,
		srv: &http.Server{
			Addr:    ":" + strconv.Itoa(cfg.Http.Port),
			Handler: engine,
		},
		handlers:    handlers,
		stopTimeout: time.Duration(cfg.Http.StopTimeout) * time.Millisecond,
	}

	srv.addSwaggerSettings(&cfg.SwaggerUI)
	srv.routers()

	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		return
	}

	//if cfg.Http.ProfilingEnabled {
	//	pprof.Register(engine)
	//}

	return srv, nil
}
