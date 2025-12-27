package main

import (
	handlerscontainerregistry "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/container-plugin/handlers/containerregistry"
	handlerskaas "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/container-plugin/handlers/kaas"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/health"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/server"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Aruba Cloud Container Plugin API for Krateo Operator Generator (KOG)
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
		Log: &log.Logger,
	}

	// Kaas
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Container/kaas", handlerskaas.PostKaas(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Container/kaas", handlerskaas.ListKaass(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Container/kaas/{id}", handlerskaas.GetKaas(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Container/kaas/{id}", handlerskaas.PutKaas(opts))

	// Containerregistry
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Container/registries", handlerscontainerregistry.PostContainerregistry(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Container/registries", handlerscontainerregistry.ListContainerregistrys(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Container/registries/{id}", handlerscontainerregistry.GetContainerregistry(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Container/registries/{id}", handlerscontainerregistry.PutContainerregistry(opts))

	// Swagger UI
	srv.Mux().Handle("/swagger/", httpSwagger.WrapHandler)

	// Kubernetes health check endpoints
	srv.Mux().HandleFunc("GET /healthz", health.LivenessHandler(srv.Healthy()))
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready()))

	srv.Run()
}
