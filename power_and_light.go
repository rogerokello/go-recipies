package main

import (
	"fmt"
)

func main() {
	var plantCapacities []float64

	plantCapacities = []float64{30, 30, 30, 60, 60, 100}

	var capacity float64 = plantCapacities[0] + plantCapacities[1] + plantCapacities[2] + plantCapacities[3] + plantCapacities[4] + plantCapacities[5]

	var gridload = 75.

	utilization := gridload / capacity

	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", gridload)
	fmt.Printf("%-20s%.1f\n", "Utilization: ", utilization*100)
}
