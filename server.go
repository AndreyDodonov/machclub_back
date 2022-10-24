package apiserver

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

// запуск сервера
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	logrus.Info("server start at port ", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// остановка сервера
func (s *Server) Shutdown(ctx context.Context) error {
	logrus.Info("server shutdown")
	return s.httpServer.Shutdown(ctx)
}
