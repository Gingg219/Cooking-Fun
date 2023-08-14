package main

import (
	"fmt"

	"net/http"

	"time"
)

func (a *application) ListenAndServer() error {
	url := fmt.Sprintf("%s:%s", a.server.host, a.server.port)

	srv := http.Server{
		Handler: a.routes(),
		Addr: url,
		ReadTimeout: 300 * time.Second,
	}

	a.infoLog.Printf("Server listening on :%s\n", url)

	return srv.ListenAndServe()
}