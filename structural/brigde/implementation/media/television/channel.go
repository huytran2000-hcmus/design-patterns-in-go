package television

type channel struct {
	value int
}

const channelInterval = 1

func (c channel) GetNext() channel {
	return channel{value: c.value + channelInterval}
}

func (c channel) GetPrevious() channel {
	return channel{value: c.value - channelInterval}
}

func (c channel) GetChannelNumber() int {
	return c.value / channelInterval
}

func getChannelNumber(number int) channel {
	return channel{value: number * channelInterval}
}
