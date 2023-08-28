package queue

import (
	"testing"
)

func TestPolynomialAdd(t *testing.T) {
	p1 := NewPolynomial(
		[]indeterminate{
			{coefficient: 3, power: 3},
			{coefficient: 2, power: 1},
			{coefficient: 1, power: 0},
		},
	)

	p2 := NewPolynomial(
		[]indeterminate{
			{coefficient: 4, power: 3},
			{coefficient: -2, power: 1},
			{coefficient: 5, power: 2},
		},
	)

	expected := []indeterminate{
		{coefficient: 7, power: 3},
		{coefficient: 5, power: 2},
		{coefficient: 0, power: 1},
		{coefficient: 1, power: 0},
	}

	if !equal(expected, p1.Add(p2)) {
		t.Errorf("Polynomial addition failed")
	}
}

func TestPolynomialAddWithEmptyPolynomial(t *testing.T) {
	p1 := NewPolynomial(
		[]indeterminate{
			{coefficient: 3, power: 3},
		},
	)

	p2 := NewPolynomial([]indeterminate{})

	expected := []indeterminate{
		{coefficient: 3, power: 3},
	}

	if !equal(expected, p1.Add(p2)) {
		t.Errorf("Polynomial addition with empty polynomial failed")
	}
}

func TestPolynomialMultipliedBy(t *testing.T) {
	p1 := NewPolynomial(
		[]indeterminate{
			{coefficient: 3, power: 2},
			{coefficient: 2, power: 1},
			{coefficient: 1, power: 0},
		},
	)

	p2 := NewPolynomial(
		[]indeterminate{
			{coefficient: 4, power: 1},
			{coefficient: -2, power: 0},
		},
	)

	expected := []indeterminate{
		{coefficient: 12, power: 3},
		{coefficient: 2, power: 2},
		{coefficient: 10, power: 1},
		{coefficient: -2, power: 0},
	}

	result := p1.MultipliedBy(p2)

	if !equal(expected, result) {
		t.Errorf("Polynomial multiplication failed")
	}
}

func equal(expected []indeterminate, results *Polynomial) bool {
	m1 := map[uint]int{}
	for _, v := range expected {
		m1[v.power] = v.coefficient
	}
	ind := (*indeterminate)(results)
	if ind == nil {
		return false
	}
	for ind != nil {
		if v, ok := m1[ind.power]; !ok || m1[ind.power] != v {
			return false
		}
		ind = ind.next
	}
	return true
}
