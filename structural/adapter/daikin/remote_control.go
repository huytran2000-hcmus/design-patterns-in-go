package daikin

import (
	"errors"
	"fmt"
)

type AirCondition interface {
	SetTemperature(temperature int) bool
	GetTemperature() (min int, tempt int, max int)
}

type ACRemote struct {
	ac AirCondition
}

func NewACRemote(ac AirCondition) *ACRemote {
	return &ACRemote{ac: ac}
}

func (r *ACRemote) GetTemperature() (min int, tempt int, max int) {
	return r.ac.GetTemperature()
}

func (r *ACRemote) SetTemperature(temperature int) error {
	result := r.ac.SetTemperature(temperature)

	if !result {
		return r.reportError(temperature)
	}

	return nil
}

func (r *ACRemote) reportError(temperature int) error {
	min, _, max := r.ac.GetTemperature()
	if temperature < min || temperature > max {
		return fmt.Errorf(
			"can't set AC to %d°C ([%d°C, %d°C])",
			temperature,
			min,
			max,
		)
	}

	return errors.New("air condition is has some problem, can't set degree")
}
