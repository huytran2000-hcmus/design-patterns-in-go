package adapter

type PanasonicACModel interface {
	GetTemperature() int
	IncreaseTemperature() bool
	DecreaseTemperature() bool
	GetInfo() (min int, max int)
}
type PanasonicAC struct {
	model PanasonicACModel
}

func NewPanasonicAC(model PanasonicACModel) *PanasonicAC {
	return &PanasonicAC{model: model}
}

func (p *PanasonicAC) SetTemperature(temperature int) bool {
	currentTemperature := p.model.GetTemperature()
	difference := temperature - currentTemperature

	if difference > 0 {
		return p.increaseTemperatureBy(difference)
	} else {
		difference := -difference
		return p.decreaseTemperatureBy(difference)
	}
}

func (p *PanasonicAC) decreaseTemperatureBy(difference int) bool {
	for i := 0; i < difference; i++ {
		result := p.model.DecreaseTemperature()
		if !result {
			return false
		}
	}
	return true
}

func (p *PanasonicAC) increaseTemperatureBy(n int) bool {
	for i := 0; i < n; i++ {
		result := p.model.IncreaseTemperature()
		if !result {
			return false
		}
	}
	return true
}

func (p *PanasonicAC) GetTemperature() (min int, tempt int, max int) {
	min, max = p.model.GetInfo()
	tempt = p.model.GetTemperature()

	return
}
