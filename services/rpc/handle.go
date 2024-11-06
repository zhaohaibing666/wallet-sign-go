package rpc

import (
	"context"
	"strconv"

	"github.com/zhaohaibing666/wallet-sign-go/protobuf/wallet"
	"github.com/zhaohaibing666/wallet-sign-go/ssm"
)

func (s *RpcServer) GetSupportSignWay(ctx context.Context, in *wallet.GetSupportSignWayRequest) (*wallet.GetSupportSignWayResponse, error) {
	if in.Type == "ecdsa" || in.Type == "eddsa" {
		return &wallet.GetSupportSignWayResponse{
			Code:    strconv.Itoa(1),
			Msg:     "support the sign way",
			Support: true,
		}, nil
	} else {
		return &wallet.GetSupportSignWayResponse{
			Code:    strconv.Itoa(-1),
			Msg:     "unsupport the sign way",
			Support: false,
		}, nil
	}
}

func (s *RpcServer) ExportPublicKeyList(ctx context.Context, in *wallet.ExportPublicKeyListRequest) (*wallet.ExportPublicKeyListResponse, error) {
	var retKeyList []*wallet.PublicKey
	for i := 0; i < int(in.Number); i++ {
		_, pubKeyStr, decPubKey, err := ssm.CreateESDAKeyPair()
		if err != nil {
			return nil, err
		}
		pubItem := &wallet.PublicKey{
			CompressPubkey:   pubKeyStr,
			DecompressPubkey: decPubKey,
		}
		retKeyList = append(retKeyList, pubItem)
	}

	return &wallet.ExportPublicKeyListResponse{
		Code:      strconv.Itoa(1),
		Msg:       "support the sign way",
		PublicKey: retKeyList,
	}, nil
}

func (s *RpcServer) SignTxMessage(ctx context.Context, in *wallet.SignTxMessageRequest) (*wallet.SignTxMessageResponse, error) {

	privateKey := "111"

	signature, err := ssm.SignMessage(privateKey, in.MessageHash)
	if err != nil {
		return nil, err
	}
	return &wallet.SignTxMessageResponse{
		Code:      strconv.Itoa(1),
		Msg:       "sign tx success",
		Signature: signature,
	}, nil

}
