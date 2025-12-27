package backup

import (
	"encoding/json"
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/sdk-go/pkg/types"
)

func GetBackup(opts handlers.HandlerOptions) handlers.Handler {
	return &getHandler{baseHandler: newBaseHandler(opts)}
}

func PostBackup(opts handlers.HandlerOptions) handlers.Handler {
	return &postHandler{baseHandler: newBaseHandler(opts)}
}

func PutBackup(opts handlers.HandlerOptions) handlers.Handler {
	return &putHandler{baseHandler: newBaseHandler(opts)}
}

func ListBackups(opts handlers.HandlerOptions) handlers.Handler {
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
	id := r.PathValue("id")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
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

	h.Log.Printf("Getting storage backup %s for project %s", id, projectId)

	// Call Aruba Cloud SDK to get backup
	response, err := client.FromStorage().Backups().Get(r.Context(), projectId, id, params)
	if err != nil {
		h.Log.Printf("Failed to get backup: %v", err)
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

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
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
	var req types.StorageBackupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Build request parameters from query string
	params := handlers.BuildRequestParameters(r.URL.Query())

	h.Log.Printf("Creating storage backup for project %s", projectId)

	// Call Aruba Cloud SDK to create backup
	response, err := client.FromStorage().Backups().Create(r.Context(), projectId, req, params)
	if err != nil {
		h.Log.Printf("Failed to create backup: %v", err)
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
	// Storage backups resource does not support Update operation
	h.Log.Print("PUT storage backup operation not supported by Aruba Cloud API")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   "Method Not Allowed",
		"message": "PUT operation is not supported for storage backups",
	})
}

// ServeHTTP implementation for LIST handler
func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")

	if projectId == "" {
		http.Error(w, "projectId is required", http.StatusBadRequest)
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

	h.Log.Printf("Listing storage backups for project %s", projectId)

	// Call Aruba Cloud SDK to list backups
	response, err := client.FromStorage().Backups().List(r.Context(), projectId, params)
	if err != nil {
		h.Log.Printf("Failed to list backups: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if err := json.NewEncoder(w).Encode(response.Data); err != nil {
		h.Log.Printf("Failed to encode response: %v", err)
	}
}
