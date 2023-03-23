package algorithms

import (
	"errors"
	"sort"
)

type knapsack struct {
	price    float64
	weight   float64
	solution float64
	optimal  float64
}

func knapsackGreedy(knapsacks []knapsack, capacity int) (float64, error) {

	if capacity < 0 {
		return 0.0, errors.New("negative capacity is not allowed")
	}

	for _, knapsack := range knapsacks {
		if knapsack.price < 0 || knapsack.weight < 0 {
			return 0.0, errors.New("negative price or negative value is not allowed")
		}
	}

	for i := 0; i < len(knapsacks); i++ {
		knapsacks[i].solution = knapsacks[i].price / knapsacks[i].weight
	}

	sort.Slice(knapsacks, func(i, j int) bool {
		if knapsacks[i].solution == knapsacks[j].solution {
			return knapsacks[i].weight > knapsacks[i].weight
		}
		return knapsacks[i].solution > knapsacks[j].solution
	})

	max := float64(capacity)
	totalProfit := 0.0
	for i := 0; i < len(knapsacks); i++ {
		if knapsacks[i].weight < max {
			max -= knapsacks[i].weight
			totalProfit += knapsacks[i].price
			knapsacks[i].optimal = 1.0
		} else {
			knapsacks[i].optimal = max / knapsacks[i].weight
			totalProfit += knapsacks[i].price * (max / knapsacks[i].weight)
			break
		}
	}

	return totalProfit, nil
}
