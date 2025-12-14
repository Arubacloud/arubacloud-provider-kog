package restore

import (
	"encoding/json"
	"net/http"

	"github.com/Arubacloud/arubacloud-provider-kog/plugins/pkg/handlers"
)

func GetRestore(opts handlers.HandlerOptions) handlers.Handler {
	return &getHandler{baseHandler: newBaseHandler(opts)}
}

func PostRestore(opts handlers.HandlerOptions) handlers.Handler {
	return &postHandler{baseHandler: newBaseHandler(opts)}
}

func PutRestore(opts handlers.HandlerOptions) handlers.Handler {
	return &putHandler{baseHandler: newBaseHandler(opts)}
}

func ListRestores(opts handlers.HandlerOptions) handlers.Handler {
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
	// Restores are nested resources under backups in the SDK
	// The current route structure doesn't include backupId, so we cannot support this operation
	h.Log.Print("GET restore operation requires backupId which is not in the route path")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   "Method Not Allowed",
		"message": "Restore operations require backupId in the path. SDK expects nested resource: /backups/{backupId}/restores/{id}",
	})
}

// ServeHTTP implementation for POST handler
func (h *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Restores are nested resources under backups in the SDK
	// The current route structure doesn't include backupId, so we cannot support this operation
	h.Log.Print("POST restore operation requires backupId which is not in the route path")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   "Method Not Allowed",
		"message": "Restore operations require backupId in the path. SDK expects nested resource: /backups/{backupId}/restores",
	})
}

// ServeHTTP implementation for PUT handler
func (h *putHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Restores are nested resources under backups in the SDK
	// The current route structure doesn't include backupId, so we cannot support this operation
	h.Log.Print("PUT restore operation requires backupId which is not in the route path")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   "Method Not Allowed",
		"message": "Restore operations require backupId in the path. SDK expects nested resource: /backups/{backupId}/restores/{id}",
	})
}

// ServeHTTP implementation for LIST handler
func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Restores are nested resources under backups in the SDK
	// The current route structure doesn't include backupId, so we cannot support this operation
	h.Log.Print("LIST restores operation requires backupId which is not in the route path")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   "Method Not Allowed",
		"message": "Restore operations require backupId in the path. SDK expects nested resource: /backups/{backupId}/restores",
	})
}
