// Program grpc_keyval_server is a memory-backed key/value gRPC server written
// in Go.
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"sync"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/api-definitions/keyval/go/internal/memserver"
	"github.com/api-definitions/keyval/go/keyvalpb"
)

var (
	grpcPort = flag.Int("port", 3133, "Port the key/value server should use for gRPC serving.")
)

var (
	maxMessageSizeOpt = grpc.MaxMsgSize(100 * 1000 * 1000)
)

func main() {
	// If beamx or Go flags are used, flags must be parsed first.
	flag.Parse()

	s := &memserver.Server{}

	stopKeyValCh := make(chan struct{})
	if err := listen(context.Background(), s, *grpcPort, stopKeyValCh, func() {}); err != nil {
		glog.Exitf("problem with server: %v", err)
	}
}

func listen(ctx context.Context, s *memserver.Server, port int, done <-chan struct{}, onReady func()) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer(maxMessageSizeOpt)

	if done == nil {
		done = make(chan struct{})
	}
	stopOnce := sync.Once{}
	stop := func() {
		stopOnce.Do(func() {
			grpcServer.GracefulStop()
		})
	}
	go func() {
		select {
		case <-ctx.Done():
			stop()
		case <-done:
			stop()
		}
	}()

	keyvalpb.RegisterKeyValueServiceServer(grpcServer, s)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	go onReady()
	glog.Infof("starting to server gRPC traffic on %v", lis.Addr())
	grpcServer.Serve(lis)
	return nil
}
