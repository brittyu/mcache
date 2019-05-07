package tcp

import (
	"net"

	"../cache"
)

type Server struct {
	cache.Cache
}

func (this *Server) Listen() {
	l, err := net.Listen("tcp", ":12346")
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go this.process(c)
	}
}

func New(this cache.Cache) *Server {
	return &Server{this}
}
