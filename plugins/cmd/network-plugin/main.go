package main

import (
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/pkg/handlers"
	"github.com/Arubacloud/arubacloud-provider-kog/pkg/health"
	"github.com/Arubacloud/arubacloud-provider-kog/pkg/server"
	handlersvpc "github.com/Arubacloud/arubacloud-provider-kog/network-plugin/handlers/vpc"
	handlerssubnet "github.com/Arubacloud/arubacloud-provider-kog/network-plugin/handlers/subnet"
	handlerssecuritygroup "github.com/Arubacloud/arubacloud-provider-kog/network-plugin/handlers/securitygroup"
	handlerssecurityrule "github.com/Arubacloud/arubacloud-provider-kog/network-plugin/handlers/securityrule"
	handlerselasticip "github.com/Arubacloud/arubacloud-provider-kog/network-plugin/handlers/elasticip"
	handlersloadbalancer "github.com/Arubacloud/arubacloud-provider-kog/network-plugin/handlers/loadbalancer"
	handlersvpntunnel "github.com/Arubacloud/arubacloud-provider-kog/network-plugin/handlers/vpntunnel"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Aruba Cloud Network Plugin API for Krateo Operator Generator (KOG)
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

	// Vpc
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/vpcs", handlersvpc.PostVpc(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpcs", handlersvpc.ListVpcs(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpcs/{id}", handlersvpc.GetVpc(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/vpcs/{id}", handlersvpc.PutVpc(opts))

	// Subnet
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/subnets", handlerssubnet.PostSubnet(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/subnets", handlerssubnet.ListSubnets(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/subnets/{id}", handlerssubnet.GetSubnet(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/subnets/{id}", handlerssubnet.PutSubnet(opts))

	// Securitygroup
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/securitygroups", handlerssecuritygroup.PostSecuritygroup(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/securitygroups", handlerssecuritygroup.ListSecuritygroups(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/securitygroups/{id}", handlerssecuritygroup.GetSecuritygroup(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/securitygroups/{id}", handlerssecuritygroup.PutSecuritygroup(opts))

	// Securityrule
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/securitygroups/{securityGroupId}/rules", handlerssecurityrule.PostSecurityrule(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/securitygroups/{securityGroupId}/rules", handlerssecurityrule.ListSecurityrules(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/securitygroups/{securityGroupId}/rules/{id}", handlerssecurityrule.GetSecurityrule(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/securitygroups/{securityGroupId}/rules/{id}", handlerssecurityrule.PutSecurityrule(opts))

	// Elasticip
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/elasticips", handlerselasticip.PostElasticip(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/elasticips", handlerselasticip.ListElasticips(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/elasticips/{id}", handlerselasticip.GetElasticip(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/elasticips/{id}", handlerselasticip.PutElasticip(opts))

	// Loadbalancer
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/loadbalancers", handlersloadbalancer.PostLoadbalancer(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/loadbalancers", handlersloadbalancer.ListLoadbalancers(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/loadbalancers/{id}", handlersloadbalancer.GetLoadbalancer(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/loadbalancers/{id}", handlersloadbalancer.PutLoadbalancer(opts))

	// Vpntunnel
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/vpntunnels", handlersvpntunnel.PostVpntunnel(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpntunnels", handlersvpntunnel.ListVpntunnels(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpntunnels/{id}", handlersvpntunnel.GetVpntunnel(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/vpntunnels/{id}", handlersvpntunnel.PutVpntunnel(opts))

	// Swagger UI
	srv.Mux().Handle("/swagger/", httpSwagger.WrapHandler)

	// Kubernetes health check endpoints
	srv.Mux().HandleFunc("GET /healthz", health.LivenessHandler(srv.Healthy()))
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready(), opts.Client.(*http.Client)))

	srv.Run()
}
