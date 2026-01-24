package myntp

import (
	"errors"
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

// NewTime is new struct for task
type NewTime struct {
	Time      time.Time
	Precision time.Duration
}

// CurrentTime is func for task 2
func CurrentTime(network string) (*NewTime, error) {
	resp := NewTime{}
	if network == "" {
		network = "pool.ntp.org"
	}
	response, err := ntp.Query(network)
	if err != nil {
		return &resp, errors.New("Ошибка запроса")
	}
	err = response.Validate()
	if err != nil {
		return &resp, fmt.Errorf("Ошибка валидации")
	}
	resp = NewTime{
		time.Now().Add(response.ClockOffset),
		response.Precision,
	}
	return &resp, nil
}
