package panasonic

import (
	"errors"
	"fmt"
)

type ACModel interface {
	GetTemperature() int
	IncreaseTemperature() bool
	DecreaseTemperature() bool
	GetInfo() (min int, max int)
}
type ACRemote struct {
	model ACModel
}

func NewACRemote(model ACModel) *ACRemote {
	return &ACRemote{model: model}
}

func (r *ACRemote) GetTemperature() int {
	return r.model.GetTemperature()
}

func (r *ACRemote) IncreaseTemperature() error {
	result := r.model.IncreaseTemperature()
	if !result {
		temperature := r.GetTemperature() + 1
		return r.reportError(temperature)
	}

	return nil
}

func (r *ACRemote) DecreaseTemperature() error {
	result := r.model.DecreaseTemperature()
	if !result {
		temperature := r.GetTemperature() - 1
		return r.reportError(temperature)
	}

	return nil
}

func (r *ACRemote) reportError(temperature int) error {
	min, max := r.model.GetInfo()
	if temperature < min || temperature > max {
		return fmt.Errorf(
			"can't set air condition to %d°C, valid range [%d, %d]",
			temperature,
			min,
			max,
		)
	}

	return errors.New("air condition is has some problem, can't set degree")
}
