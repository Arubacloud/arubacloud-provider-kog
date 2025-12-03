package main

import (
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/pkg/handlers"
	"github.com/Arubacloud/arubacloud-provider-kog/pkg/health"
	"github.com/Arubacloud/arubacloud-provider-kog/pkg/server"
	handlersblockstorage "github.com/Arubacloud/arubacloud-provider-kog/storage-plugin/handlers/blockstorage"
	handlerssnapshot "github.com/Arubacloud/arubacloud-provider-kog/storage-plugin/handlers/snapshot"
	handlersbackup "github.com/Arubacloud/arubacloud-provider-kog/storage-plugin/handlers/backup"
	handlersrestore "github.com/Arubacloud/arubacloud-provider-kog/storage-plugin/handlers/restore"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Aruba Cloud Storage Plugin API for Krateo Operator Generator (KOG)
// @version         1.0
// @description     Simple wrapper around Aruba Cloud API to provide consistency of API response for Krateo Operator Generator (KOG)
// @termsOfService  http://swagger.io/terms/
// @contact.name    Krateo Support
// @contact.url     https://krateo.io
// @contact.email   contact@krateoplatformops.io
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8080
// @BasePath        /
// @schemes         http
func main() {
	srv := server.New()

	opts := handlers.HandlerOptions{
		Log:    &log.Logger,
		Client: http.DefaultClient,
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
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready(), opts.Client.(*http.Client)))

	srv.Run()
}
