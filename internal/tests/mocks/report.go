package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/usecases/report/dto"
)

type ReportUseCaseMock struct {
	mock.Mock
}

func (m *ReportUseCaseMock) Execute(input dto.ReportInput) *dto.ReportOutput {
	args := m.Called(input)
	return args.Get(0).(*dto.ReportOutput)
}
