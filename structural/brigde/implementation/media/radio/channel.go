package radio

type channel struct {
	value hz
}

type hz float64

const channelInterval hz = 1

func (c channel) GetNext() channel {
	return channel{value: c.value + channelInterval}
}

func (c channel) GetPrevious() channel {
	return channel{value: c.value - channelInterval}
}

func (c channel) GetChannelNumber() int {
	return int(c.value / channelInterval)
}

func getChannelNumber(number int) channel {
	return channel{value: hz(number) * channelInterval}
}
