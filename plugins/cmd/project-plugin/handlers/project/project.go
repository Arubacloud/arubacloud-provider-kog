package project

import (
	"encoding/json"
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/sdk-go/pkg/types"
)

func GetProject(opts handlers.HandlerOptions) handlers.Handler {
	return &getHandler{baseHandler: newBaseHandler(opts)}
}

func PostProject(opts handlers.HandlerOptions) handlers.Handler {
	return &postHandler{baseHandler: newBaseHandler(opts)}
}

func PutProject(opts handlers.HandlerOptions) handlers.Handler {
	return &putHandler{baseHandler: newBaseHandler(opts)}
}

func ListProjects(opts handlers.HandlerOptions) handlers.Handler {
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
	id := r.PathValue("id")

	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Build request parameters from query string
	params := &types.RequestParameters{}
	if apiVersion := r.URL.Query().Get("api-version"); apiVersion != "" {
		params.APIVersion = &apiVersion
	}

	// Call Aruba Cloud SDK to get project
	h.Log.Printf("Getting project %s", id)
	response, err := client.FromProject().Get(r.Context(), id, params)
	if err != nil {
		h.Log.Printf("Failed to get project: %v", err)
		http.Error(w, "Failed to get project", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response.Data)
}

// ServeHTTP implementation for POST handler
func (h *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Decode request body
	var req types.ProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Build request parameters
	params := &types.RequestParameters{}
	if apiVersion := r.URL.Query().Get("api-version"); apiVersion != "" {
		params.APIVersion = &apiVersion
	}

	// Call Aruba Cloud SDK to create project
	h.Log.Printf("Creating project")
	response, err := client.FromProject().Create(r.Context(), req, params)
	if err != nil {
		h.Log.Printf("Failed to create project: %v", err)
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response.Data)
}

// ServeHTTP implementation for PUT handler
func (h *putHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Decode request body
	var req types.ProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Build request parameters
	params := &types.RequestParameters{}
	if apiVersion := r.URL.Query().Get("api-version"); apiVersion != "" {
		params.APIVersion = &apiVersion
	}

	// Call Aruba Cloud SDK to update project
	h.Log.Printf("Updating project %s", id)
	response, err := client.FromProject().Update(r.Context(), id, req, params)
	if err != nil {
		h.Log.Printf("Failed to update project: %v", err)
		http.Error(w, "Failed to update project", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response.Data)
}

// ServeHTTP implementation for LIST handler
func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Build request parameters from query string
	params := &types.RequestParameters{}
	query := r.URL.Query()
	if apiVersion := query.Get("api-version"); apiVersion != "" {
		params.APIVersion = &apiVersion
	}
	if filter := query.Get("filter"); filter != "" {
		params.Filter = &filter
	}
	if sort := query.Get("sort"); sort != "" {
		params.Sort = &sort
	}

	// Call Aruba Cloud SDK to list projects
	h.Log.Printf("Listing projects")
	response, err := client.FromProject().List(r.Context(), params)
	if err != nil {
		h.Log.Printf("Failed to list projects: %v", err)
		http.Error(w, "Failed to list projects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response.Data)
}
