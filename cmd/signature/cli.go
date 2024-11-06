package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/params"

	"github.com/zhaohaibing666/wallet-sign-go/config"
	flags2 "github.com/zhaohaibing666/wallet-sign-go/flags"

	"github.com/zhaohaibing666/wallet-sign-go/common/cliapp"
	"github.com/zhaohaibing666/wallet-sign-go/services/rpc"
)

func runRpc(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	fmt.Println("running grpc services...")
	cfg := config.NewConfig(ctx)
	grpcServerCfg := &rpc.RpcServerConfig{
		GrpcHostName: cfg.RpcServer.Host,
		GrpcPort:     cfg.RpcServer.Port,
	}

	return rpc.NewRpcServer(grpcServerCfg)
}

func NewCli(GitCommit string, GitData string) *cli.App {
	flags := flags2.Flags
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitData),
		Description:          "An exchange wallet scanner services with rpc and rest api services",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "rpc",
				Flags:       flags,
				Description: "Run rpc services",
				Action:      cliapp.LifecycleCmd(runRpc),
			},
			{
				Name:        "version",
				Description: "Show project version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
