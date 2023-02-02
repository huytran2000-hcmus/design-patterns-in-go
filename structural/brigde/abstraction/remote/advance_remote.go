package remote

type AdvanceRemote struct {
	*BasicRemote
}

func NewAdvanceRemote(device Device) *AdvanceRemote {
	return &AdvanceRemote{BasicRemote: NewBasicRemote(device)}
}

func (rm *AdvanceRemote) SetChannel(val int) {
	rm.device.SetChannel(val)
}

func (rm *AdvanceRemote) SetVolume(val int) {
	rm.device.SetVolume(val)
}
