package pupservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// KennelPups stores the pup IDs so they can be used to sniff out new pups during the next fetch
func (s *pupsvc) KennelPups(pupIDs []int) error {
	now := time.Now().UTC().Format(time.RFC3339)

	fname := fmt.Sprintf("%s/%s.json", s.kennelPath, now)

	j, err := json.Marshal(pupIDs)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fname, j, 0644); err != nil {
		return err
	}

	return nil
}
