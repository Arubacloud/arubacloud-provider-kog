package securityrule

import (
	"encoding/json"
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/sdk-go/pkg/types"
)

func GetSecurityrule(opts handlers.HandlerOptions) handlers.Handler {
	return &getHandler{baseHandler: newBaseHandler(opts)}
}

func PostSecurityrule(opts handlers.HandlerOptions) handlers.Handler {
	return &postHandler{baseHandler: newBaseHandler(opts)}
}

func PutSecurityrule(opts handlers.HandlerOptions) handlers.Handler {
	return &putHandler{baseHandler: newBaseHandler(opts)}
}

func ListSecurityrules(opts handlers.HandlerOptions) handlers.Handler {
	return &listHandler{baseHandler: newBaseHandler(opts)}
}

// Interface compliance verification
var _ handlers.Handler = &getHandler{}
var _ handlers.Handler = &postHandler{}
var _ handlers.Handler = &putHandler{}
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

type listHandler struct {
	*baseHandler
}

// ServeHTTP implementation for GET handler
func (h *getHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")
	vpcId := r.PathValue("vpcId")
	securityGroupId := r.PathValue("securityGroupId")
	id := r.PathValue("id")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
		return
	}
	if securityGroupId == "" {
		http.Error(w, "securityGroupId is required", http.StatusBadRequest)
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

	h.Log.Printf("Getting security rule %s for security group %s in vpc %s in project %s", id, securityGroupId, vpcId, projectId)

	// Call Aruba Cloud SDK to get security rule
	response, err := client.FromNetwork().SecurityGroupRules().Get(r.Context(), projectId, vpcId, securityGroupId, id, params)
	if err != nil {
		h.Log.Printf("Failed to get security rule: %v", err)
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
	securityGroupId := r.PathValue("securityGroupId")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
		return
	}
	if securityGroupId == "" {
		http.Error(w, "securityGroupId is required", http.StatusBadRequest)
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
	var req types.SecurityRuleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Creating security rule for security group %s in vpc %s in project %s", securityGroupId, vpcId, projectId)

	// Call Aruba Cloud SDK to create security rule
	response, err := client.FromNetwork().SecurityGroupRules().Create(r.Context(), projectId, vpcId, securityGroupId, req, params)
	if err != nil {
		h.Log.Printf("Failed to create security rule: %v", err)
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
	securityGroupId := r.PathValue("securityGroupId")
	id := r.PathValue("id")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
		return
	}
	if securityGroupId == "" {
		http.Error(w, "securityGroupId is required", http.StatusBadRequest)
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
	var req types.SecurityRuleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Updating security rule %s for security group %s in vpc %s in project %s", id, securityGroupId, vpcId, projectId)

	// Call Aruba Cloud SDK to update security rule
	response, err := client.FromNetwork().SecurityGroupRules().Update(r.Context(), projectId, vpcId, securityGroupId, id, req, params)
	if err != nil {
		h.Log.Printf("Failed to update security rule: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if err := json.NewEncoder(w).Encode(response.Data); err != nil {
		h.Log.Printf("Failed to encode response: %v", err)
	}
}

// ServeHTTP implementation for LIST handler
func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")
	vpcId := r.PathValue("vpcId")
	securityGroupId := r.PathValue("securityGroupId")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if vpcId == "" {
		http.Error(w, "vpcId is required", http.StatusBadRequest)
		return
	}
	if securityGroupId == "" {
		http.Error(w, "securityGroupId is required", http.StatusBadRequest)
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

	h.Log.Printf("Listing security rules for security group %s in vpc %s in project %s", securityGroupId, vpcId, projectId)

	// Call Aruba Cloud SDK to list security rules
	response, err := client.FromNetwork().SecurityGroupRules().List(r.Context(), projectId, vpcId, securityGroupId, params)
	if err != nil {
		h.Log.Printf("Failed to list security rules: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if err := json.NewEncoder(w).Encode(response.Data); err != nil {
		h.Log.Printf("Failed to encode response: %v", err)
	}
}
