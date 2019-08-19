package pupservice

// SniffForNewPups takes a set of pup IDs and detects if there are any new pups, not previously seen
func (s *pupsvc) SniffOutNew(pupIDs []int) ([]int, error) {
	kennelPupIDs, err := s.fetcher.FetchRecentKennel(s.kennelPath)
	if err != nil {
		return nil, err
	}

	kennelPupSet := make(map[int]bool)
	for _, pupID := range kennelPupIDs {
		kennelPupSet[pupID] = true
	}

	newPupIDs := []int{}
	for _, pupID := range pupIDs {
		if !kennelPupSet[pupID] {
			newPupIDs = append(newPupIDs, pupID)
		}
	}

	return newPupIDs, nil
}
