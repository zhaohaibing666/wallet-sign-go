package rpc

import (
	"context"
	"fmt"
	"net"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zhaohaibing666/wallet-sign-go/protobuf/wallet"
)

const MaxRecvMessageSize = 1024 * 1024 * 30000

type RpcServerConfig struct {
	GrpcHostName string
	GrpcPort     int
}

type RpcServer struct {
	*RpcServerConfig
	wallet.UnimplementedWalletServiceServer
	stopped atomic.Bool
}

func NewRpcServer(conf *RpcServerConfig) (*RpcServer, error) {
	return &RpcServer{
		RpcServerConfig: conf,
	}, nil
}

func (s *RpcServer) Stop(ctx context.Context) error {
	s.stopped.Store(true)
	return nil
}

func (s *RpcServer) Stopped() bool {
	return s.stopped.Load()
}

func (s *RpcServer) Start(ctx context.Context) error {
	go func(s *RpcServer) {
		addr := fmt.Sprintf("%s:%d", s.RpcServerConfig.GrpcHostName, s.RpcServerConfig.GrpcPort)
		log.Info("start rpc services", "addr", addr)

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Error("could not start tcp listener")
		}

		gs := grpc.NewServer(grpc.MaxRecvMsgSize(MaxRecvMessageSize))

		reflection.Register(gs)

		wallet.RegisterWalletServiceServer(gs, s)

		if err := gs.Serve(listener); err != nil {
			log.Error("Could not GRPC services")
		}

	}(s)
	return nil
}
