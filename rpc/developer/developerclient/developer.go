// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: developer.proto

package developerclient

import (
	"context"

	"github.com/wushiling50/aster/gen/developer"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddDeveloperReq         = developer.AddDeveloperReq
	AddDeveloperResp        = developer.AddDeveloperResp
	DelDeveloperByIdReq     = developer.DelDeveloperByIdReq
	DelDeveloperByIdResp    = developer.DelDeveloperByIdResp
	DelDeveloperByLoginReq  = developer.DelDeveloperByLoginReq
	DelDeveloperByLoginResp = developer.DelDeveloperByLoginResp
	Developer               = developer.Developer
	GetDeveloperByIdReq     = developer.GetDeveloperByIdReq
	GetDeveloperByIdResp    = developer.GetDeveloperByIdResp
	GetDeveloperByLoginReq  = developer.GetDeveloperByLoginReq
	GetDeveloperByLoginResp = developer.GetDeveloperByLoginResp
	UpdateDeveloperReq      = developer.UpdateDeveloperReq
	UpdateDeveloperResp     = developer.UpdateDeveloperResp

	DeveloperZrpcClient interface {
		AddDeveloper(ctx context.Context, in *AddDeveloperReq, opts ...grpc.CallOption) (*AddDeveloperResp, error)
		UpdateDeveloper(ctx context.Context, in *UpdateDeveloperReq, opts ...grpc.CallOption) (*UpdateDeveloperResp, error)
		DelDeveloperById(ctx context.Context, in *DelDeveloperByIdReq, opts ...grpc.CallOption) (*DelDeveloperByIdResp, error)
		DelDeveloperByLogin(ctx context.Context, in *DelDeveloperByLoginReq, opts ...grpc.CallOption) (*DelDeveloperByLoginResp, error)
		GetDeveloperById(ctx context.Context, in *GetDeveloperByIdReq, opts ...grpc.CallOption) (*GetDeveloperByIdResp, error)
		GetDeveloperByLogin(ctx context.Context, in *GetDeveloperByLoginReq, opts ...grpc.CallOption) (*GetDeveloperByLoginResp, error)
	}

	defaultDeveloperZrpcClient struct {
		cli zrpc.Client
	}
)

func NewDeveloperZrpcClient(cli zrpc.Client) DeveloperZrpcClient {
	return &defaultDeveloperZrpcClient{
		cli: cli,
	}
}

func (m *defaultDeveloperZrpcClient) AddDeveloper(ctx context.Context, in *AddDeveloperReq, opts ...grpc.CallOption) (*AddDeveloperResp, error) {
	client := developer.NewDeveloperClient(m.cli.Conn())
	return client.AddDeveloper(ctx, in, opts...)
}

func (m *defaultDeveloperZrpcClient) UpdateDeveloper(ctx context.Context, in *UpdateDeveloperReq, opts ...grpc.CallOption) (*UpdateDeveloperResp, error) {
	client := developer.NewDeveloperClient(m.cli.Conn())
	return client.UpdateDeveloper(ctx, in, opts...)
}

func (m *defaultDeveloperZrpcClient) DelDeveloperById(ctx context.Context, in *DelDeveloperByIdReq, opts ...grpc.CallOption) (*DelDeveloperByIdResp, error) {
	client := developer.NewDeveloperClient(m.cli.Conn())
	return client.DelDeveloperById(ctx, in, opts...)
}

func (m *defaultDeveloperZrpcClient) DelDeveloperByLogin(ctx context.Context, in *DelDeveloperByLoginReq, opts ...grpc.CallOption) (*DelDeveloperByLoginResp, error) {
	client := developer.NewDeveloperClient(m.cli.Conn())
	return client.DelDeveloperByLogin(ctx, in, opts...)
}

func (m *defaultDeveloperZrpcClient) GetDeveloperById(ctx context.Context, in *GetDeveloperByIdReq, opts ...grpc.CallOption) (*GetDeveloperByIdResp, error) {
	client := developer.NewDeveloperClient(m.cli.Conn())
	return client.GetDeveloperById(ctx, in, opts...)
}

func (m *defaultDeveloperZrpcClient) GetDeveloperByLogin(ctx context.Context, in *GetDeveloperByLoginReq, opts ...grpc.CallOption) (*GetDeveloperByLoginResp, error) {
	client := developer.NewDeveloperClient(m.cli.Conn())
	return client.GetDeveloperByLogin(ctx, in, opts...)
}
