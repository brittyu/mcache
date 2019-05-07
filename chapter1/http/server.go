package http

import (
	"net/http"

	"../cache"
)

type Server struct {
	cache.Cache
}

func (this *Server) Listen() {
	http.Handle("/cache/", this.cacheHandler())
	http.Handle("/status", this.statusHandler())
	http.ListenAndServe(":12345", nil)
}

func New(this cache.Cache) *Server {
	return &Server{this}
}
