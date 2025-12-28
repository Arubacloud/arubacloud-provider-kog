package vpcpeeringroute

import (
	"encoding/json"
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/sdk-go/pkg/types"
)

func GetVPCPeeringRoute(opts handlers.HandlerOptions) handlers.Handler {
	return &getHandler{baseHandler: newBaseHandler(opts)}
}

func PostVPCPeeringRoute(opts handlers.HandlerOptions) handlers.Handler {
	return &postHandler{baseHandler: newBaseHandler(opts)}
}

func PutVPCPeeringRoute(opts handlers.HandlerOptions) handlers.Handler {
	return &putHandler{baseHandler: newBaseHandler(opts)}
}

func DeleteVPCPeeringRoute(opts handlers.HandlerOptions) handlers.Handler {
	return &deleteHandler{baseHandler: newBaseHandler(opts)}
}

func ListVPCPeeringRoutes(opts handlers.HandlerOptions) handlers.Handler {
	return &listHandler{baseHandler: newBaseHandler(opts)}
}

// Interface compliance verification
var _ handlers.Handler = &getHandler{}
var _ handlers.Handler = &postHandler{}
var _ handlers.Handler = &putHandler{}
var _ handlers.Handler = &deleteHandler{}
var _ handlers.Handler = &listHandler{}

// Base handler with common functionality
type baseHandler struct {
	handlers.HandlerOptions
}

// Constructor for the base handler
func newBaseHandler(opts handlers.HandlerOptions) *baseHandler {
	return &baseHandler{HandlerOptions: opts}
}

// Handler types embedding the base handler
type getHandler struct {
	*baseHandler
}

type postHandler struct {
	*baseHandler
}

type putHandler struct {
	*baseHandler
}

type deleteHandler struct {
	*baseHandler
}

type listHandler struct {
	*baseHandler
}

// ServeHTTP implementation for GET handler
func (h *getHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")
	vpcId := r.PathValue("vpcId")
	vpcPeeringId := r.PathValue("vpcPeeringId")
	id := r.PathValue("id")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
		return
	}
	if vpcPeeringId == "" {
		http.Error(w, "vpcPeeringId is required", http.StatusBadRequest)
		return
	}
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Getting VPC peering route %s for vpc peering %s in vpc %s in project %s", id, vpcPeeringId, vpcId, projectId)

	// Call Aruba Cloud SDK to get VPC peering route
	response, err := client.FromNetwork().VPCPeeringRoutes().Get(r.Context(), projectId, vpcId, vpcPeeringId, id, params)
	if err != nil {
		h.Log.Printf("Failed to get VPC peering route: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if err := json.NewEncoder(w).Encode(response.Data); err != nil {
		h.Log.Printf("Failed to encode response: %v", err)
	}
}

// ServeHTTP implementation for POST handler
func (h *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")
	vpcId := r.PathValue("vpcId")
	vpcPeeringId := r.PathValue("vpcPeeringId")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
		return
	}
	if vpcPeeringId == "" {
		http.Error(w, "vpcPeeringId is required", http.StatusBadRequest)
		return
	}

	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Decode request body
	var req types.VPCPeeringRouteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Creating VPC peering route for vpc peering %s in vpc %s in project %s", vpcPeeringId, vpcId, projectId)

	// Call Aruba Cloud SDK to create VPC peering route
	response, err := client.FromNetwork().VPCPeeringRoutes().Create(r.Context(), projectId, vpcId, vpcPeeringId, req, params)
	if err != nil {
		h.Log.Printf("Failed to create VPC peering route: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if err := json.NewEncoder(w).Encode(response.Data); err != nil {
		h.Log.Printf("Failed to encode response: %v", err)
	}
}

// ServeHTTP implementation for PUT handler
func (h *putHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")
	vpcId := r.PathValue("vpcId")
	vpcPeeringId := r.PathValue("vpcPeeringId")
	id := r.PathValue("id")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
		return
	}
	if vpcPeeringId == "" {
		http.Error(w, "vpcPeeringId is required", http.StatusBadRequest)
		return
	}
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Decode request body
	var req types.VPCPeeringRouteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Updating VPC peering route %s for vpc peering %s in vpc %s in project %s", id, vpcPeeringId, vpcId, projectId)

	// Call Aruba Cloud SDK to update VPC peering route
	response, err := client.FromNetwork().VPCPeeringRoutes().Update(r.Context(), projectId, vpcId, vpcPeeringId, id, req, params)
	if err != nil {
		h.Log.Printf("Failed to update VPC peering route: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if err := json.NewEncoder(w).Encode(response.Data); err != nil {
		h.Log.Printf("Failed to encode response: %v", err)
	}
}

// ServeHTTP implementation for DELETE handler
func (h *deleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")
	vpcId := r.PathValue("vpcId")
	vpcPeeringId := r.PathValue("vpcPeeringId")
	id := r.PathValue("id")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
		return
	}
	if vpcPeeringId == "" {
		http.Error(w, "vpcPeeringId is required", http.StatusBadRequest)
		return
	}
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Deleting VPC peering route %s for vpc peering %s in vpc %s in project %s", id, vpcPeeringId, vpcId, projectId)

	// Call Aruba Cloud SDK to delete VPC peering route
	response, err := client.FromNetwork().VPCPeeringRoutes().Delete(r.Context(), projectId, vpcId, vpcPeeringId, id, params)
	if err != nil {
		h.Log.Printf("Failed to delete VPC peering route: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if response.Data != nil {
		if err := json.NewEncoder(w).Encode(response.Data); err != nil {
			h.Log.Printf("Failed to encode response: %v", err)
		}
	}
}

// ServeHTTP implementation for LIST handler
func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")
	vpcId := r.PathValue("vpcId")
	vpcPeeringId := r.PathValue("vpcPeeringId")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
		return
	}
	if vpcPeeringId == "" {
		http.Error(w, "vpcPeeringId is required", http.StatusBadRequest)
		return
	}

	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Listing VPC peering routes for vpc peering %s in vpc %s in project %s", vpcPeeringId, vpcId, projectId)

	// Call Aruba Cloud SDK to list VPC peering routes
	response, err := client.FromNetwork().VPCPeeringRoutes().List(r.Context(), projectId, vpcId, vpcPeeringId, params)
	if err != nil {
		h.Log.Printf("Failed to list VPC peering routes: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if err := json.NewEncoder(w).Encode(response.Data); err != nil {
		h.Log.Printf("Failed to encode response: %v", err)
	}
}

