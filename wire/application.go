package wire

import "net/http"

type Application struct {
	HttpServer *http.Server
}

func newApplication(httpServer *http.Server) Application {
	return Application{
		HttpServer: httpServer,
	}
}
