package di

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/DrDowellsHead/Go_Bootcamp_4/src/datasource"
	"github.com/DrDowellsHead/Go_Bootcamp_4/src/domain"
	"github.com/DrDowellsHead/Go_Bootcamp_4/src/web"
	"go.uber.org/fx"
)

// Module собирает зависимости; хранилище используется одно на всё приложение.
var Module = fx.Options(
	fx.Provide(
		datasource.NewStorage,
		fx.Annotate(datasource.NewRepository, fx.As(new(domain.Repository))),
		domain.NewService,
		web.NewHandler,
		newServer,
	),
	fx.Invoke(registerServer),
)

func newServer(handler *web.Handler) *http.Server {
	return &http.Server{
		Addr:              ":8080",
		Handler:           handler.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func registerServer(lifecycle fx.Lifecycle, server *http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			listener, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			go func() {
				_ = server.Serve(listener)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
