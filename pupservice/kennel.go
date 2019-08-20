package pupservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

// KennelPups stores the pup IDs so they can be used to sniff out new pups during the next fetch
func (s *pupsvc) KennelPups(pupIDs []int) error {
	now := time.Now().Format("20060102_150405")

	fname := fmt.Sprintf("%s.json", now)
	filepath := fmt.Sprintf(filepath.Join(s.kennelPath, fname))

	j, err := json.Marshal(pupIDs)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, j, 0644); err != nil {
		return err
	}

	return nil
}
