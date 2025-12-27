package main

import (
	handlerscloudserver "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/compute-plugin/handlers/cloudserver"
	handlerskeypair "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/compute-plugin/handlers/keypair"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/health"
	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/server"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Aruba Cloud Compute Plugin API for Krateo Operator Generator (KOG)
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

	// Cloudserver
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Compute/cloudservers", handlerscloudserver.PostCloudserver(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Compute/cloudservers", handlerscloudserver.ListCloudservers(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Compute/cloudservers/{id}", handlerscloudserver.GetCloudserver(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Compute/cloudservers/{id}", handlerscloudserver.PutCloudserver(opts))

	// Keypair
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Compute/keypairs", handlerskeypair.PostKeypair(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Compute/keypairs", handlerskeypair.ListKeypairs(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Compute/keypairs/{id}", handlerskeypair.GetKeypair(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Compute/keypairs/{id}", handlerskeypair.PutKeypair(opts))

	// Swagger UI
	srv.Mux().Handle("/swagger/", httpSwagger.WrapHandler)

	// Kubernetes health check endpoints
	srv.Mux().HandleFunc("GET /healthz", health.LivenessHandler(srv.Healthy()))
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready()))

	srv.Run()
}
