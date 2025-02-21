package server

import (
	"Library/config"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	Serve(ctx context.Context) error
}
type HttpServer struct {
	conf config.Server
	log  *zap.Logger
	srv  *http.Server
}

func NewHttpServer(conf config.Server, logger *zap.Logger, server *http.Server) Server {
	return &HttpServer{conf: conf, log: logger, srv: server}
}
func (s *HttpServer) Serve(ctx context.Context) error {
	serverStarted := make(chan struct{})

	go func() {
		s.log.Info("Starting server on :8080")
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Error("Error starting server:", zap.Error(err))
		}

	}()

	go func() {
		for {
			conn, err := net.DialTimeout("tcp", s.srv.Addr, 1*time.Second)
			if err == nil {
				conn.Close()
				close(serverStarted)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {
	case <-serverStarted:
		s.log.Info(fmt.Sprintf("Server started successfully on %s", s.srv.Addr))
	case <-time.After(5 * time.Second):
		s.log.Error("Server failed to start within 5 seconds")
		return fmt.Errorf("server failed to start within 5 seconds")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Error("Error shutting down server: %v", zap.Error(err))
	}

	s.log.Info("Server stopped")

	return nil
}
