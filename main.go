package main

import (
	"api-catalog-auto/config"
	"api-catalog-auto/internal/adapter/repository"
	"api-catalog-auto/internal/common/logger"
	"api-catalog-auto/internal/core/service"
	"api-catalog-auto/internal/http"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Init("config.json")
	if err != nil {
		panic("config init error")
	}

	db, err := repository.NewDB(cfg.DB)
	if err != nil {
		panic("init db error")
	}

	templageRep, err := repository.NewTemplate(db)
	if err != nil {
		panic("init auto_catalog mock error")
	}

	log := logger.New()

	catalog := service.NewCatalog(cfg, templageRep)
	handlers := http.NewHandlers(catalog)

	srvNew, err := http.New(cfg, handlers, log)

	startSrvErr := srvNew.Start()

	var stopReason string
	select {
	case err = <-startSrvErr:
		stopReason = fmt.Sprintf("start server error %v", err)
	case qs := <-quit:
		stopReason = fmt.Sprintf("receiving signal %v", qs)
	}
	fmt.Printf("stop reason = %v", stopReason)

	err = db.CloseDB()
	if err != nil {
		fmt.Printf("stop db error %v", err)
		return
	}

	err = srvNew.Stop()
	if err != nil {
		fmt.Printf("stop server error %v", err)
		return
	}
	fmt.Println("server stopped")
}
