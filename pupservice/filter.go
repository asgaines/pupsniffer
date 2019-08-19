package pupservice

import "github.com/asgaines/pupsniffer/pup"

func (p *pupsvc) FilterPups(pups []*pup.Pup) ([]*pup.Pup, error) {
	filteredPups := []*pup.Pup{}

	for _, pup := range pups {
		available, err := pup.IsAvailable()
		if err != nil {
			return filteredPups, err
		}

		rightAge, err := pup.IsRightAge()
		if err != nil {
			return filteredPups, err
		}

		if available && rightAge {
			filteredPups = append(filteredPups, pup)
		}
	}

	return filteredPups, nil
}
