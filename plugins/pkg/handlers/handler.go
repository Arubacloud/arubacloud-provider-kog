package handlers

import (
	"net/http"
	"strings"
	"sync"

	"github.com/Arubacloud/sdk-go/pkg/aruba"
)

// Logger interface allows mocking of logger
type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type HandlerOptions struct {
	Log Logger // Logger interface
}

// Handler interface
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// clientCache stores SDK clients by token to avoid recreating them
var (
	clientCache = sync.Map{}
)

// CreateClientFromRequest creates or retrieves a cached Aruba SDK client using the Bearer token from the request
func CreateClientFromRequest(r *http.Request) (aruba.Client, error) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Check if client already exists in cache
	if cached, ok := clientCache.Load(token); ok {
		return cached.(aruba.Client), nil
	}

	// Create new client
	options := aruba.NewOptions()
	options.WithNativeLogger()
	options.WithToken(token)

	client, err := aruba.NewClient(options)
	if err != nil {
		return nil, err
	}

	// Store in cache
	clientCache.Store(token, client)

	return client, nil
}
