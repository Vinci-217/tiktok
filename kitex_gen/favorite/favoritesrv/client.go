// Code generated by Kitex v0.3.2. DO NOT EDIT.

package favoritesrv

import (
	"context"
	"github.com/a76yyyy/tiktok/kitex_gen/favorite"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FavoriteAction(ctx context.Context, Req *favorite.DouyinFavoriteActionRequest, callOptions ...callopt.Option) (r *favorite.DouyinFavoriteActionResponse, err error)
	FavoriteList(ctx context.Context, Req *favorite.DouyinFavoriteListRequest, callOptions ...callopt.Option) (r *favorite.DouyinFavoriteListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFavoriteSrvClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFavoriteSrvClient struct {
	*kClient
}

func (p *kFavoriteSrvClient) FavoriteAction(ctx context.Context, Req *favorite.DouyinFavoriteActionRequest, callOptions ...callopt.Option) (r *favorite.DouyinFavoriteActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, Req)
}

func (p *kFavoriteSrvClient) FavoriteList(ctx context.Context, Req *favorite.DouyinFavoriteListRequest, callOptions ...callopt.Option) (r *favorite.DouyinFavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, Req)
}
