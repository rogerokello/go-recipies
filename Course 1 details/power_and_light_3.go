package main

import (
	"fmt"
	"strings"
)

func main() {
	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavailable},
		PowerPlant{solar, 40, inactive},
	}

	grid := PowerGrid{300, plants}

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
			grid.generatePlantReport()
		case "2":
			grid.generateGridReport()
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

type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "InActive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for idx, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", idx)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type:", p.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity:", p.capacity)
		fmt.Printf("%-20s%s\n", "Status:", p.status)
		fmt.Println("")
	}
}

func (pg *PowerGrid) generateGridReport() {
	capacity := 0.

	for _, p := range pg.plants {
		if p.status == active {
			capacity += p.capacity
		}
	}

	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", pg.load)
	fmt.Printf("%-20s%.2f%%\n", "Utilization:", (pg.load/capacity)*100)
}
