package graph

import (
	"testing"
)

func TestFindTwColoring(t *testing.T) {
	tests := loadTestsFromJSON("testdata/col-tests.json", []string{"Bipartite"})

	if len(tests) == 0 {
		t.Errorf("no tests generated")
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			err := test.validate()
			if err != nil {
				t.Errorf("invalid input [%s]", test.Name)
			}

			_, err = test.FindTwoColoring()
			gotErr := (err == nil)
			wantErr := test.Bipartite
			if wantErr != gotErr {
				t.Errorf("Got %s but was expecting %t", err, test.Bipartite)
			}
		})
	}
}
