package pupservice

import (
	"bytes"
	"context"
	"io"

	"github.com/asgaines/pupsniffer/pup"
	"github.com/asgaines/pupsniffer/pupservice/fetcher"
)

type PupService interface {
	FetchPupIDs(ctx context.Context) ([]int, error)
	FetchPups(ctx context.Context, pupIDs []int) ([]*pup.Pup, error)
	SniffOutNew(pupIDs []int) ([]int, error)
	KennelPups(pupIDs []int) error
	FilterPups(pups []*pup.Pup) ([]*pup.Pup, error)
	PupReport(total int, pups []*pup.Pup, wr io.Writer) error
	Mailman(buf *bytes.Buffer, recipients []string) error
}

func New(f fetcher.Fetcher, kennelPath, staticPath string) PupService {
	return &pupsvc{
		fetcher:    f,
		kennelPath: kennelPath,
		staticPath: staticPath,
	}
}

type pupsvc struct {
	fetcher    fetcher.Fetcher
	kennelPath string
	staticPath string
}
