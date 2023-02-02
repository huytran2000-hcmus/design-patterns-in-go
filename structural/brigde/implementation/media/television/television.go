package television

type Television struct {
	ch     channel
	vol    volume
	enable bool
}

const name = "Television"

var (
	defaultChannel = channel{value: 1}
	defaultVolume  = volume{value: 10}
)

func New() *Television {
	return &Television{ch: defaultChannel, vol: defaultVolume}
}

func (t *Television) GetName() string {
	return name
}

func (t *Television) IsEnable() bool {
	return t.enable
}

func (t *Television) Enable() {
	t.enable = true
}

func (t *Television) Disable() {
	t.enable = false
}

func (t *Television) GetChannel() int {
	return t.ch.GetChannelNumber()
}

func (t *Television) SetChannel(val int) {
	t.ch = getChannelNumber(val)
}

func (t *Television) ChannelUp() {
	t.ch = t.ch.GetNext()
}

func (t *Television) ChannelDown() {
	t.ch = t.ch.GetPrevious()
}

func (t *Television) GetVolume() int {
	return t.vol.GetVolumeNumber()
}

func (t *Television) SetVolume(val int) {
	t.vol = getVolumeNumber(val)
}

func (t *Television) VolumeUp() {
	t.vol = t.vol.GetNext()
}

func (t *Television) VolumeDown() {
	t.vol = t.vol.GetPrevious()
}
