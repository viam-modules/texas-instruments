//go:build linux

// Package ti implements a ti based board.
package ti

import (
	"context"

	"github.com/pkg/errors"
	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/board/genericlinux"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const modelName = "ti"

// Model for viam supported texas-instruments ti board.
var Model = resource.NewModel("viam", "texas-instruments", "ti")

func init() {
	initLogger := logging.NewLogger("ti.init")
	gpioMappings, err := genericlinux.GetGPIOBoardMappings(modelName, boardInfoMappings, initLogger)
	var noBoardErr genericlinux.NoBoardFoundError
	if errors.As(err, &noBoardErr) {
		initLogger.Debugw("error getting ti GPIO board mapping", "error", err)
	}

	resource.RegisterComponent(
		board.API,
		Model,
		resource.Registration[board.Board, *genericlinux.Config]{
			Constructor: func(
				ctx context.Context,
				_ resource.Dependencies,
				conf resource.Config,
				logger logging.Logger,
			) (board.Board, error) {
				return genericlinux.NewBoard(ctx, conf, genericlinux.ConstPinDefs(gpioMappings), logger)
			},
		})
}
