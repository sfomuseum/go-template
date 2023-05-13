package funcs

import (
	"testing"
)

func TestMod(t *testing.T) {

	m := Mod(3.0, 2.0)

	if m != 1.0 {
		t.Fatalf("Invalid mod. Expected 1.0 but got %f", m)
	}
}

func TestIsEven(t *testing.T) {

	iseven := []int{
		2,
		24,
		264,
	}

	isodd := []int{
		1,
		55,
		127,
	}

	for _, i := range iseven {

		if !IsEven(i) {
			t.Fatalf("Wut? %d is even", i)
		}
	}

	for _, i := range isodd {

		if IsEven(i) {
			t.Fatalf("Wut? %d is not even", i)
		}
	}
}

func TestIsOdd(t *testing.T) {

	iseven := []int{
		2,
		24,
		264,
	}

	isodd := []int{
		1,
		55,
		127,
	}

	for _, i := range iseven {

		if IsOdd(i) {
			t.Fatalf("Wut? %d is not odd", i)
		}
	}

	for _, i := range isodd {

		if !IsOdd(i) {
			t.Fatalf("Wut? %d is odd", i)
		}
	}
}
