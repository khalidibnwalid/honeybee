package handlers

import (
	"errors"
	"khalidibnwalid/luma_server/internal/server"
)

// wraps the server context to be used in handlers.
type ServerHandlerContext struct {
	*server.ServerContext
}

func NewServerHandlerContext(ctx ...*server.ServerContext) (*ServerHandlerContext , error){
	var _ctx *server.ServerContext
	if len(ctx) == 0 {
		ctx, err := server.NewServerContext()
		if err != nil {
			return nil, errors.New("failed to create server context: " + err.Error())
		}
		_ctx = ctx
	} else {
		_ctx = ctx[0]
	}

	return &ServerHandlerContext{
		ServerContext: _ctx,
	}, nil
}