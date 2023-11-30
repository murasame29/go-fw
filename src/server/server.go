package server

type httpServer struct {
}

func NewServer() *httpServer {
	return &httpServer{}
}

func (s *httpServer) Run(router any) {
	runWithGracefulShutdown(router)
}
