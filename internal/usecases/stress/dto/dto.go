package dto

type StressTestInput struct {
	URL         string `validate:"required,url"`
	Requests    uint64 `validate:"required,gt=0"`
	Concurrency uint64 `validate:"required,gt=0"`
}
