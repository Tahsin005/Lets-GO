package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	// first approach
	// maxDay := 0
	// for i := 0; i < len(costs); i++ {
	// 	if costs[i].day > maxDay {
	// 		maxDay = costs[i].day
	// 	}
	// }

	// costSlice := make([]float64, maxDay + 1)

	// for i := 0; i < len(costs); i++ {
	// 	day := costs[i].day
	// 	value := costs[i].value

	// 	costSlice[day] += value
	// }
	// return costSlice

	
	// alternative approach
	costByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		currCost := costs[i]
		for currCost.day >= len(costByDay) {
			costByDay = append(costByDay, 0.0)
		}
		costByDay[currCost.day] += currCost.value
	}
	return costByDay
}

func test(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})
}