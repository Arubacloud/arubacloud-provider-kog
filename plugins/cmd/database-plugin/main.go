package main

import (
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/pkg/handlers"
	"github.com/Arubacloud/arubacloud-provider-kog/pkg/health"
	"github.com/Arubacloud/arubacloud-provider-kog/pkg/server"
	handlersdbaas "github.com/Arubacloud/arubacloud-provider-kog/database-plugin/handlers/dbaas"
	handlersdatabase "github.com/Arubacloud/arubacloud-provider-kog/database-plugin/handlers/database"
	handlersuser "github.com/Arubacloud/arubacloud-provider-kog/database-plugin/handlers/user"
	handlersgrant "github.com/Arubacloud/arubacloud-provider-kog/database-plugin/handlers/grant"
	handlersbackup "github.com/Arubacloud/arubacloud-provider-kog/database-plugin/handlers/backup"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Aruba Cloud Database Plugin API for Krateo Operator Generator (KOG)
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
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready(), opts.Client.(*http.Client)))

	srv.Run()
}
