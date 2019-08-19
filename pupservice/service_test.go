package pupservice

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/asgaines/pupsniffer/pupservice/fetcher"
)

type mockFetcher struct {
	fetcher.Fetcher
}

func newMockPupSvc() PupService {
	return New(&mockFetcher{
		Fetcher: fetcher.NewFetcher(),
	}, "testdata/kennel")
}

func (m *mockFetcher) FetchPack(ctx context.Context) ([]byte, error) {
	return m.fetch("testdata/fetch/pack/pack.json")
}

func (m *mockFetcher) FetchPup(ctx context.Context, pupID int) ([]byte, error) {
	path := fmt.Sprintf("testdata/fetch/pups/%d.json", pupID)
	return m.fetch(path)
}
func (m *mockFetcher) fetch(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}
