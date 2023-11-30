package gorilla

import (
	"encoding/json"
	"murasame29/go-fw/src/domain"
	"net/http"

	"github.com/gorilla/mux"
)

type gorillaServer struct {
	router *mux.Router
}

func NewGorillaServer() *mux.Router {
	server := &gorillaServer{
		router: mux.NewRouter(),
	}
	server.helloWorld()

	return server.router
}

func (s *gorillaServer) helloWorld() {
	s.router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(domain.Response{
			Message: "Hello World",
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})).Methods(http.MethodGet)
}
