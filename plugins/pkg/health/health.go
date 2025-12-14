package health

import (
	"net/http"
	"sync/atomic"
)

// LivenessHandler implements Kubernetes liveness probe
// Returns 200 if the application is running and hasn't deadlocked
func LivenessHandler(healthy *int32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(healthy) == 1 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service Unavailable"))
	}
}

// ReadinessHandler implements Kubernetes readiness probe
// Returns 200 if the application is ready to serve traffic
func ReadinessHandler(ready *int32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the service is marked as ready
		if atomic.LoadInt32(ready) == 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Service Not Ready"))
			return
		}

		// Service is ready
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ready"))
	}
}
