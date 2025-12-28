package main

import (
	handlerselasticip "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/network-plugin/handlers/elasticip"
	handlersloadbalancer "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/network-plugin/handlers/loadbalancer"
	handlerssecuritygroup "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/network-plugin/handlers/securitygroup"
	handlerssecurityrule "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/network-plugin/handlers/securityrule"
	handlerssubnet "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/network-plugin/handlers/subnet"
	handlersvpc "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/network-plugin/handlers/vpc"
	handlersvpcpeering "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/network-plugin/handlers/vpcpeering"
	handlersvpcpeeringroute "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/network-plugin/handlers/vpcpeeringroute"
	handlersvpntunnel "github.com/Arubacloud/arubacloud-provider-kog/plugins/cmd/network-plugin/handlers/vpntunnel"
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
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/securitygroups/{securityGroupId}/securityrules", handlerssecurityrule.PostSecurityrule(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/securitygroups/{securityGroupId}/securityrules", handlerssecurityrule.ListSecurityrules(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/securitygroups/{securityGroupId}/securityrules/{id}", handlerssecurityrule.GetSecurityrule(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/securitygroups/{securityGroupId}/securityrules/{id}", handlerssecurityrule.PutSecurityrule(opts))

	// Elasticip
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/elasticips", handlerselasticip.PostElasticip(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/elasticips", handlerselasticip.ListElasticips(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/elasticips/{id}", handlerselasticip.GetElasticip(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/elasticips/{id}", handlerselasticip.PutElasticip(opts))

	// Loadbalancer (read-only: only GET and LIST are supported)
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/loadbalancers", handlersloadbalancer.ListLoadbalancers(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/loadbalancers/{id}", handlersloadbalancer.GetLoadbalancer(opts))

	// Vpntunnel
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/vpntunnels", handlersvpntunnel.PostVpntunnel(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpntunnels", handlersvpntunnel.ListVpntunnels(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpntunnels/{id}", handlersvpntunnel.GetVpntunnel(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/vpntunnels/{id}", handlersvpntunnel.PutVpntunnel(opts))

	// Vpcpeering
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/vpcPeerings", handlersvpcpeering.PostVPCPeering(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/vpcPeerings", handlersvpcpeering.ListVPCPeerings(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/vpcPeerings/{id}", handlersvpcpeering.GetVPCPeering(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/vpcPeerings/{id}", handlersvpcpeering.PutVPCPeering(opts))
	srv.Mux().Handle("DELETE /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/vpcPeerings/{id}", handlersvpcpeering.DeleteVPCPeering(opts))

	// Vpcpeeringroute
	srv.Mux().Handle("POST /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/vpcPeerings/{vpcPeeringId}/vpcPeeringRoutes", handlersvpcpeeringroute.PostVPCPeeringRoute(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/vpcPeerings/{vpcPeeringId}/vpcPeeringRoutes", handlersvpcpeeringroute.ListVPCPeeringRoutes(opts))
	srv.Mux().Handle("GET /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/vpcPeerings/{vpcPeeringId}/vpcPeeringRoutes/{id}", handlersvpcpeeringroute.GetVPCPeeringRoute(opts))
	srv.Mux().Handle("PUT /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/vpcPeerings/{vpcPeeringId}/vpcPeeringRoutes/{id}", handlersvpcpeeringroute.PutVPCPeeringRoute(opts))
	srv.Mux().Handle("DELETE /projects/{projectId}/providers/Aruba.Network/vpcs/{vpcId}/vpcPeerings/{vpcPeeringId}/vpcPeeringRoutes/{id}", handlersvpcpeeringroute.DeleteVPCPeeringRoute(opts))

	// Swagger UI
	srv.Mux().Handle("/swagger/", httpSwagger.WrapHandler)

	// Kubernetes health check endpoints
	srv.Mux().HandleFunc("GET /healthz", health.LivenessHandler(srv.Healthy()))
	srv.Mux().HandleFunc("GET /readyz", health.ReadinessHandler(srv.Ready()))

	srv.Run()
}
