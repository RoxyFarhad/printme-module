package main

import (
	"context"

	"github.com/RoxyFarhad/printme-module/component"
	"github.com/edaniels/golog"
	"go.viam.com/rdk/components/generic"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("printme"))
}

func mainWithArgs(ctx context.Context, args []string, logger golog.Logger) error {
	module, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	if err := module.AddModelFromRegistry(ctx, generic.API, component.PrintMeModel); err != nil {
		logger.Errorw("error starting the module", "err", err)
		return err
	}

	err = module.Start(ctx)
	defer module.Close(ctx)
	if err != nil {
		logger.Errorw("error starting the module", "err", err)
		return err
	}
	logger.Infof("i made it here")
	<-ctx.Done()
	return nil
}
