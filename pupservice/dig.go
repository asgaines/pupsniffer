package pupservice

import (
	"encoding/json"
	"fmt"

	"github.com/asgaines/pupsniffr/pup"
)

func ParsePups(j []byte) ([]int, error) {
	var pack pup.Pack

	if err := json.Unmarshal(j, &pack); err != nil {
		return nil, fmt.Errorf("Error unmarshaling the pack: %s", err)
	}

	pupIDs, err := pack.GetPupIDs()
	if err != nil {
		return nil, err
	}

	return pupIDs, nil
}

func ParsePup(j []byte) (*pup.Pup, error) {
	var pup pup.Pup

	if err := json.Unmarshal(j, &pup); err != nil {
		return nil, fmt.Errorf("Error unmarshaling a pup: %s", err)
	}

	return &pup, nil
}
