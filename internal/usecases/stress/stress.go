package stress

import (
	"log"

	"github.com/go-playground/validator/v10"

	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/pkg/httpclient"
	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/usecases/stress/dto"
)

type StressTestUseCaseInterface interface {
	Execute(input dto.StressTestInput) error
}

type StressTestUseCase struct {
	HTTPClient httpclient.HttpClientInterface
	validator  *validator.Validate
}

func NewStressTestUseCase(c httpclient.HttpClientInterface) *StressTestUseCase {
	return &StressTestUseCase{
		HTTPClient: c,
		validator:  validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (uc *StressTestUseCase) Execute(input dto.StressTestInput) error {
	log.Printf("Executing stress test with input: %+v", input)

	if err := uc.validateInput(input); err != nil {
		return err
	}

	return nil
}

func (uc *StressTestUseCase) validateInput(input dto.StressTestInput) error {
	return uc.validator.Struct(input)
}
