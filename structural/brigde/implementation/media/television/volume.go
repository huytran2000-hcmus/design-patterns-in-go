package television

type volume struct {
	value int
}

const volumeInterval = 5

func (v volume) GetNext() volume {
	return volume{value: v.value + volumeInterval}
}

func (v volume) GetPrevious() volume {
	return volume{value: v.value - volumeInterval}
}

func (v volume) GetVolumeNumber() int {
	return v.value / volumeInterval
}

func getVolumeNumber(number int) volume {
	return volume{value: number * volumeInterval}
}
