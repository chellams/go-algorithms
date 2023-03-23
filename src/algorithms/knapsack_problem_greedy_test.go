package algorithms

import (
	"math"
	"testing"
)

func TestKnapsackGreedy(t *testing.T) {
	result, err := knapsackGreedy(createTestData(), 15)
	if err != nil {
		t.Errorf("error %d", err)
	}
	if getNumberWithPrecision(result, 2) != 55.33 {
		t.Errorf("%f not an optimal solution", result)
	}
}

func TestKnapsackGreedy_SimpleData(t *testing.T) {
	var knapsacks []knapsack
	knapsacks = append(knapsacks, knapsack{price: 15, weight: 10})
	knapsacks = append(knapsacks, knapsack{price: 24, weight: 15})
	knapsacks = append(knapsacks, knapsack{price: 25, weight: 18})

	result, err := knapsackGreedy(knapsacks, 20)
	if err != nil {
		t.Errorf("error %d", err)
	}
	if getNumberWithPrecision(result, 2) != 31.50 {
		t.Errorf("%f not an optimal solution", result)
	}
}

func TestKnapsackGreedy_NegativeCapacity(t *testing.T) {
	_, err := knapsackGreedy(createTestData(), -3)
	if err == nil {
		t.Errorf("error %d", err)
	}
}

func TestKnapsackGreedy_NegativeWeight(t *testing.T) {
	var knapsacks []knapsack
	knapsacks = append(knapsacks, knapsack{price: 15, weight: 10})
	knapsacks = append(knapsacks, knapsack{price: 24, weight: -15})
	knapsacks = append(knapsacks, knapsack{price: 25, weight: 18})

	_, err := knapsackGreedy(knapsacks, 20)
	if err == nil {
		t.Errorf("error %d", err)
	}
}

func TestKnapsackGreedy_NegativePrice(t *testing.T) {
	var knapsacks []knapsack
	knapsacks = append(knapsacks, knapsack{price: 15, weight: 10})
	knapsacks = append(knapsacks, knapsack{price: 24, weight: -15})
	knapsacks = append(knapsacks, knapsack{price: 25, weight: 18})

	_, err := knapsackGreedy(knapsacks, 20)
	if err == nil {
		t.Errorf("error %d", err)
	}
}

func createTestData() []knapsack {
	var knapsacks []knapsack
	knapsacks = append(knapsacks, knapsack{price: 10, weight: 2})
	knapsacks = append(knapsacks, knapsack{price: 5, weight: 3})
	knapsacks = append(knapsacks, knapsack{price: 15, weight: 5})
	knapsacks = append(knapsacks, knapsack{price: 7, weight: 7})
	knapsacks = append(knapsacks, knapsack{price: 6, weight: 1})
	knapsacks = append(knapsacks, knapsack{price: 18, weight: 4})
	knapsacks = append(knapsacks, knapsack{price: 3, weight: 1})
	return knapsacks
}

func getNumberWithPrecision(num float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(num*ratio) / ratio
}
