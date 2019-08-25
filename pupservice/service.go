package pupservice

import (
	"context"

	"github.com/asgaines/pupsniffr/pup"
	"github.com/asgaines/pupsniffr/pupservice/fetcher"
)

type PupService interface {
	FetchPupIDs(ctx context.Context) ([]int, error)
	FetchPups(ctx context.Context, pupIDs []int) ([]*pup.Pup, error)
	SniffOutNew(pupIDs []int) ([]int, error)
	KennelPups(pupIDs []int) error
	FilterPups(pups []*pup.Pup) ([]*pup.Pup, error)
}

func New(f fetcher.Fetcher, kennelPath string) PupService {
	return &pupsvc{
		fetcher:    f,
		kennelPath: kennelPath,
	}
}

type pupsvc struct {
	fetcher    fetcher.Fetcher
	kennelPath string
}
