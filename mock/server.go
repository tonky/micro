package mock

import (
	"net/http"
	"net/http/httptest"
)

// see motivation behind this: http://big-elephants.com/2017-09/this-programmer-tried-to-mock-an-http-slash-2-server-in-go-and-heres-what-happened/
func ServerMock() (baseURL string, mux *http.ServeMux, teardownFn func()) {
	mux = http.NewServeMux()
	srv := httptest.NewServer(mux)
	return srv.URL, mux, srv.Close
}
