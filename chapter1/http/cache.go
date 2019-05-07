package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type cacheHandler struct {
	*Server
}

func (this *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	key := strings.Split(r.URL.EscapedPath(), "/")[2]

	log.Println("key: " + key)

	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := r.Method
	if m == http.MethodPut {
		b, _ := ioutil.ReadAll(r.Body)
		if len(b) != 0 {
			e := this.Set(key, b)
			if e != nil {
				log.Println(e)
				w.WriteHeader(http.StatusInsufficientStorage)
			}
		}

		return
	}

	if m == http.MethodGet {
		b, e := this.Get(key)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(b) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Write(b)
		return
	}

	if m == http.MethodDelete {
		e := this.Del(key)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (this *Server) cacheHandler() http.Handler {
	return &cacheHandler{this}
}
