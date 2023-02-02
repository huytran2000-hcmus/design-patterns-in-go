package remote

type Device interface {
	GetName() string
	IsEnable() bool
	Enable()
	Disable()
	GetChannel() int
	SetChannel(int)
	ChannelUp()
	ChannelDown()
	GetVolume() int
	SetVolume(int)
	VolumeUp()
	VolumeDown()
}
