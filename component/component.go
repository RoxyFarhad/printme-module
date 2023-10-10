package component

import (
	"context"
	"time"

	"github.com/edaniels/golog"
	"go.viam.com/rdk/components/generic"
	"go.viam.com/rdk/resource"
	"go.viam.com/utils"
)

type Config struct {
	Message string `json:"message"`
}

// checks the converted attributes but not sure what the path here is. I think it is the key?
func (c *Config) Validate(path string) ([]string, error) {
	print(path)
	return make([]string, 0), nil
}

var (
	PrintMeModel = resource.NewModel("roxy", "simple", "printme")
)

func init() {
	registration := resource.Registration[resource.Resource, *Config]{
		Constructor: func(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger golog.Logger) (resource.Resource, error) {
			return build(deps, conf, logger)
		},
	}
	resource.RegisterComponent(generic.API, PrintMeModel, registration)
}

type component struct {
	resource.Named
	resource.AlwaysRebuild
	cfg    *Config
	logger golog.Logger
	cancel func()
}

func (c *component) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	c.logger.Info("i am doing the do command")
	return make(map[string]interface{}), nil
}

func (c *component) Close(ctx context.Context) error {
	c.logger.Infof("i am trying to close")
	c.cancel()
	return nil
}

// build creates the component and starts running it in the background
func build(deps resource.Dependencies, conf resource.Config, logger golog.Logger) (resource.Resource, error) {
	logger.Info("building my simple component")
	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	newConf, err := resource.NativeConfig[*Config](conf)
	if err != nil {
		cancelFunc()
		return nil, err
	}

	comp := &component{
		Named:  conf.ResourceName().AsNamed(),
		cfg:    newConf,
		logger: logger,
		cancel: cancelFunc,
	}

	// runs the background go routine
	comp.run(cancelCtx)
	return comp, nil
}

// run runs the main functionality of the component
func (c *component) run(ctx context.Context) {
	utils.PanicCapturingGo(func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case tick := <-ticker.C:
				c.logger.Infof("i am running bish for %d second", tick.Second())
			case <-ctx.Done():
				c.logger.Infof("exiting the loop as the context is done")
				return
			}
		}
	})

}
