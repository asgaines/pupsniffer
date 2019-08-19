package pupservice

import (
	"context"
	"reflect"
	"testing"
)

func TestFetchPupIDs(t *testing.T) {
	mockPupSvc := newMockPupSvc()

	cases := []struct {
		message  string
		expected []int
	}{
		{
			message: "All correct pup IDs are pulled from the pack",
			expected: []int{
				39796455,
				41306187,
				41666771,
				41677095,
				41741516,
				41928161,
				42082729,
				42125851,
				42388940,
				42388945,
				42388948,
				42388949,
				42388951,
				42388953,
				42388954,
				42388958,
				42399012,
				42422649,
				42427772,
				42429042,
				42429899,
				42431336,
				42431362,
				42431385,
				42433570,
				42442033,
				42446250,
				42450228,
				42450259,
				42450274,
				42450285,
				42450292,
				42450293,
				42450297,
				42450300,
				42450310,
				42450315,
				42450320,
				42457019,
				42461508,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.message, func(t *testing.T) {
			out, err := mockPupSvc.FetchPupIDs(context.Background())
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(out, c.expected) {
				t.Errorf("Expected %v, got %v", c.expected, out)
			}
		})
	}
}

func TestFetchPups(t *testing.T) {
	mockPupSvc := newMockPupSvc()

	cases := []struct {
		message       string
		in            []int
		expectedNames []string
	}{
		{
			message: "Fetching all pups by ID finds and parses them correctly",
			in: []int{
				39796455,
				41306187,
				41666771,
				41677095,
				41741516,
				41928161,
				42082729,
				42125851,
				42388940,
				42388945,
				42388948,
				42388949,
				42388951,
				42388953,
				42388954,
				42388958,
				42399012,
				42422649,
				42427772,
				42429042,
				42429899,
				42431336,
				42431362,
				42431385,
				42433570,
				42442033,
				42446250,
				42450228,
				42450259,
				42450274,
				42450285,
				42450292,
				42450293,
				42450297,
				42450300,
				42450310,
				42450315,
				42450320,
				42457019,
				42461508,
			},
			expectedNames: []string{
				"Poppy",
				"Blue",
				"Taffy",
				"Skipper",
				"Snoop",
				"Bruno",
				"Alfie",
				"Ramona",
				"Suzy",
				"Sam",
				"Sonny",
				"Sunshine",
				"Starshine",
				"Sean",
				"Sarah",
				"Seth",
				"Dewey",
				"Rave",
				"Matilda",
				"Captain",
				"Dolphin",
				"Cody",
				"Dobie",
				"Lenny",
				"Bella",
				"July",
				"Marley",
				"Blu",
				"Linus",
				"Cosmos",
				"Dolly",
				"Toby",
				"Carrie",
				"Garth",
				"Clint",
				"Duke",
				"Pippa",
				"Cal",
				"Teddy",
				"Sdah",
			},
		},
	}

	settify := func(ss []string) map[string]int {
		m := make(map[string]int)

		for _, s := range ss {
			if _, ok := m[s]; ok {
				m[s]++
			} else {
				m[s] = 1
			}
		}

		return m
	}

	for _, c := range cases {
		t.Run(c.message, func(t *testing.T) {
			out, err := mockPupSvc.FetchPups(context.Background(), c.in)
			if err != nil {
				t.Error(err)
			}

			pupNames := make([]string, 0, len(out))
			for _, pup := range out {
				pupNames = append(pupNames, pup.AnimalName)
			}

			if !reflect.DeepEqual(settify(pupNames), settify(c.expectedNames)) {
				t.Errorf("Expected %v, got %v", c.expectedNames, pupNames)
			}
		})
	}
}
