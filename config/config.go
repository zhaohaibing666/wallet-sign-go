package config

import (
	"github.com/urfave/cli/v2"
	"github.com/zhaohaibing666/wallet-sign-go/flags"
)

type ServerConfig struct {
	Host string
	Port int
}

type Config struct {
	RpcServer ServerConfig
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		RpcServer: ServerConfig{
			Host: ctx.String(flags.RpcHostFlag.Name),
			Port: ctx.Int(flags.RpcPortFlag.Name),
		},
	}
}
