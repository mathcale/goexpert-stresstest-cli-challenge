package commands

import (
	"github.com/spf13/cobra"

	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/usecases/stress"
	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/usecases/stress/dto"
)

type StressTestCmd struct {
	StressTestUseCase stress.StressTestUseCaseInterface
}

func NewStressTestCmd(s stress.StressTestUseCaseInterface) *StressTestCmd {
	return &StressTestCmd{
		StressTestUseCase: s,
	}
}

func (s *StressTestCmd) Build() *cobra.Command {
	cmd := &cobra.Command{
		Short: "Stress test a given URL",
		Long:  "Executes a stress test on a given URL with a given number of requests and concurrency.",
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")
			requests, _ := cmd.Flags().GetUint64("requests")
			concurrency, _ := cmd.Flags().GetUint64("concurrency")

			input := dto.StressTestInput{
				URL:         url,
				Requests:    requests,
				Concurrency: concurrency,
			}

			return s.StressTestUseCase.Execute(input)
		},
	}

	cmd.Flags().String("url", "", "service URL to test")
	cmd.Flags().Uint64("requests", 0, "number of requests to perform")
	cmd.Flags().Uint64("concurrency", 0, "number of simultaneous requests to make at a time")

	cmd.MarkFlagRequired("url")
	cmd.MarkFlagRequired("requests")
	cmd.MarkFlagRequired("concurrency")

	return cmd
}
