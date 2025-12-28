package main

import (
	handlerskms "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/security-plugin/handlers/kms"
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

	// Kms
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Security/kms", handlerskms.PostKms(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Security/kms", handlerskms.ListKmss(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Security/kms/{kmsId}", handlerskms.GetKms(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Security/kms/{kmsId}", handlerskms.PutKms(opts))
	srv.Mux().Handle("DELETE /projects/{projectId}/providers/Aruba.Security/kms/{kmsId}", handlerskms.DeleteKms(opts))

	// Swagger UI
	srv.Mux().Handle("/swagger/", httpSwagger.WrapHandler)

	// Kubernetes health check endpoints
	srv.Mux().HandleFunc("GET /healthz", health.LivenessHandler(srv.Healthy()))
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready()))

	srv.Run()
}
