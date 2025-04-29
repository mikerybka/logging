package logging

import (
	"fmt"
	"net/http"

	"github.com/mikerybka/util"
)

type Server struct {
	Handler http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Printf("%s\t%s\t%s\n", util.ClientIP(r), r.Method, r.Host+r.URL.String())
	if err != nil {
		panic(err)
	}
	s.Handler.ServeHTTP(w, r)
}
