package remote

import (
	"fmt"
	"strings"
)

type BasicRemote struct {
	device Device
}

func NewBasicRemote(device Device) *BasicRemote {
	return &BasicRemote{device: device}
}

func (rm *BasicRemote) Power() {
	if !rm.device.IsEnable() {
		rm.device.Enable()
	} else {
		rm.device.Disable()
	}
}

func (rm *BasicRemote) ChannelUp() {
	if rm.device.IsEnable() {
		rm.device.ChannelUp()
	}
}

func (rm *BasicRemote) ChannelDown() {
	if rm.device.IsEnable() {
		rm.device.ChannelDown()
	}
}

func (rm *BasicRemote) VolumeUp() {
	if rm.device.IsEnable() {
		rm.device.VolumeUp()
	}
}

func (rm *BasicRemote) VolumeDown() {
	if rm.device.IsEnable() {
		rm.device.VolumeDown()
	}
}

func (rm *BasicRemote) GetReport() string {
	builder := statusBuilder{}

	name := strings.ToUpper(rm.device.GetName())
	builder.WriteString("========================================================")
	builder.WriteRune('\n')
	if rm.device.IsEnable() {
		builder.WriteString(name + " is enable")
		builder.WriteRune('\n')
		builder.WriteString(fmt.Sprintf("Channel: %d", rm.device.GetChannel()))
		builder.WriteRune('\n')
		builder.WriteString(fmt.Sprintf("Volume: %d", rm.device.GetVolume()))
		builder.WriteRune('\n')
	} else {
		builder.WriteString(name + " is disable")
		builder.WriteRune('\n')
	}
	builder.WriteString("========================================================")

	status := builder.String()
	return status
}

type statusBuilder struct {
	strings.Builder
}

func (sb *statusBuilder) WriteString(s string) (int, error) {
	sb.Builder.WriteRune('|')
	n, err := sb.Builder.WriteString(s)

	return n, err
}
