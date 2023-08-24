package main

import (
	"context"

	"github.com/edaniels/golog"
	"go.viam.com/rdk/components/generic"
	"go.viam.com/rdk/resource"
)

type Config struct {
	message string `json:"message"`
}

// checks the converted attributes but not sure what the path here is. I think it is the key?
func (c *Config) Validate(path string) ([]string, error) {
	print(path)
	return make([]string, 0), nil
}

var (
	Model = resource.NewModel("roxy", "simple", "printme")
)

func init() {
	registration := resource.Registration[resource.Resource, *Config]{
		Constructor: func(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger golog.Logger) (resource.Resource, error) {
			return build(ctx, deps, conf, logger)
		},
	}
	resource.RegisterComponent(generic.API, Model, registration)
}

type component struct {
	resource.Named
	resource.AlwaysRebuild
	cfg    *Config
	logger golog.Logger
	cancel func()
}

func (c *component) DoCommand() {
	c.logger.Info("i am doing the do command")
}

func build(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger golog.Logger) (resource.Resource, error) {
	logger.Info("building my simple component")
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	newConf, err := resource.NativeConfig[*Config](conf)
	if err != nil {

	}

	comp := &component{
		Named:  conf.ResourceName().AsNamed(),
		cfg:    newConf,
		logger: logger,
		cancel: cancelFunc,
	}

	// start a background process using this

	return comp, nil
}
