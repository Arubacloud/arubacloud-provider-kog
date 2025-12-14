package dbaas

import (
	"encoding/json"
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/sdk-go/pkg/types"
)

func GetDbaas(opts handlers.HandlerOptions) handlers.Handler {
	return &getHandler{baseHandler: newBaseHandler(opts)}
}

func PostDbaas(opts handlers.HandlerOptions) handlers.Handler {
	return &postHandler{baseHandler: newBaseHandler(opts)}
}

func PutDbaas(opts handlers.HandlerOptions) handlers.Handler {
	return &putHandler{baseHandler: newBaseHandler(opts)}
}

func ListDbaass(opts handlers.HandlerOptions) handlers.Handler {
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

	h.Log.Printf("Getting dbaas %s for project %s", id, projectId)

	// Call Aruba Cloud SDK to get dbaas
	response, err := client.FromDatabase().DBaaS().Get(r.Context(), projectId, id, params)
	if err != nil {
		h.Log.Printf("Failed to get dbaas: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response.Data)
}

// ServeHTTP implementation for POST handler
func (h *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")

	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Decode request body
	var reqBody interface{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert to typed request
	reqBytes, _ := json.Marshal(reqBody)
	var req types.DBaaSRequest
	json.Unmarshal(reqBytes, &req)

	// Build request parameters from query string
	params := &types.RequestParameters{}
	if apiVersion := r.URL.Query().Get("api-version"); apiVersion != "" {
		params.APIVersion = &apiVersion
	}

	h.Log.Printf("Creating dbaas for project %s", projectId)

	// Call Aruba Cloud SDK to create dbaas
	response, err := client.FromDatabase().DBaaS().Create(r.Context(), projectId, req, params)
	if err != nil {
		h.Log.Printf("Failed to create dbaas: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response.Data)
}

// ServeHTTP implementation for PUT handler
func (h *putHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")
	id := r.PathValue("id")

	// Create SDK client from request's Bearer token
	client, err := handlers.CreateClientFromRequest(r)
	if err != nil {
		h.Log.Printf("Failed to create Aruba Cloud client: %v", err)
		http.Error(w, "Failed to initialize API client", http.StatusInternalServerError)
		return
	}

	// Decode request body
	var reqBody interface{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert to typed request
	reqBytes, _ := json.Marshal(reqBody)
	var req types.DBaaSRequest
	json.Unmarshal(reqBytes, &req)

	// Build request parameters from query string
	params := &types.RequestParameters{}
	if apiVersion := r.URL.Query().Get("api-version"); apiVersion != "" {
		params.APIVersion = &apiVersion
	}

	h.Log.Printf("Updating dbaas %s for project %s", id, projectId)

	// Call Aruba Cloud SDK to update dbaas
	response, err := client.FromDatabase().DBaaS().Update(r.Context(), projectId, id, req, params)
	if err != nil {
		h.Log.Printf("Failed to update dbaas: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response.Data)
}

// ServeHTTP implementation for LIST handler
func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectId")

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

	h.Log.Printf("Listing dbaass for project %s", projectId)

	// Call Aruba Cloud SDK to list dbaass
	response, err := client.FromDatabase().DBaaS().List(r.Context(), projectId, params)
	if err != nil {
		h.Log.Printf("Failed to list dbaass: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response.Data)
}
