package main

import (
	handlersjob "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/schedule-plugin/handlers/job"
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

	// Job
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Schedule/jobs", handlersjob.PostJob(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Schedule/jobs", handlersjob.ListJobs(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Schedule/jobs/{id}", handlersjob.GetJob(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Schedule/jobs/{id}", handlersjob.PutJob(opts))

	// Swagger UI
	srv.Mux().Handle("/swagger/", httpSwagger.WrapHandler)

	// Kubernetes health check endpoints
	srv.Mux().HandleFunc("GET /healthz", health.LivenessHandler(srv.Healthy()))
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready()))

	srv.Run()
}
