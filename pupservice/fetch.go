package pupservice

import (
	"context"

	"github.com/asgaines/pupsniffr/pup"
)

func (p *pupsvc) FetchPupIDs(ctx context.Context) ([]int, error) {
	b, err := p.fetcher.FetchPack(ctx)
	if err != nil {
		return nil, err
	}

	pupIDs, err := ParsePups(b)
	if err != nil {
		return nil, err
	}

	return pupIDs, nil
}

func (p *pupsvc) FetchPups(ctx context.Context, pupIDs []int) ([]*pup.Pup, error) {
	hole := make(chan *pup.Pup)
	errs := make(chan error)

	ctx, cancel := context.WithCancel(ctx)

	for _, pupID := range pupIDs {
		go func(pupID int) {
			b, err := p.fetcher.FetchPup(ctx, pupID)
			if err != nil {
				errs <- err
				return
			}

			pup, err := ParsePup(b)
			if err != nil {
				errs <- err
				return
			}

			hole <- pup
		}(pupID)
	}

	pups := make([]*pup.Pup, 0, len(pupIDs))

	for range pupIDs {
		select {
		case pup := <-hole:
			pups = append(pups, pup)
		case err := <-errs:
			cancel()
			return pups, err
		}
	}

	return pups, nil
}
