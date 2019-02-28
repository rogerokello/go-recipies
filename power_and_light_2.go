package main

import (
	"fmt"
)

func main() {

	plantCapacities := []float64{30, 30, 30, 60, 60, 100}

	activePlants := []int{0, 1}

	gridLoad := 75.

	for {
		println("**** Menu Options ****")
		println("1) Generate power plant report")
		println("2) Generate power grid report")
		println("3) To exit")
		print("Please choose an option: ")

		var option string

		fmt.Scanln(&option)

		switch option {
		case "1":
			plantReport(plantCapacities)
		case "2":
			gridReport(activePlants, plantCapacities, gridLoad)
		default:
			fmt.Println("Unknown option, no action taken")
			fmt.Println()
		}

		if option == "3" {
			fmt.Println("Exiting now, Bye Bye !!!!!")
			break
		}
	}

}

func plantReport(plantCapacities []float64) {
	println()
	println("**Power Plant report**")
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
	}
	fmt.Println()
}

func gridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantID := range activePlants {
		capacity += plantCapacities[plantID]
	}
	println()
	println("**Power Grid report**")
	fmt.Printf("%-20s%.0f%%\n", "Capacity", capacity)
	fmt.Printf("%-20s%.0f%%\n", "Load", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization", (gridLoad/capacity)*100)
	fmt.Println()
}
