package main

import (
	handlersbackup "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/storage-plugin/handlers/backup"
	handlersblockstorage "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/storage-plugin/handlers/blockstorage"
	handlersrestore "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/storage-plugin/handlers/restore"
	handlerssnapshot "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/storage-plugin/handlers/snapshot"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/health"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/server"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	srv := server.New()

	opts := handlers.HandlerOptions{
		Log: &log.Logger,
	}

	// Blockstorage
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Storage/blockstorages", handlersblockstorage.PostBlockstorage(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Storage/blockstorages", handlersblockstorage.ListBlockstorages(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Storage/blockstorages/{id}", handlersblockstorage.GetBlockstorage(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Storage/blockstorages/{id}", handlersblockstorage.PutBlockstorage(opts))

	// Snapshot
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Storage/snapshots", handlerssnapshot.PostSnapshot(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Storage/snapshots", handlerssnapshot.ListSnapshots(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Storage/snapshots/{id}", handlerssnapshot.GetSnapshot(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Storage/snapshots/{id}", handlerssnapshot.PutSnapshot(opts))

	// Backup
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Storage/backups", handlersbackup.PostBackup(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Storage/backups", handlersbackup.ListBackups(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Storage/backups/{id}", handlersbackup.GetBackup(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Storage/backups/{id}", handlersbackup.PutBackup(opts))

	// Restore
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Storage/restores", handlersrestore.PostRestore(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Storage/restores", handlersrestore.ListRestores(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Storage/restores/{id}", handlersrestore.GetRestore(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Storage/restores/{id}", handlersrestore.PutRestore(opts))

	// Swagger UI
	srv.Mux().Handle("/swagger/", httpSwagger.WrapHandler)

	// Kubernetes health check endpoints
	srv.Mux().HandleFunc("GET /healthz", health.LivenessHandler(srv.Healthy()))
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready()))

	srv.Run()
}
