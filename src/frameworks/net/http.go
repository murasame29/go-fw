package net

import (
	"encoding/json"
	"murasame29/go-fw/src/domain"
	"net/http"
)

type netServer struct {
	router *http.ServeMux
}

func NewNetHttpServer() *http.ServeMux {
	server := &netServer{
		router: http.NewServeMux(),
	}
	server.helloWorld()

	return server.router
}

func (s *netServer) helloWorld() {
	s.router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(domain.Response{
			Message: "Hello World",
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}))
}
