package pupservice

import (
	"log"

	"github.com/asgaines/pupsniffer/pup"
)

func (p *pupsvc) FilterPups(pups []*pup.Pup) ([]*pup.Pup, error) {
	filteredPups := []*pup.Pup{}

	for _, pup := range pups {
		available, err := pup.IsAvailable()
		if err != nil {
			log.Printf("while filtering, getting availability: %s", err)
			// return filteredPups, err
		}

		isPuppy, err := pup.IsPuppy()
		if err != nil {
			log.Printf("while filtering, getting puppy status: %s", err)
			// return filteredPups, err
		}

		isOlderAndBigger, err := pup.IsOlderAndBigger()
		if err != nil {
			log.Printf("while filtering, getting older and bigger status: %s", err)
			// return filteredPups, err
		}

		if available && (isPuppy || isOlderAndBigger) {
			log.Printf("Pup passes filter. Age: %s, weight: %s", pup.Age, pup.BodyWeight)
			filteredPups = append(filteredPups, pup)
		}
	}

	return filteredPups, nil
}
