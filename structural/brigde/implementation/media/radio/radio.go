package radio

type Radio struct {
	ch     channel
	vol    volume
	enable bool
}

const name = "Radio"

var (
	defaultChannel = channel{value: 1}
	defaultVolume  = volume{value: 10}
)

func New() *Radio {
	return &Radio{ch: defaultChannel, vol: defaultVolume}
}

func (t *Radio) GetName() string {
	return name
}

func (t *Radio) IsEnable() bool {
	return t.enable
}

func (t *Radio) Enable() {
	t.enable = true
}

func (t *Radio) Disable() {
	t.enable = false
}

func (t *Radio) GetChannel() int {
	return t.ch.GetChannelNumber()
}

func (t *Radio) SetChannel(val int) {
	t.ch = getChannelNumber(val)
}

func (t *Radio) ChannelUp() {
	t.ch = t.ch.GetNext()
}

func (t *Radio) ChannelDown() {
	t.ch = t.ch.GetPrevious()
}

func (t *Radio) GetVolume() int {
	return t.vol.GetVolumeNumber()
}

func (t *Radio) SetVolume(val int) {
	t.vol = getVolumeNumber(val)
}

func (t *Radio) VolumeUp() {
	t.vol = t.vol.GetNext()
}

func (t *Radio) VolumeDown() {
	t.vol = t.vol.GetPrevious()
}
