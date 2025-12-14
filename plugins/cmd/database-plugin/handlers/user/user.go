package user

import (
	"encoding/json"
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
	"github.com/Arubacloud/sdk-go/pkg/types"
)

func GetUser(opts handlers.HandlerOptions) handlers.Handler {
	return &getHandler{baseHandler: newBaseHandler(opts)}
}

func PostUser(opts handlers.HandlerOptions) handlers.Handler {
	return &postHandler{baseHandler: newBaseHandler(opts)}
}

func PutUser(opts handlers.HandlerOptions) handlers.Handler {
	return &putHandler{baseHandler: newBaseHandler(opts)}
}

func ListUsers(opts handlers.HandlerOptions) handlers.Handler {
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

	h.Log.Printf("Getting user %s for dbaas %s in project %s", id, dbaasId, projectId)

	// Call Aruba Cloud SDK to get user
	response, err := client.FromDatabase().Users().Get(r.Context(), projectId, dbaasId, id, params)
	if err != nil {
		h.Log.Printf("Failed to get user: %v", err)
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
	dbaasId := r.PathValue("dbaasId")

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
	var req types.UserRequest
	json.Unmarshal(reqBytes, &req)

	// Build request parameters from query string
	params := &types.RequestParameters{}
	if apiVersion := r.URL.Query().Get("api-version"); apiVersion != "" {
		params.APIVersion = &apiVersion
	}

	h.Log.Printf("Creating user for dbaas %s in project %s", dbaasId, projectId)

	// Call Aruba Cloud SDK to create user
	response, err := client.FromDatabase().Users().Create(r.Context(), projectId, dbaasId, req, params)
	if err != nil {
		h.Log.Printf("Failed to create user: %v", err)
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
	dbaasId := r.PathValue("dbaasId")
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
	var req types.UserRequest
	json.Unmarshal(reqBytes, &req)

	// Build request parameters from query string
	params := &types.RequestParameters{}
	if apiVersion := r.URL.Query().Get("api-version"); apiVersion != "" {
		params.APIVersion = &apiVersion
	}

	h.Log.Printf("Updating user %s for dbaas %s in project %s", id, dbaasId, projectId)

	// Call Aruba Cloud SDK to update user
	response, err := client.FromDatabase().Users().Update(r.Context(), projectId, dbaasId, id, req, params)
	if err != nil {
		h.Log.Printf("Failed to update user: %v", err)
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
	dbaasId := r.PathValue("dbaasId")

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

	h.Log.Printf("Listing users for dbaas %s in project %s", dbaasId, projectId)

	// Call Aruba Cloud SDK to list users
	response, err := client.FromDatabase().Users().List(r.Context(), projectId, dbaasId, params)
	if err != nil {
		h.Log.Printf("Failed to list users: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response.Data)
}
