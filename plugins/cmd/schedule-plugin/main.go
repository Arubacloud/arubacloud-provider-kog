package main

import (
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/health"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/server"
	handlersjob "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/schedule-plugin/handlers/job"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Aruba Cloud Schedule Plugin API for Krateo Operator Generator (KOG)
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

	// Job
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Schedule/jobs", handlersjob.PostJob(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Schedule/jobs", handlersjob.ListJobs(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Schedule/jobs/{id}", handlersjob.GetJob(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Schedule/jobs/{id}", handlersjob.PutJob(opts))

	// Swagger UI
	srv.Mux().Handle("/swagger/", httpSwagger.WrapHandler)

	// Kubernetes health check endpoints
	srv.Mux().HandleFunc("GET /healthz", health.LivenessHandler(srv.Healthy()))
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready(), opts.Client.(*http.Client)))

	srv.Run()
}
