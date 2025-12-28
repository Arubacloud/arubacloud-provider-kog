package main

import (
	handlersproject "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/project-plugin/handlers/project"
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

	// Project
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Project/projects", handlersproject.PostProject(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Project/projects", handlersproject.ListProjects(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Project/projects/{id}", handlersproject.GetProject(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Project/projects/{id}", handlersproject.PutProject(opts))

	// Swagger UI
	srv.Mux().Handle("/swagger/", httpSwagger.WrapHandler)

	// Kubernetes health check endpoints
	srv.Mux().HandleFunc("GET /healthz", health.LivenessHandler(srv.Healthy()))
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready()))

	srv.Run()
}
