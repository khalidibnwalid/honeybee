package handlers

import (
	"fmt"
	"net/http"
)

func NewServer(port int) (*http.Server, error) {
	ctx, err := NewServerHandlerContext()
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: ctx.registrar(),
	}, nil
}
