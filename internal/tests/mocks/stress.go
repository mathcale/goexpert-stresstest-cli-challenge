package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/usecases/stress/dto"
)

type StressTestUseCaseMock struct {
	mock.Mock
}

func (m *StressTestUseCaseMock) Execute(input dto.StressTestInput) (*dto.StressTestOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dto.StressTestOutput), args.Error(1)
}
