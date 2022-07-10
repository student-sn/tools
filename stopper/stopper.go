package stopper

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Stopper struct {
	client *mongo.Client
	ctx    context.Context
	server *http.Server
}

func Init(client *mongo.Client, ctx context.Context, server *http.Server) *Stopper {
	return &Stopper{
		client: client,
		ctx:    ctx,
		server: server,
	}
}

func (s *Stopper) Stop(code int, error ...any) {
	err := s.server.Shutdown(s.ctx)
	if err != nil {
		return
	}
	err = s.client.Disconnect(s.ctx)
	if err != nil {
		return
	}
	signal.NotifyContext(s.ctx, syscall.SIGINT)
	if code != 0 {
		panic(error[0])
	}
	os.Exit(code)
}
