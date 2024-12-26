// package main is a module with raspberry pi board component.
package main

import (
	"context"

	"github.com/viam-modules/texas-instruments/ina"
	"github.com/viam-modules/texas-instruments/ti"

	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/powersensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("texas-instruments"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	module, err := module.NewModuleFromArgs(ctx)
	if err != nil {
		return err
	}

	if err = module.AddModelFromRegistry(ctx, board.API, ti.Model); err != nil {
		return err
	}

	if err = module.AddModelFromRegistry(ctx, powersensor.API, ina.Model219); err != nil {
		return err
	}

	if err = module.AddModelFromRegistry(ctx, powersensor.API, ina.Model226); err != nil {
		return err
	}

	err = module.Start(ctx)
	defer module.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
