package graph

import (
	"fmt"
	"testing"
)

func TestComponentFinding(t *testing.T) {
	tests := loadTestsFromJSON("testdata/components-tests.json")

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			g, err := New(test.NumVertices, test.Adj)
			if err != nil {
				t.Errorf("while generating graph: %s", err)
			}

			g.updateConnectedComponents()
			got := g.components

			want := test.ExpectedComponents
			if len(got) != len(want) {
				t.Errorf("got component size %d different from expected size %d", len(got), len(want))
			}
			for i := range want {
				err := isEqual[int](got[i], want[i])
				if err != nil {
					t.Errorf("while checking components %s: %s", name, err)
				}
			}
		})
	}
}

func isEqual[K comparable](got []K, want []K) error {
	if len(want) != len(got) {
		return fmt.Errorf("got length %d different from expected length %d", len(got), len(want))
	}

	for i := range want {
		if want[i] != got[i] {
			return fmt.Errorf("got entry %v different from expected entry %v at position %d", want[i], got[i], i)
		}
	}
	return nil
}
