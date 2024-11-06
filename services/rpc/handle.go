package rpc

import (
	"context"

	"github.com/zhaohaibing666/wallet-sign-go/protobuf/wallet"
)

func (s *RpcServer) GetSupportSignWay(ctx context.Context, in *wallet.GetSupportSignWayRequest) (*wallet.GetSupportSignWayResponse, error) {

	return nil, nil
}

func (s *RpcServer) ExportPublicKeyList(ctx context.Context, in *wallet.ExportPublicKeyListRequest) (*wallet.ExportPublicKeyListResponse, error) {
	return nil, nil
}

func (s *RpcServer) SignTxMessage(ctx context.Context, in *wallet.SignTxMessageRequest) (*wallet.SignTxMessageResponse, error) {
	return nil, nil
}
