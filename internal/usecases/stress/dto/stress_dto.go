package dto

import (
	"time"

	"github.com/mathcale/goexpert-stresstest-cli-challenge/internal/pkg/httpclient"
)

type StressTestInput struct {
	URL         string `validate:"required,url"`
	Requests    uint64 `validate:"required,gt=0"`
	Concurrency uint64 `validate:"required,gt=0"`
}

type StressTestOutput struct {
	Duration time.Duration
	Results  []*httpclient.HttpClientResponse
}
