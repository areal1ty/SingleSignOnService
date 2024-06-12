package app

import (
	appgrpc "SingleSignOnService/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *appgrpc.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	grpcApp := appgrpc.NewApp(log, grpcPort)
	return &App{
		GRPCServer: grpcApp,
	}
}
