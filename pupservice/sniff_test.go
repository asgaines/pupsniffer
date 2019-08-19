package pupservice

import (
	"reflect"
	"testing"
)

// TestSniffOutNew compares the pup IDs retrieved (mocked in test case) against the most recent kennel in the testdata/kennel directory, ensuring the correct new IDs are reported
func TestSniffOutNew(t *testing.T) {
	pupsvc := newMockPupSvc()

	cases := []struct {
		message  string
		pupIDs   []int
		expected []int
		err      error
	}{
		{
			message:  "Everything except for 66666666 was already in most recent kennel (in testdata/kennel dir)",
			pupIDs:   []int{11111111, 22222222, 33333333, 44444444, 55555555, 66666666},
			expected: []int{66666666},
			err:      nil,
		},
		{
			message:  "All pup IDs being checked against are new (not in most recent kennel)",
			pupIDs:   []int{18971235, 89713073, 76137893},
			expected: []int{18971235, 89713073, 76137893},
			err:      nil,
		},
		{
			message:  "Pup IDs with a more limited set than in the kennel yields no new pups",
			pupIDs:   []int{11111111, 22222222},
			expected: []int{},
			err:      nil,
		},
		{
			message:  "No pup IDs yields no new pups",
			pupIDs:   []int{},
			expected: []int{},
			err:      nil,
		},
	}

	for _, c := range cases {
		t.Run(c.message, func(t *testing.T) {
			newPupIDs, err := pupsvc.SniffOutNew(c.pupIDs)
			if err != c.err {
				t.Errorf("Expected %v, got %v", c.err, err)
			}

			if !reflect.DeepEqual(newPupIDs, c.expected) {
				t.Errorf("Expected %v, got %v", c.expected, newPupIDs)
			}
		})
	}
}
