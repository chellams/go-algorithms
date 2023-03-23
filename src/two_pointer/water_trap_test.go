package two_pointer

import (
	"testing"
)

func TestMaxWaterTrap(t *testing.T) {
	ints := []int{4, 8, 1, 2, 3, 9}
	expected := 32
	actual := maxWater(ints)

	if expected != actual {
		t.Errorf(" Expected is %d, but received %d", expected, actual)
	}
}

func TestMaxWaterTrap_Nil(t *testing.T) {
	expected := 0
	actual := maxWater(nil)

	if expected != actual {
		t.Errorf(" Expected is %d, but received %d", expected, actual)
	}
}

func TestMaxWaterTrap_Empty(t *testing.T) {
	ints := []int{}
	expected := 0
	actual := maxWater(ints)

	if expected != actual {
		t.Errorf(" Expected is %d, but received %d", expected, actual)
	}

	ints = []int{1}
	expected = 0
	actual = maxWater(ints)
	if expected != actual {
		t.Errorf(" Expected is %d, but received %d", expected, actual)
	}
}
