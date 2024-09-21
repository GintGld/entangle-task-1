package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	controller "task1/internal/controller/NDL"
	"task1/internal/lib/utils/sl"
	nglSrv "task1/internal/service/NGL"
)

type App struct {
	log    *slog.Logger
	addr   string
	router *mux.Router

	wg     *sync.WaitGroup
	server *http.Server
}

func New(
	log *slog.Logger,
	port int,
	client nglSrv.NGLClient,
) *App {
	// Init service.
	ngl := nglSrv.New(
		log,
		client,
	)

	// Init router.
	router := controller.New(
		log,
		ngl,
	)

	return &App{
		log:    log,
		addr:   fmt.Sprintf(":%d", port),
		router: router,
		wg:     &sync.WaitGroup{},
	}
}

// Run start http server.
func (a *App) Run() {
	srv := &http.Server{
		Addr:    a.addr,
		Handler: a.router,
	}

	a.wg.Add(1)
	defer a.wg.Done()

	// always returns error. ErrServerClosed on graceful close
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		a.log.Error("failed to run ListendAndServe", sl.Err(err))
	}
}

// Stop stops server.
func (a *App) Stop() {
	if err := a.server.Shutdown(context.Background()); err != nil {
		a.log.Error("failed to shut down server", sl.Err(err))
	}

	a.wg.Wait()
}
