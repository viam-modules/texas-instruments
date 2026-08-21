//go:build linux

// Package ti implements a ti based board.
package ti

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/board/genericlinux"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const modelName = "ti"

// Model for viam supported texas-instruments ti board.
var Model = resource.NewModel("viam", "texas-instruments", "ti")

var (
	gpioMappingsOnce   sync.Once
	cachedGPIOMappings map[string]genericlinux.GPIOBoardMapping
)

// getGPIOMappings looks up the board's GPIO mappings once per process; every
// call after the first returns the cached result, so only the first caller's
// logger receives the lookup's log output, including the debug message when
// no board is found.
func getGPIOMappings(logger logging.Logger) map[string]genericlinux.GPIOBoardMapping {
	gpioMappingsOnce.Do(func() {
		var err error
		cachedGPIOMappings, err = genericlinux.GetGPIOBoardMappings(modelName, boardInfoMappings, logger)
		var noBoardErr genericlinux.NoBoardFoundError
		if errors.As(err, &noBoardErr) {
			logger.Debugw("error getting ti GPIO board mapping", "error", err)
		}
	})
	return cachedGPIOMappings
}

func init() {
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
				return genericlinux.NewBoard(ctx, conf, genericlinux.ConstPinDefs(getGPIOMappings(logger)), logger)
			},
		})
}
