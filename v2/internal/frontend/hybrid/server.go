//go:build exp && server
// +build exp,server

package hybrid

import (
	"context"

	"github.com/qimingzihaonanya1/wails/v2/internal/binding"
	"github.com/qimingzihaonanya1/wails/v2/internal/frontend"
	"github.com/qimingzihaonanya1/wails/v2/internal/frontend/desktop/null"
	"github.com/qimingzihaonanya1/wails/v2/internal/frontend/devserver"
	"github.com/qimingzihaonanya1/wails/v2/internal/logger"
	"github.com/qimingzihaonanya1/wails/v2/pkg/options"
)

// New returns a new Server frontend
// A server Frontend implementation contains a devserver.Frontend wrapping a null.Frontend
func NewFrontend(ctx context.Context, appoptions *options.App, logger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher) frontend.Frontend {
	return devserver.NewFrontend(ctx, appoptions, logger, appBindings, dispatcher, nil, null.NewFrontend(ctx, appoptions))
}
