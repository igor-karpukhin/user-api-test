package server

import (
	"context"
	"fmt"
	"net/http"

	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type HTTPServer struct {
	server *http.Server
	routes *mux.Router
	log    *zap.Logger
}

func NewHTTPServer(host string, port uint, router *mux.Router, log *zap.Logger) *HTTPServer {
	s := &HTTPServer{
		routes: router,
		log:    log,
	}
	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: s,
	}
	return s
}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.log.Debug("request", zap.String("from", r.Host))
	s.routes.ServeHTTP(w, r)
}

func (s *HTTPServer) Start(ctx context.Context) {
	go s.server.ListenAndServe()

	go func() {
		for {
			select {
			case <-ctx.Done():
				s.server.Shutdown(context.Background())
				return
			case <-time.After(1 * time.Second):
			}
		}
	}()

}
