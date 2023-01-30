package panasonic

import "fmt"

func TestAC(remote *ACRemote) {
	printTemperature(remote.GetTemperature())
	increaseToMaxTemperature(remote)
	printTemperature(remote.GetTemperature())
	increaseToMinimumTemperature(remote)
	printTemperature(remote.GetTemperature())
}

func increaseToMaxTemperature(remote *ACRemote) {
	var err error
	for i := 0; i < 20; i++ {
		err = remote.IncreaseTemperature()
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
	}
}

func increaseToMinimumTemperature(remote *ACRemote) {
	var err error
	for i := 0; i < 20; i++ {
		err = remote.DecreaseTemperature()
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
	}
}

func printTemperature(temperature int) {
	fmt.Printf("Current temperature: %d°C\n", temperature)
}
