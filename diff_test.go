package diff

import (
	"testing"
)

func TestSimple(t *testing.T) {
	tests := []struct {
		A      []string
		B      []string
		Result []Result
	}{
		{
			A:      []string{},
			B:      []string{},
			Result: []Result{},
		},
		{
			A:      []string{"hello"},
			B:      []string{"hello"},
			Result: []Result{},
		},
		{
			A: []string{"hello"},
			B: []string{"goodbye"},
			Result: []Result{
				Result{ActionChanged, 0, "goodbye"},
			},
		},
		{
			A: []string{"1"},
			B: []string{"1", "2"},
			Result: []Result{
				Result{ActionAdded, 1, "2"},
			},
		},
		{
			A: []string{"1", "2"},
			B: []string{"1"},
			Result: []Result{
				Result{ActionRemoved, 1, "2"},
			},
		},
		{
			A: []string{"1", "2", "x", "3"},
			B: []string{"1", "2", "3"},
			Result: []Result{
				Result{ActionRemoved, 2, "x"},
			},
		},
		{
			A: []string{"1", "2", "x", "y", "3"},
			B: []string{"1", "2", "3"},
			Result: []Result{
				Result{ActionRemoved, 2, "x"},
				Result{ActionRemoved, 3, "y"},
			},
		},
		{
			A: []string{"1", "2", "x", "y", "3"},
			B: []string{"1", "2", "y"},
			Result: []Result{
				Result{ActionRemoved, 2, "x"},
				Result{ActionRemoved, 4, "3"},
			},
		},
	}

	for i, test := range tests {
		result := Diff(test.A, test.B)
		if len(result) != len(test.Result) {
			t.Errorf("%d. Expected\n%v\nGot:\n%v\n", i, test.Result, result)
		}
	}
}
