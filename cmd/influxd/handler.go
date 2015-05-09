package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"net/url"
	"strings"

	"github.com/influxdb/influxdb"
	"github.com/influxdb/influxdb/httpd"
)

// Handler represents an HTTP handler for InfluxDB node.
// Depending on its role, it will serve many different endpoints.
type Handler struct {
	Server *influxdb.Server
	Config *Config
}

// NewHandler returns a new instance of Handler.
func NewHandler() *Handler {
	return &Handler{}
}

// ServeHTTP responds to HTTP request to the handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Debug routes.
	if h.Config.Debugging.PprofEnabled && strings.HasPrefix(r.URL.Path, "/debug/pprof") {
		switch r.URL.Path {
		case "/debug/pprof/cmdline":
			pprof.Cmdline(w, r)
		case "/debug/pprof/profile":
			pprof.Profile(w, r)
		case "/debug/pprof/symbol":
			pprof.Symbol(w, r)
		default:
			pprof.Index(w, r)
		}
		return
	}

	// Data node endpoints.  These are handled by data nodes and allow brokers and data
	// nodes to transfer state, process queries, etc..
	if strings.HasPrefix(r.URL.Path, "/data") {
		h.serveData(w, r)
		return
	}

	// These are public API endpoints
	h.serveAPI(w, r)
}

// serveData responds to broker requests
func (h *Handler) serveData(w http.ResponseWriter, r *http.Request) {
	if h.Server == nil {
		log.Println("no server configured to handle metadata endpoints")
		http.Error(w, "server unavailable", http.StatusServiceUnavailable)
		return
	}

	sh := httpd.NewClusterHandler(h.Server, h.Config.Authentication.Enabled,
		h.Config.Snapshot.Enabled, h.Config.Logging.HTTPAccess, version)
	sh.ServeHTTP(w, r)
	return
}

// serveAPI responds to data requests
func (h *Handler) serveAPI(w http.ResponseWriter, r *http.Request) {
	if h.Server == nil {
		log.Println("no server configured to handle data endpoints")
		http.Error(w, "server unavailable", http.StatusServiceUnavailable)
		return
	}

	sh := httpd.NewAPIHandler(h.Server, h.Config.Authentication.Enabled,
		h.Config.Logging.HTTPAccess, version)
	sh.WriteTrace = h.Config.Logging.WriteTracing
	sh.ServeHTTP(w, r)
	return
}

// redirect redirects a request to URL in u.  If u is an empty slice,
// a 503 is returned
func (h *Handler) redirect(u []url.URL, w http.ResponseWriter, r *http.Request) {
	// No valid URLs, return an error
	if len(u) == 0 {
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}

	// TODO: Log to internal stats to track how frequently this is happening. If
	// this is happening frequently, the clients are using a suboptimal endpoint

	// Redirect the client to a valid data node that can handle the request
	http.Redirect(w, r, u[rand.Intn(len(u))].String()+r.RequestURI, http.StatusTemporaryRedirect)
}
