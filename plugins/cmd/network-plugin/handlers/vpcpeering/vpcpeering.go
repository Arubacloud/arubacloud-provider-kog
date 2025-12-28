package vpcpeering

import (
	"encoding/json"
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/sdk-go/pkg/types"
)

func GetVPCPeering(opts handlers.HandlerOptions) handlers.Handler {
	return &getHandler{baseHandler: newBaseHandler(opts)}
}

func PostVPCPeering(opts handlers.HandlerOptions) handlers.Handler {
	return &postHandler{baseHandler: newBaseHandler(opts)}
}

func PutVPCPeering(opts handlers.HandlerOptions) handlers.Handler {
	return &putHandler{baseHandler: newBaseHandler(opts)}
}

func DeleteVPCPeering(opts handlers.HandlerOptions) handlers.Handler {
	return &deleteHandler{baseHandler: newBaseHandler(opts)}
}

func ListVPCPeerings(opts handlers.HandlerOptions) handlers.Handler {
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
	id := r.PathValue("id")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
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

	h.Log.Printf("Getting VPC peering %s for vpc %s in project %s", id, vpcId, projectId)

	// Call Aruba Cloud SDK to get VPC peering
	response, err := client.FromNetwork().VPCPeerings().Get(r.Context(), projectId, vpcId, id, params)
	if err != nil {
		h.Log.Printf("Failed to get VPC peering: %v", err)
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

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
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
	var req types.VPCPeeringRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Creating VPC peering for vpc %s in project %s", vpcId, projectId)

	// Call Aruba Cloud SDK to create VPC peering
	response, err := client.FromNetwork().VPCPeerings().Create(r.Context(), projectId, vpcId, req, params)
	if err != nil {
		h.Log.Printf("Failed to create VPC peering: %v", err)
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
	id := r.PathValue("id")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
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
	var req types.VPCPeeringRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Updating VPC peering %s for vpc %s in project %s", id, vpcId, projectId)

	// Call Aruba Cloud SDK to update VPC peering
	response, err := client.FromNetwork().VPCPeerings().Update(r.Context(), projectId, vpcId, id, req, params)
	if err != nil {
		h.Log.Printf("Failed to update VPC peering: %v", err)
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
	id := r.PathValue("id")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
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

	h.Log.Printf("Deleting VPC peering %s for vpc %s in project %s", id, vpcId, projectId)

	// Call Aruba Cloud SDK to delete VPC peering
	response, err := client.FromNetwork().VPCPeerings().Delete(r.Context(), projectId, vpcId, id, params)
	if err != nil {
		h.Log.Printf("Failed to delete VPC peering: %v", err)
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

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
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

	h.Log.Printf("Listing VPC peerings for vpc %s in project %s", vpcId, projectId)

	// Call Aruba Cloud SDK to list VPC peerings
	response, err := client.FromNetwork().VPCPeerings().List(r.Context(), projectId, vpcId, params)
	if err != nil {
		h.Log.Printf("Failed to list VPC peerings: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if err := json.NewEncoder(w).Encode(response.Data); err != nil {
		h.Log.Printf("Failed to encode response: %v", err)
	}
}

