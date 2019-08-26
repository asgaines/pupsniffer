package pupservice

import (
	"reflect"
	"testing"

	"github.com/asgaines/pupsniffer/pup"
)

func TestFilterPups(t *testing.T) {
	mockPupSvc := newMockPupSvc()

	cases := []struct {
		pups     []*pup.Pup
		expected []*pup.Pup
		err      error
	}{
		{
			pups: []*pup.Pup{
				{
					AnimalName: "Rufus",
					Age:        "36",
					OnHold:     "No",
				},
				{
					AnimalName: "Shelley",
					Age:        "41",
					OnHold:     "No",
				},
				{
					AnimalName: "Boris",
					Age:        "101",
					OnHold:     "No",
				},
				{
					AnimalName: "Pip",
					Age:        "3",
					OnHold:     "No",
				},
				{
					AnimalName: "Polly",
					Age:        "33",
					OnHold:     "Yes",
				},
			},
			expected: []*pup.Pup{
				{
					AnimalName: "Rufus",
					Age:        "36",
					OnHold:     "No",
				},
				{
					AnimalName: "Shelley",
					Age:        "41",
					OnHold:     "No",
				},
			},
			err: nil,
		},
	}

	for _, c := range cases {
		filteredPups, err := mockPupSvc.FilterPups(c.pups)
		if err != c.err {
			t.Errorf("Expected %s, got %s", c.err, err)
		}

		if !reflect.DeepEqual(filteredPups, c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, filteredPups)
		}
	}
}
