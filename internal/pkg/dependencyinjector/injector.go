package dependencyinjector

import (
	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/infra/cli"
	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/infra/cli/commands"
	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/pkg/httpclient"
	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/usecases/stress"
)

type DependencyInjectorInterface interface {
	Inject() (*Dependencies, error)
}

type DependencyInjector struct{}

type Dependencies struct {
	CLI cli.CLIInterface
}

func NewDependencyInjector() *DependencyInjector {
	return &DependencyInjector{}
}

func (d *DependencyInjector) Inject() (*Dependencies, error) {
	httpClient := httpclient.NewHttpClient()
	stressTestUseCase := stress.NewStressTestUseCase(httpClient)
	stressTestCmd := commands.NewStressTestCmd(stressTestUseCase)

	cli := cli.NewCLI(stressTestCmd.Build())

	return &Dependencies{
		CLI: cli,
	}, nil
}
