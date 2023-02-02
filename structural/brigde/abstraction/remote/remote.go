package remote

type Remote interface {
	Power()
	ChannelUp()
	ChannelDown()
	VolumeUp()
	VolumeDown()
	GetReport() string
}
