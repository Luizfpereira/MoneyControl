package webserver

import "net/http"

type WebServer struct {
	Router chi.
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		WebServerPort: webServerPort,
		Handlers:      make(map[string]http.HandlerFunc),
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() {

}
