package upstream

import (
	"fmt"
	"net/http"
)

const defaultStaticResponseCode = 200

// newStaticResponseHandler creates a new staticResponseHandler that serves a
// a static response code.
func newStaticResponseHandler(code *int) http.Handler {
	if code == nil {
		c := defaultStaticResponseCode
		code = &c
	}
	return &staticResponseHandler{code: *code}
}

// staticResponseHandler responds with a static response with the given response code.
type staticResponseHandler struct {
	code int
}

// ServeHTTP serves a static response.
func (s *staticResponseHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(s.code)
	fmt.Fprintf(rw, "Authenticated")
}
