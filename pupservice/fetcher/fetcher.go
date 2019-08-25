package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"time"

	"github.com/asgaines/pupsniffr/config"
)

type Fetcher interface {
	FetchPack(ctx context.Context) ([]byte, error)
	FetchPup(ctx context.Context, pupID int) ([]byte, error)
	FetchRecentKennel(kennelPath string) ([]int, error)
}

func NewFetcher() Fetcher {
	return &fetcher{
		kennelPath: "./kennel/",
	}
}

type fetcher struct {
	kennelPath string
}

var client = http.Client{
	Timeout: 30 * time.Second,
}

func (f *fetcher) FetchPack(ctx context.Context) ([]byte, error) {
	return f.fetch(ctx, config.PackURL)
}

func (f *fetcher) FetchPup(ctx context.Context, pupID int) ([]byte, error) {
	url := fmt.Sprintf(config.PupURL, pupID)
	return f.fetch(ctx, url)
}

func (f *fetcher) fetch(ctx context.Context, url string) ([]byte, error) {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	r = r.WithContext(ctx)

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unable to retrieve data from: %v (status: %d)", url, resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f *fetcher) FetchRecentKennel(kennelPath string) ([]int, error) {
	files, err := ioutil.ReadDir(kennelPath)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return []int{}, nil
	}

	recentF := path.Join(kennelPath, files[len(files)-1].Name())

	b, err := ioutil.ReadFile(recentF)
	if err != nil {
		return nil, err
	}

	var kennelPupIDs []int
	if err := json.Unmarshal(b, &kennelPupIDs); err != nil {
		return nil, err
	}

	return kennelPupIDs, nil
}
