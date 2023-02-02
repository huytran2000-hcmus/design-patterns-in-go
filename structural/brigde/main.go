package main

import (
	"fmt"

	radio2 "github.com/huytran2000-hcmus/design-patterns-in-go/structural/brigde/implementation/media/radio"

	"github.com/huytran2000-hcmus/design-patterns-in-go/structural/brigde/abstraction/remote"
	"github.com/huytran2000-hcmus/design-patterns-in-go/structural/brigde/implementation/media/television"
)

func main() {
	tv := television.New()
	remoteTV := remote.NewBasicRemote(tv)
	fmt.Println(remoteTV.GetReport())
	fmt.Println("Turn on")
	remoteTV.Power()
	fmt.Println(remoteTV.GetReport())
	fmt.Println("Forward to next channel")
	remoteTV.ChannelUp()
	fmt.Println("Turn up volume one notch")
	remoteTV.VolumeUp()
	fmt.Println("Turn up volume one notch")
	remoteTV.VolumeUp()
	fmt.Println(remoteTV.GetReport())

	radio := radio2.New()
	remoteRadio := remote.NewAdvanceRemote(radio)
	fmt.Println(remoteRadio.GetReport())
	fmt.Println("Turn on")
	remoteRadio.Power()
	fmt.Println("Turn down volume one notch")
	remoteRadio.VolumeDown()
	fmt.Println("Go back one channel")
	remoteRadio.ChannelDown()
	fmt.Println(remoteRadio.GetReport())
	fmt.Println("Go to channel 10")
	remoteRadio.SetChannel(10)
	fmt.Println("Set volume to 20")
	remoteRadio.SetVolume(20)
	fmt.Println(remoteRadio.GetReport())
}
