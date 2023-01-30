package daikin

import "fmt"

func TestAC(remote *ACRemote) {
	min, temperature, max := remote.GetTemperature()
	fmt.Printf("Current temperature: %d°C\n", temperature)
	err := remote.SetTemperature(max)
	err = remote.SetTemperature(max + 1)
	fmt.Println("Error: ", err)
	_, temperature, _ = remote.GetTemperature()
	fmt.Printf("Current temperature: %d°C\n", temperature)
	err = remote.SetTemperature(min)
	err = remote.SetTemperature(min - 1)
	_, temperature, _ = remote.GetTemperature()
	fmt.Println("Error: ", err)
	fmt.Printf("Current temperature: %d°C\n", temperature)
}
