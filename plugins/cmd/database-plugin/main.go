package main

import (
	handlersbackup "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/database-plugin/handlers/backup"
	handlersdatabase "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/database-plugin/handlers/database"
	handlersdbaas "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/database-plugin/handlers/dbaas"
	handlersgrant "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/database-plugin/handlers/grant"
	handlersuser "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/database-plugin/handlers/user"
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

	// Dbaas
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Database/dbaas", handlersdbaas.PostDbaas(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Database/dbaas", handlersdbaas.ListDbaass(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Database/dbaas/{id}", handlersdbaas.GetDbaas(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Database/dbaas/{id}", handlersdbaas.PutDbaas(opts))

	// Database
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/databases", handlersdatabase.PostDatabase(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/databases", handlersdatabase.ListDatabases(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/databases/{id}", handlersdatabase.GetDatabase(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/databases/{id}", handlersdatabase.PutDatabase(opts))

	// User
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/users", handlersuser.PostUser(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/users", handlersuser.ListUsers(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/users/{id}", handlersuser.GetUser(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/users/{id}", handlersuser.PutUser(opts))

	// Grant
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/grants", handlersgrant.PostGrant(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/grants", handlersgrant.ListGrants(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/grants/{id}", handlersgrant.GetGrant(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/grants/{id}", handlersgrant.PutGrant(opts))

	// Backup
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/backups", handlersbackup.PostBackup(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/backups", handlersbackup.ListBackups(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/backups/{id}", handlersbackup.GetBackup(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Database/dbaas/{dbaasId}/backups/{id}", handlersbackup.PutBackup(opts))

	// Swagger UI
	srv.Mux().Handle("/swagger/", httpSwagger.WrapHandler)

	// Kubernetes health check endpoints
	srv.Mux().HandleFunc("GET /healthz", health.LivenessHandler(srv.Healthy()))
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready()))

	srv.Run()
}
