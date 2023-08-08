package graph

import (
	"testing"
)

func TestFindTwColoring(t *testing.T) {
	tests := loadTestsFromJSON("testdata/col-tests.json")

	if len(tests) == 0 {
		t.Errorf("no tests generated")
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			g, err := New(test.NumVertices, test.Adj)
			if err != nil {
				t.Errorf("invalid input [%s]", name)
			}

			_, err = g.FindTwoColoring()
			gotErr := (err == nil)
			wantErr := test.ExpectedBipartite
			if wantErr != gotErr {
				t.Errorf("got %s but was expecting %t", err, test.ExpectedBipartite)
			}
		})
	}
}
