package panasonic

type airTemperature int

type AirCondition struct {
	temperature airTemperature
}

const (
	minTemperature     airTemperature = 16
	maxTemperature     airTemperature = 31
	defaultTemperature airTemperature = 20
)

func (t airTemperature) isValid() bool {
	if minTemperature <= t && t <= maxTemperature {
		return true
	}

	return false
}

func NewAirCondition() *AirCondition {
	return &AirCondition{temperature: defaultTemperature}
}

func (ac *AirCondition) GetTemperature() int {
	return int(ac.temperature)
}

func (ac *AirCondition) IncreaseTemperature() bool {
	increasedTemperature := ac.temperature + 1

	if !increasedTemperature.isValid() {
		return false
	}

	ac.temperature = increasedTemperature
	return true
}

func (ac *AirCondition) DecreaseTemperature() bool {
	decreasedTemperature := ac.temperature - 1

	if !decreasedTemperature.isValid() {
		return false
	}

	ac.temperature = decreasedTemperature
	return true
}

func (ac *AirCondition) GetInfo() (min int, max int) {
	return int(minTemperature), int(maxTemperature)
}
