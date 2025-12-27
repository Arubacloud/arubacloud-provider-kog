package grant

import (
	"encoding/json"
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/sdk-go/pkg/types"
)

func GetGrant(opts handlers.HandlerOptions) handlers.Handler {
	return &getHandler{baseHandler: newBaseHandler(opts)}
}

func PostGrant(opts handlers.HandlerOptions) handlers.Handler {
	return &postHandler{baseHandler: newBaseHandler(opts)}
}

func PutGrant(opts handlers.HandlerOptions) handlers.Handler {
	return &putHandler{baseHandler: newBaseHandler(opts)}
}

func ListGrants(opts handlers.HandlerOptions) handlers.Handler {
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
	dbaasId := r.PathValue("dbaasId")
	databaseId := r.PathValue("databaseId")
	userId := r.PathValue("userId")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if dbaasId == "" {
		http.Error(w, "dbaasId is required", http.StatusBadRequest)
		return
	}
	if databaseId == "" {
		http.Error(w, "databaseId is required", http.StatusBadRequest)
		return
	}
	if userId == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
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

	h.Log.Printf("Getting grant for user %s", userId)

	// Call Aruba Cloud SDK to get grant
	response, err := client.FromDatabase().Grants().Get(r.Context(), projectId, dbaasId, databaseId, userId, params)
	if err != nil {
		h.Log.Printf("Failed to get grant: %v", err)
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
	dbaasId := r.PathValue("dbaasId")
	databaseId := r.PathValue("databaseId")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if dbaasId == "" {
		http.Error(w, "dbaasId is required", http.StatusBadRequest)
		return
	}
	if databaseId == "" {
		http.Error(w, "databaseId is required", http.StatusBadRequest)
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
	var req types.GrantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Creating grant for database %s", databaseId)

	// Call Aruba Cloud SDK to create grant
	response, err := client.FromDatabase().Grants().Create(r.Context(), projectId, dbaasId, databaseId, req, params)
	if err != nil {
		h.Log.Printf("Failed to create grant: %v", err)
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
	dbaasId := r.PathValue("dbaasId")
	databaseId := r.PathValue("databaseId")
	userId := r.PathValue("userId")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if dbaasId == "" {
		http.Error(w, "dbaasId is required", http.StatusBadRequest)
		return
	}
	if databaseId == "" {
		http.Error(w, "databaseId is required", http.StatusBadRequest)
		return
	}
	if userId == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
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
	var req types.GrantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Updating grant for user %s", userId)

	// Call Aruba Cloud SDK to update grant
	response, err := client.FromDatabase().Grants().Update(r.Context(), projectId, dbaasId, databaseId, userId, req, params)
	if err != nil {
		h.Log.Printf("Failed to update grant: %v", err)
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
	dbaasId := r.PathValue("dbaasId")
	databaseId := r.PathValue("databaseId")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	if dbaasId == "" {
		http.Error(w, "dbaasId is required", http.StatusBadRequest)
		return
	}
	if databaseId == "" {
		http.Error(w, "databaseId is required", http.StatusBadRequest)
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

	h.Log.Printf("Listing grants for database %s", databaseId)

	// Call Aruba Cloud SDK to list grants
	response, err := client.FromDatabase().Grants().List(r.Context(), projectId, dbaasId, databaseId, params)
	if err != nil {
		h.Log.Printf("Failed to list grants: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if err := json.NewEncoder(w).Encode(response.Data); err != nil {
		h.Log.Printf("Failed to encode response: %v", err)
	}
}
